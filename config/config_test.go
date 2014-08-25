package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// This is the directory where our test fixtures are.
const fixtureDir = "./test-fixtures"

func fixturePath(t *testing.T, name string) string {
	dir, err := filepath.Abs(filepath.Join(fixtureDir, name))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return dir
}

// No config file found should result in nil config struct
func TestNoEnvVarSet(t *testing.T) {

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c != nil {
		t.Fatalf("fail: expecting a %s config", "nil")
	}
}

// No config file found should result in nil config struct
func TestLoadNoFile(t *testing.T) {

	dir := fixturePath(t, "nofile")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c != nil {
		t.Fatalf("fail: expecting a %s config", "nil")
	}
}

// Empty config file found should result in nil config struct
func TestLoadEmpty(t *testing.T) {

	dir := fixturePath(t, "empty")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c != nil {
		t.Fatalf("fail: expecting a %s config", "nil")
	}
}

func TestLoadOnlyOrg(t *testing.T) {

	dir := fixturePath(t, "onlyorg")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "XYZCorp", User: "", Dir: "", Token: ""}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestLoadOnlyUser(t *testing.T) {

	dir := fixturePath(t, "onlyuser")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "", User: "billyjoe", Dir: "", Token: ""}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestLoadOnlyDir(t *testing.T) {

	dir := fixturePath(t, "onlydir")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "", User: "", Dir: "/home/billyjoe/projects", Token: ""}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestLoadOnlyToken(t *testing.T) {

	dir := fixturePath(t, "onlytoken")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "", User: "", Dir: "", Token: "KFDJLDJad$!39AF"}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestLoadUserOrg(t *testing.T) {

	dir := fixturePath(t, "userorg")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "XYZCorp", User: "billyjoe", Dir: "", Token: ""}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestLoadUserDir(t *testing.T) {

	dir := fixturePath(t, "userdir")
	os.Setenv(ConfigPathEnvVar, dir)

	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c == nil {
		t.Fatalf("fail: expecting a %s config", "non-nil")
	}

	expected := &Config{Org: "", User: "billyjoe", Dir: "/home/billyjoe/projects", Token: ""}
	if !reflect.DeepEqual(c, expected) {
		t.Fatalf("bad: %#v", c)
	}
}

func TestWriteNilConfig(t *testing.T) {

	if err := WriteConfig(nil); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	}
}

func TestWriteEmptyConfig(t *testing.T) {

	expectedDir := "dir is required"
	expectedUser := "user is required"
	expectedToken := "oauth token is required"

	c := &Config{}
	if err := WriteConfig(c); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expectedDir) {
		t.Fatalf("fail: error should include '%s'", expectedDir)
	} else if !strings.Contains(err.Error(), expectedUser) {
		t.Fatalf("fail: error should include '%s'", expectedUser)
	} else if !strings.Contains(err.Error(), expectedToken) {
		t.Fatalf("fail: error should include '%s'", expectedToken)
	}
}

func TestWriteBadDir(t *testing.T) {
	expectedDir := "dir '/tmpp' does not exist"
	c := &Config{User: "bob", Dir: "/tmpp", Token: "DKFK"}

	if err := WriteConfig(c); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expectedDir) {
		t.Fatalf("fail: error should include '%s'", expectedDir)
	}
}

func TestBadToken(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "L2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := "unable to authenticate against GitHub using the oauth token provided"

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expected) {
		t.Fatalf("fail: error should include '%s'", expected)
	}
}

func TestBadUserName(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		User:  "xxxxyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := "username provided does not match GitHub login or email associated with your authenicated user"

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expected) {
		t.Fatalf("fail: error should include '%s'", expected)
	}
}

func TestBadOrg(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		Org:   "XXXXXXXXXYZCorp",
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := "org does not exist on GitHub"

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expected) {
		t.Fatalf("fail: error should include '%s'", expected)
	}
}

func TestNotOrgMember(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		Org:   "XYZCorp",
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := "user is not a member of the organization provided"

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), expected) {
		t.Fatalf("fail: error should include '%s'", expected)
	}
}

func TestWritGood(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		Org:   "xyzinc",
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := &Config{Org: "xyzinc", User: "xyzincuser", Dir: "/tmp", Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317"}

	if err := WriteConfig(c); err != nil {
		t.Fatalf("fail: %s", err.Error())
	}

	out, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("bad: %#v", out)
	}
}

func TestWritNoOrg(t *testing.T) {

	dir := fixturePath(t, "")
	os.Setenv(ConfigPathEnvVar, dir)

	c := &Config{
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	expected := &Config{User: "xyzincuser", Dir: "/tmp", Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317"}

	if err := WriteConfig(c); err != nil {
		t.Fatalf("fail: %s", err.Error())
	}

	out, err := LoadConfig()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("bad: %#v", out)
	}
}
