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
)

// ConfigPathEnvVar identifies the environment variable that overrides the default location (i.e. user's home directory) of the HSC configuration file.
const ConfigPathEnvVar = "HSC_CONFIG_HOME"

// NOTE: This should get moved to ghutils
// GitHubTokenEnvVar identifies the environment variable that contains a user's GitHub API token.
const GitHubTokenEnvVar = "GITHUB_TOKEN"

// ConfigFileName is the name of the HSC configuration file.
const ConfigFileName = ".hsc"

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
		errors = append(errors, fmt.Errorf("config: dir is required"))
	} else if _, err := os.Stat(c.Dir); os.IsNotExist(err) {
		errors = append(errors, fmt.Errorf("config: dir '%s' does not exist", c.Dir))
	}

	if c.Token == "" {
		errors = append(errors, fmt.Errorf("config: a GitHub oauth token is required. Goto https://github.com/blog/1509-personal-api-tokens to learn about tokens"))
	}

	if c.User == "" {
		errors = append(errors, fmt.Errorf("config: a user is required. This should be your GitHub username"))
		//	} else if _, err := ghutil.User(c.User); ghutil.IsNotExist(err) {
		//		errors = append(errors, fmt.Errorf("config: user '%s' does not exist in GitHub", c.User))
	}

	if c.Org != "" {
		//		if _, err := ghutil.Org(c.Org); ghutil.IsNotExist(err) {
		//			errors = append(errors, fmt.Errorf("config: org '%s' does not exist in GitHub", c.Org))
		//		}
	}

	if len(errors) > 0 {
		return &multierror.Error{Errors: errors}
	}

	return nil
}

func getField(c *Config, field string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

// HTTPClient returns an client built with Token. This client is suitable for OAuth authentication
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
