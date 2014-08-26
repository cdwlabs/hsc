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

	c := &Config{}
	if err := WriteConfig(c); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrMissingDir) {
		t.Fatalf("fail: error should include '%s'", ErrMissingDir)
	} else if !strings.Contains(err.Error(), ErrMissingUser) {
		t.Fatalf("fail: error should include '%s'", ErrMissingUser)
	} else if !strings.Contains(err.Error(), ErrMissingToken) {
		t.Fatalf("fail: error should include '%s'", ErrMissingToken)
	}
}

func TestWriteBadDir(t *testing.T) {
	c := &Config{User: "bob", Dir: "/tmpp", Token: "DKFK"}

	if err := WriteConfig(c); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrDirDoesNotExist) {
		t.Fatalf("fail: error should include '%s'", ErrDirDoesNotExist)
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

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrAuthUser) {
		t.Fatalf("fail: error should include '%s'", ErrAuthUser)
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

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrUserMismatch) {
		t.Fatalf("fail: error should include '%s'", ErrUserMismatch)
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

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrOrgDoesNotExist) {
		t.Fatalf("fail: error should include '%s'", ErrOrgDoesNotExist)
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

	if err := c.Validate(); err == nil {
		t.Fatalf("fail: expecting a %s error", "non-nil")
	} else if !strings.Contains(err.Error(), ErrUserNotOrgMember) {
		t.Fatalf("fail: error should include '%s'", ErrUserNotOrgMember)
	}
}

/*
To Do: Determine if Config should allow an Org even if User is NOT a member of org...

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
*/

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
