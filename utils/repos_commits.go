package utils

import "github.com/google/go-github/github"

// ListCommits lists the commits of a repository.
func (rt *RepoUtils) ListCommits(owner, repo string, opt *github.CommitsListOptions) ([]github.RepositoryCommit, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListCommits(owner, repo, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// GetCommit fetches the specified commit, including all details about it.
// todo: support media formats - https://github.com/google/go-github/issues/6
func (rt *RepoUtils) GetCommit(owner, repo, sha string) (*github.RepositoryCommit, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.GetCommit(owner, repo, sha)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// CompareCommits compares a range of commits with each other.
// todo: support media formats - https://github.com/google/go-github/issues/6
func (rt *RepoUtils) CompareCommits(owner, repo string, base, head string) (*github.CommitsComparison, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.CompareCommits(owner, repo, base, head)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}
