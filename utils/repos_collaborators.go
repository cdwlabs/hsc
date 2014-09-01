package utils

import "github.com/google/go-github/github"

// ListCollaborators lists the Github users that have access to the repository.
func (rt *RepoUtils) ListCollaborators(owner, repo string, opt *github.ListOptions) ([]github.User, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListCollaborators(owner, repo, opt)
	if err != nil {
		return nil, resp, err
	}

	return repoinfo, resp, nil

}

// IsCollaborator checks whether the specified Github user has collaborator
// access to the given repo.
// Note: This will return false if the user is not a collaborator OR the user
// is not a GitHub user.
func (rt *RepoUtils) IsCollaborator(owner, repo, user string) (bool, *github.Response, error) {




}


// AddCollaborator adds the specified Github user as collaborator to the given repo.
func (rt *RepoUtils) AddCollaborator(owner, repo, user string) (*github.Response, error) {




}


// RemoveCollaborator removes the specified Github user as collaborator from the given repo.
// Note: Does not return error if a valid user that is not a collaborator is removed.
func (rt *RepoUtils) RemoveCollaborator(owner, repo, user string) (*github.Response, error) {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



