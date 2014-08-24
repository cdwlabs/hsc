package utils

import "github.com/google/go-github/github"

// UserUtils handles activities around managing users in a development workflow.
// This could include things like:
// 1. communication with the user related methods of the GitHub API.
//    Note: GitHub API docs: http://developer.github.com/v3/users/
//
type UserUtils struct {
	*Utils
}

// GetGitHubUser fetches a repository user.  Passing the empty string will fetch the authenticated
// user.
// TODO: This should provide generic return values (Not, the go-github structs)
func (ut *UserUtils) GetGitHubUser(user string) (*github.User, *github.Response, error) {

	//u, resp, err := ut.client.client.Users.Get(user)
	u, resp, err := ut.client.Users.Get(user)
	if err != nil {
		return nil, resp, err
	}

	return u, resp, nil
}

// IsGitHubUser is a convenience method for determining if user is a valid GitHub user.
func (ut *UserUtils) IsGitHubUser(user string) (bool, error) {

	_, resp, err := ut.GetGitHubUser(user)
	if err != nil && resp.StatusCode != 404 {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}
