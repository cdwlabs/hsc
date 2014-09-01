package utils

import "github.com/google/go-github/github"

func (rt *RepoUtils) ListCommits(owner, repo string, opt *github.CommitsListOptions) ([]github.RepositoryCommit, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListCommits(owner, repo, opt)
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



