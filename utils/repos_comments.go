package utils

import "github.com/google/go-github/github"

// ListComments lists all the comments for the repository.
func (rt *RepoUtils) ListComments(owner, repo string, opt *github.ListOptions) ([]github.RepositoryComment, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListComments(owner, repo, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// ListCommitComments lists all the comments for a given commit SHA.
func (rt *RepoUtils) ListCommitComments(owner, repo, sha string, opt *github.ListOptions) ([]github.RepositoryComment, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListCommitComments(owner, repo, sha, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// CreateComment creates a comment for the given commit.
// Note: GitHub allows for comments to be created for non-existing files and positions.
func (rt *RepoUtils) CreateComment(owner, repo, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.CreateComment(owner, repo, sha, comment)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// GetComment gets a single comment from a repository.
func (rt *RepoUtils) GetComment(owner, repo string, id int) (*github.RepositoryComment, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.GetComment(owner, repo, id)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// UpdateComment updates the body of a single comment.
func (rt *RepoUtils) UpdateComment(owner, repo string, id int, comment *github.RepositoryComment) (*github.RepositoryComment, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.UpdateComment(owner, repo, id, comment)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

func (rt *RepoUtils) DeleteComment(owner, repo string, id int) (*Response, error) {

	resp, err := rt.client.Repositories.DeleteComment(owner, repo, id)
	respinfo := NewResponse(resp)
	return respinfo, err

}
