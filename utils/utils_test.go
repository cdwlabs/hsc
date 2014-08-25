package utils

import (
	"testing"

	"github.com/pinterb/hsc/config"
)

func TestNilConfig(t *testing.T) {
	utils := NewUtils(nil)
	if utils == nil {
		t.Fatal("fail: Creating a Utils instance with nil Config should not fail."
	}
}

func TestUserClient(t *testing.T) {
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
