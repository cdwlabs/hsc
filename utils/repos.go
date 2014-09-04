package utils

import "github.com/google/go-github/github"

// RepoUtils handles activities around GitHub repositories.
// This code is basically just a wrapper around the library that google has graciously
// put together.
//
// TODO: Methods implemented on this struct should provide generic return values (Not, the go-github structs)
type RepoUtils struct {
	*Utils
}

// List the repositories for a user.  Passing the empty string will list
// repositories for the authenticated user.
func (rt *RepoUtils) List(user string, opt *github.RepositoryListOptions) ([]github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.List(user, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListByOrg lists the repositories for an organization.
func (rt *RepoUtils) ListByOrg(org string, opt *github.RepositoryListByOrgOptions) ([]github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListByOrg(org, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListAll lists all GitHub repositories in the order that they were created.
func (rt *RepoUtils) ListAll(opt *github.RepositoryListAllOptions) ([]github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListAll(opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// Create a new repository.  If an organization is specified, the new
// repository will be created under that org.  If the empty string is
// specified, it will be created for the authenticated user.
func (rt *RepoUtils) Create(org string, repo *github.Repository) (*github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.Create(org, repo)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// Get fetches a repository.
func (rt *RepoUtils) Get(owner, repo string) (*github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.Get(owner, repo)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// Edit updates a repository.
func (rt *RepoUtils) Edit(owner, repo string, repository *github.Repository) (*github.Repository, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.Edit(owner, repo, repository)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// Delete a repository.
func (rt *RepoUtils) Delete(owner, repo string) (*Response, error) {

	resp, err := rt.client.Repositories.Delete(owner, repo)
	if err != nil {
		respinfo := NewResponse(resp)
		return respinfo, err
	}

	respinfo := NewResponse(resp)
	return respinfo, nil

}

// ListContributors lists contributors for a repository
func (rt *RepoUtils) ListContributors(owner string, repository string, opt *github.ListContributorsOptions) ([]github.Contributor, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListContributors(owner, repository, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListLanguages lists languages for the specified repository. The returned map
// specifies the languages and the number of bytes of code written in that
// language. For example:
//
//     {
//       "C": 78769,
//       "Python": 7769
//     }
func (rt *RepoUtils) ListLanguages(owner string, repo string) (map[string]int, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListLanguages(owner, repo)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListTeams lists the teams for the specified repository.
func (rt *RepoUtils) ListTeams(owner string, repo string, opt *github.ListOptions) ([]github.Team, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListTeams(owner, repo, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListTags lists tags for the specified repository.
func (rt *RepoUtils) ListTags(owner string, repo string, opt *github.ListOptions) ([]github.RepositoryTag, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListTags(owner, repo, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// ListBranches lists branches for the specified repository.
func (rt *RepoUtils) ListBranches(owner string, repo string, opt *github.ListOptions) ([]github.Branch, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListBranches(owner, repo, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}

// GetBranch gets the specified branch for a repository.
func (rt *RepoUtils) GetBranch(owner, repo, branch string) (*github.Branch, *Response, error) {

	repoinfo, resp, err := rt.client.Repositories.GetBranch(owner, repo, branch)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return repoinfo, respinfo, nil

}
