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
func (rt *RepoUtils) List(user string, opt *github.RepositoryListOptions) ([]github.Repository, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.List(user, opt)
	if err != nil {
		return nil, resp, err
	}

	return repoinfo, resp, nil

}


// ListByOrg lists the repositories for an organization.
func (rt *RepoUtils) ListByOrg(org string, opt *github.RepositoryListByOrgOptions) ([]github.Repository, *github.Response, error) {




}

// ListAll lists all GitHub repositories in the order that they were created.
func (rt *RepoUtils) ListAll(opt *github.RepositoryListAllOptions) ([]github.Repository, *github.Response, error) {




}

// Create a new repository.  If an organization is specified, the new
// repository will be created under that org.  If the empty string is
// specified, it will be created for the authenticated user.
func (rt *RepoUtils) Create(org string, repo *github.Repository) (*github.Repository, *github.Response, error) {




}

// Get fetches a repository.
func (rt *RepoUtils) Get(owner, repo string) (*github.Repository, *github.Response, error) {




}

// Edit updates a repository.
func (rt *RepoUtils) Edit(owner, repo string, repository *github.Repository) (*github.Repository, *github.Response, error) {




}

// Delete a repository.
func (rt *RepoUtils) Delete(owner, repo string) (*github.Response, error) {




}

// ListContributors lists contributors for a repository
func (rt *RepoUtils) ListContributors(owner string, repository string, opt *github.ListContributorsOptions) ([]github.Contributor, *github.Response, error) {




}

// ListLanguages lists languages for the specified repository. The returned map
// specifies the languages and the number of bytes of code written in that
// language. For example:
//
//     {
//       "C": 78769,
//       "Python": 7769
//     }
func (rt *RepoUtils) ListLanguages(owner string, repo string) (map[string]int, *github.Response, error) {




}


// ListTeams lists the teams for the specified repository.
func (rt *RepoUtils) ListTeams(owner string, repo string, opt *github.ListOptions) ([]Team, *github.Response, error) {




}


// ListTags lists tags for the specified repository.
func (rt *RepoUtils) ListTags(owner string, repo string, opt *github.ListOptions) ([]github.RepositoryTag, *github.Response, error) {




}


// ListBranches lists branches for the specified repository.
func (rt *RepoUtils) ListBranches(owner string, repo string, opt *github.ListOptions) ([]github.Branch, *github.Response, error) {




}

// GetBranch gets the specified branch for a repository.
func (rt *RepoUtils) GetBranch(owner, repo, branch string) (*github.Branch, *github.Response, error) {




}


func (rt *RepoUtils) ??? {




}
