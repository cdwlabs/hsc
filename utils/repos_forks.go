package utils

import "github.com/google/go-github/github"


func (rt *RepoUtils) ListForks(owner, repo string, opt *github.RepositoryListForksOptions) ([]github.Repository, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListForks(owner, repo, opt)
	if err != nil {
		return nil, resp, err
	}

	return repoinfo, resp, nil

}



func (rt *RepoUtils) ??? {




}




func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}



func (rt *RepoUtils) ??? {




}


func (rt *RepoUtils) ??? {




}



