package utils

import "github.com/google/go-github/github"

// GetReadme gets the Readme file for the repository.
func (rt *RepoUtils) GetReadme(owner, repo string, opt *github.RepositoryContentGetOptions) (*github.RepositoryContent, *github.Response, error) {

	repoinfo, resp, err := rt.client.Repositories.GetReadme(owner, repo, opt)
	if err != nil {
		return nil, resp, err
	}

	return repoinfo, resp, nil

}


// Decode decodes the file content if it is base64 encoded.
func (rt *RepoUtils) Decode() ([]byte, error) {




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



