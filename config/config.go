package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"
	"unsafe"

	"github.com/hashicorp/terraform/helper/multierror"
	"github.com/mitchellh/go-homedir"
	ini "github.com/rakyll/goini"

	"net/http"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

const (
	// ConfigPathEnvVar identifies the environment variable that overrides the default location (i.e. user's home directory) of the HSC configuration file.
	ConfigPathEnvVar = "HSC_CONFIG_HOME"

	// GitHubTokenEnvVar identifies the environment variable that contains a user's GitHub API token.
	GitHubTokenEnvVar = "GITHUB_TOKEN"

	// ConfigFileName is the name of the HSC configuration file.
	ConfigFileName = ".hsc"

	// Validation Error Messages:

	// ErrMissingDir is an error message indicating a Config without a Dir variable
	ErrMissingDir = "dir is required"

	// ErrDirDoesNotExist is an error message indicating a Config's Dir variable references a directory that does not exist
	ErrDirDoesNotExist = "directory does not exist"

	// ErrMissingUser is an error message indicating a Config without a User variable
	ErrMissingUser = "a user is required. This should be your GitHub username"

	// ErrAuthUser is an error message indicating a Config's User cannot be authenticated against GitHub (using Token)
	ErrAuthUser = "unable to authenticate against GitHub using the oauth token provided"

	// ErrUserMismatch is an error message indicating a Config's User variable does not match data returned from GitHub
	ErrUserMismatch = "username provided does not match GitHub login or email associated with your authenicated user"

	// ErrMissingToken is an error message indicating a Config without a Token variable
	ErrMissingToken = "a GitHub oauth token is required. Goto https://github.com/blog/1509-personal-api-tokens to learn about tokens"

	// ErrOrgDoesNotExist is an error message indicating a Config's Org variable references an organization that does not exist on GitHub
	ErrOrgDoesNotExist = "org does not exist on GitHub"

	// ErrUserNotOrgMember is an error message indicating a Config's User is not a member of the GitHub organization identified by the Org variable
	ErrUserNotOrgMember = "user is not a member of the organization provided. Contact GitHub organization owner to have yourself added to organization before re-trying this command"
)

// Config is the configuration for HSC. The idea is to collect (as configuration) any data attributes that tend to be used across sub-commands.
type Config struct {
	Org   string
	User  string
	Dir   string
	Token string
}

// WriteConfig attempts to write the configuration file to disk
func WriteConfig(config *Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	configdir, err := configDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configdir); os.IsNotExist(err) {
		return err
	}

	writepath := filepath.Join(configdir, ConfigFileName)
	dict := make(ini.Dict)
	section := ""
	dict[section] = make(map[string]string)

	tmpintslice := []int{0}
	ielements := reflect.TypeOf(config).Elem().NumField()
	for i := 0; i < ielements; i++ {
		tmpintslice[0] = i
		f := reflect.TypeOf(config).Elem().FieldByIndex(tmpintslice)
		value := reflect.ValueOf(config).Elem().FieldByName(f.Name).String()
		if value != "" {
			dict[section][strings.ToLower(f.Name)] = value
		}
	}

	if err := ini.Write(writepath, &dict); err != nil {
		return err
	}

	return nil
}

// LoadConfig attempts to read the configuration file and return a Config
func LoadConfig() (*Config, error) {

	configdir, err := configDir()
	if err != nil {
		return nil, err
	}

	return loadConfig(configdir)
}

func loadConfig(path string) (*Config, error) {

	configpath := filepath.Join(path, ConfigFileName)
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		return nil, nil
	}

	dict, err := ini.Load(configpath)
	if err != nil {
		return nil, err
	}

	// could be an empty config file
	if len(dict) == 0 {
		return nil, nil
	}

	configs := dict[""]
	if len(configs) == 0 {
		return nil, nil
	}

	myConfig := &Config{Org: "", User: "", Dir: "", Token: ""}
	tmpintslice := []int{0}
	ielements := reflect.TypeOf(myConfig).Elem().NumField()
	for i := 0; i < ielements; i++ {
		tmpintslice[0] = i
		f := reflect.TypeOf(myConfig).Elem().FieldByIndex(tmpintslice)
		if value, ok := configs[strings.ToLower(f.Name)]; ok {
			reflect.ValueOf(myConfig).Elem().FieldByName(f.Name).SetString(value)
		}
	}

	return myConfig, nil
}

// Validate does some sanity checking of the Config
func (c *Config) Validate() error {
	var errors []error

	if c == nil || unsafe.Sizeof(c) == 0 {
		return fmt.Errorf("config: nil or empty Config instance")
	}

	if c.Dir == "" {
		errors = append(errors, fmt.Errorf(ErrMissingDir))
	} else if _, err := os.Stat(c.Dir); os.IsNotExist(err) {
		errors = append(errors, fmt.Errorf(ErrDirDoesNotExist))
	}

	if c.Token == "" {
		errors = append(errors, fmt.Errorf(ErrMissingToken))
		// Make sure token is invalid so authentication fails
		c.Token = "111222333444aaabbbcccdddeee"
	}

	// get a GitHub client to use for validation against GitHub API
	client := c.GitHubClient()

	bUserValid := false
	if c.User == "" {
		errors = append(errors, fmt.Errorf(ErrMissingUser))
	} else if u, resp, err := c.validGitHubUser(client); err != nil {
		errors = append(errors, fmt.Errorf("fatal!!!: validating user against GitHub resulted in error: %s", err.Error()))
	} else if resp.StatusCode == 401 {
		errors = append(errors, fmt.Errorf(ErrAuthUser))
	} else if *u.Login != c.User && *u.Email != c.User {
		errors = append(errors, fmt.Errorf(ErrUserMismatch))
	} else {
		bUserValid = true
	}

	if c.Org != "" {
		if _, resp, err := c.validGitHubOrg(client); err != nil {
			errors = append(errors, fmt.Errorf("fatal!!!: validating org against GitHub resulted in error: %s", err.Error()))
		} else if resp.StatusCode == 404 {
			errors = append(errors, fmt.Errorf(ErrOrgDoesNotExist))
		} else if bUserValid {
			if ok, _, err := c.isGitHubOrgMember(client); err != nil {
				errors = append(errors, fmt.Errorf("fatal!!!: validating user's org membership against GitHub resulted in error: %s", err.Error()))
			} else if !ok {
				errors = append(errors, fmt.Errorf(ErrUserNotOrgMember))
			}
		}
	}

	if len(errors) > 0 {
		return &multierror.Error{Errors: errors}
	}

	return nil
}

func (c *Config) validGitHubUser(client *github.Client) (*github.User, *github.Response, error) {
	u, resp, err := client.Users.Get("")
	if err != nil && resp.StatusCode != 401 {
		return nil, nil, err
	}

	return u, resp, nil
}

func (c *Config) validGitHubOrg(client *github.Client) (*github.Organization, *github.Response, error) {
	o, resp, err := client.Organizations.Get(c.Org)
	if err != nil && resp.StatusCode != 404 {
		return nil, nil, err
	}

	return o, resp, nil
}

func (c *Config) isGitHubOrgMember(client *github.Client) (bool, *github.Response, error) {
	o, resp, err := client.Organizations.IsMember(c.Org, c.User)
	if err != nil && resp.StatusCode != 404 {
		return false, nil, err
	}

	return o, resp, nil
}

func getField(c *Config, field string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

// GitHubClient returns a GitHub build with Token.
func (c *Config) GitHubClient() *github.Client {
	return github.NewClient(c.HTTPClient())
}

// HTTPClient returns a http client built with Token. This client is suitable for OAuth authentication
func (c *Config) HTTPClient() *http.Client {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: c.GetToken()},
	}

	return t.Client()
}

// GetToken returns the value of the user's GitHub Token environment variable or the user-supplied token value -- whichever was provided.
func (c *Config) GetToken() string {

	rtnval := c.Token
	if c.Token == GitHubTokenEnvVar {
		if envtoken := os.Getenv(GitHubTokenEnvVar); envtoken != "" {
			rtnval = envtoken
		} else {
			rtnval = ""
		}
	}
	return rtnval
}

func configDir() (string, error) {

	// First preference is the HSC_CONFIG_HOME environment variable
	if envdir := os.Getenv(ConfigPathEnvVar); envdir != "" {
		return envdir, nil
	}

	homedir, err := homeDir()
	if err != nil {
		return "", err
	}

	return homedir, nil
}

func homeDir() (string, error) {
	_, err := user.Current()
	if err != nil {
		return "", err
	}

	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	path, err := homedir.Expand(dir)
	if err != nil {
		return "", err
	}

	return path, nil
}
