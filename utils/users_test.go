package utils

import (
	"strings"
	"testing"

	"github.com/pinterb/hsc/config"
)

func TestBadCredentials(t *testing.T) {
	c := &config.Config{
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "DIEDSDKWIDS83F",
	}

	baduser := "bbbcccdddeeefffgggg"
	utils := NewUtils(c)
	_, err := utils.Users.IsGitHubUser(baduser)
	if err != nil {
		if !strings.Contains(err.Error(), "401 Bad credentials") {
			t.Fatalf("err: %v", err)
		}
	} else {
		t.Fatal("fail: expecting a 'bad credentials' error")
	}
}

func TestUndefinedUser(t *testing.T) {
	c := &config.Config{
		User:  "cuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	baduser := "bbbcccdddeeefffgggg"
	utils := NewUtils(c)
	ok, err := utils.Users.IsGitHubUser(baduser)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if ok {
		t.Fatalf("fail: not expecting user '%s' to exist on GitHub", baduser)
	}

}

func TestDefinedUser(t *testing.T) {
	c := &config.Config{
		User:  "cuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	validuser := "pinterb"
	utils := NewUtils(c)
	ok, err := utils.Users.IsGitHubUser(validuser)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !ok {
		t.Fatalf("fail: expecting user '%s' to exist on GitHub", validuser)
	}
}

func TestGetDefinedUser(t *testing.T) {
	c := &config.Config{
		User:  "xyzincuser",
		Dir:   "/tmp",
		Token: "2167c2a84fa7e09d4304aa005f6cb5e51f93d317",
	}

	utils := NewUtils(c)
	user, _, err := utils.Users.GetGitHubUser(c.User)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := "xyzincuser"
	if *user.Login != expected {
		t.Fatalf("fail: expecting user login '%s' and instead got '%v'", expected, *user.Login)
	}
}
