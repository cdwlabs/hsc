package utils

import "github.com/google/go-github/github"

// ListForks lists the forks of the specified repository.
func (rt *RepoUtils) ListForks(owner, repo string, opt *github.RepositoryListForksOptions) ([]github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListForks(owner, repo, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// CreateFork creates a fork of the specified repository.
func (rt *RepoUtils) CreateFork(owner, repo string, opt *github.RepositoryCreateForkOptions) (*github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.CreateFork(owner, repo, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}
