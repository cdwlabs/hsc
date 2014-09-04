package utils

import "github.com/google/go-github/github"

// GetReadme gets the Readme file for the repository.
func (rt *RepoUtils) GetReadme(owner, repo string, opt *github.RepositoryContentGetOptions) (*github.RepositoryContent, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.GetReadme(owner, repo, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// GetContents can return either the metadata and content of a single file
// (when path references a file) or the metadata of all the files and/or
// subdirectories of a directory (when path references a directory). To make it
// easy to distinguish between both result types and to mimic the API as much
// as possible, both result types will be returned but only one will contain a
// value and the other will be nil.
func (rt *RepoUtils) GetContents(owner, repo, path string, opt *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent,
	directoryContent []*github.RepositoryContent, resp *Response, err error) {

	fileinfo, dirinfo, respo, err := rt.client.Repositories.GetContents(owner, repo, path, opt)
	respinfo := NewResponse(respo)
	return fileinfo, dirinfo, respinfo, err

}

// CreateFile creates a new file in a repository at the given path and returns
// the commit and file metadata.
func (rt *RepoUtils) CreateFile(owner, repo, path string, opt *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.CreateFile(owner, repo, path, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// UpdateFile updates a file in a repository at the given path and returns the
// commit and file metadata. Requires the blob SHA of the file being updated.
func (rt *RepoUtils) UpdateFile(owner, repo, path string, opt *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.UpdateFile(owner, repo, path, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}

// DeleteFile deletes a file from a repository and returns the commit.
// Requires the blob SHA of the file to be deleted.
func (rt *RepoUtils) DeleteFile(owner, repo, path string, opt *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.DeleteFile(owner, repo, path, opt)
	respinfo := NewResponse(resp)
	return repoinfo, respinfo, err

}
