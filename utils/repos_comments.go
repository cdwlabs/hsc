package utils

import "github.com/google/go-github/github"

// ListComments lists all the comments for the repository.
func (rt *RepoUtils) ListComments(owner, repo string, opt *github.ListOptions) ([]github.RepositoryComment, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.ListComments(owner, repo, opt)
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



