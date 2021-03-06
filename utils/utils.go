package utils

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/pinterb/hsc/config"
)

// Utils manages interactions with all utilities
type Utils struct {
	client *github.Client
	config *config.Config

	Users *UserUtils
	//	Teams         *TeamUtils
	//	Projects      *ProjectUtils
	//	Organizations *OrganizationUtils
}

// Response is a light wrapper around the go-github Response struct
type Response struct {
	*github.Response
}

// NewUtils creates an instance of Utils
func NewUtils(config *config.Config) *Utils {

	client := github.NewClient(nil)
	if config != nil {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: config.Token},
		}

		client = github.NewClient(t.Client())
	}

	u := &Utils{config: config, client: client}
	u.Users = &UserUtils{Utils: u}

	return u
}

// NewResponse creates a new instance of Response
func NewResponse(r *github.Response) *Response {
	resp := &Response{Response: r}
	return resp
}
