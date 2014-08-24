package utils

import (
	"net/http"

	"code.google.com/p/goauth2/oauth"
	"github.com/pinterb/hsc/config"
)

// Client manages interactions with all utilities
type Client struct {
	//github *github.Client
	config *config.Config
	client *http.Client

	Users *UserUtils
	//	Teams         *TeamUtils
	//	Projects      *ProjectUtils
	//	Organizations *OrganizationUtils
}

func NewClient(config *config.Config) *Client {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: config.Token},
	}

	httpClient := t.Client()
	c := &Client{config: config, client: httpClient}
	c.Users = &UserUtils{config: config}

	return c
}
