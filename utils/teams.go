package utils

import "github.com/google/go-github/github"

// TeamUtils handles activities around managing teams within a GitHub organization.
// This code is basically just a wrapper around the library that google has graciously
// put together.
//
// TODO: Methods implemented on this struct should provide generic return values (Not, the go-github structs)
type TeamUtils struct {
	*Utils
}

// Team represents a team within a GitHub organization.  Teams are used to
// manage access to an organization's repositories.
//
// This is just a thin wrapper around the go-github Team struct
type Team struct {
	*github.Team
}

// NewTeam creates a new instance of Team
func NewTeam(t *github.Team) *Team {
	team := &Team{Team: t}
	return team
}

// ListTeams lists all of the teams for an organization.
func (tt *TeamUtils) ListTeams(org string, opt *github.ListOptions) ([]github.Team, *Response, error) {

	teamout, resp, err := tt.client.Organizations.ListTeams(org, opt)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return teamout, respinfo, nil
}

// GetTeam fetches a team by ID.
func (tt *TeamUtils) GetTeam(team int) (*github.Team, *Response, error) {

	teamout, resp, err := tt.client.Organizations.GetTeam(team)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return teamout, respinfo, nil
}

// CreateTeam creates a new team within an organization.
func (tt *TeamUtils) CreateTeam(org string, team *github.Team) (*github.Team, *Response, error) {

	teamout, resp, err := tt.client.Organizations.CreateTeam(org, team)
	if err != nil {
		respinfo := NewResponse(resp)
		return nil, respinfo, err
	}

	respinfo := NewResponse(resp)
	return teamout, respinfo, nil
}

// EditTeam edits a team.
func (tt *TeamUtils) EditTeam(id int, team *github.Team) (*github.Team, *github.Response, error) {

	teamout, resp, err := tt.client.Organizations.EditTeam(id, team)
	if err != nil {
		return nil, resp, err
	}

	return teamout, resp, err
}

// DeleteTeam deletes a team.
func (tt *TeamUtils) DeleteTeam(team int) (*github.Response, error) {

	resp, err := tt.client.Organizations.DeleteTeam(team)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListTeamMembers lists all of the users who are members of the specified
// team.
func (tt *TeamUtils) ListTeamMembers(team int, opt *github.ListOptions) ([]github.User, *github.Response, error) {

	members, resp, err := tt.client.Organizations.ListTeamMembers(team, opt)
	if err != nil {
		return nil, resp, err
	}

	return members, resp, nil
}

// IsTeamMember checks if a user is a member of the specified team.
func (tt *TeamUtils) IsTeamMember(team int, user string) (bool, *github.Response, error) {

	feh, resp, err := tt.client.Organizations.IsTeamMember(team, user)
	if err != nil {
		return false, resp, err
	}

	return feh, resp, nil
}

// AddTeamMember adds a user to a team.
func (tt *TeamUtils) AddTeamMember(team int, user string) (*github.Response, error) {

	resp, err := tt.client.Organizations.AddTeamMember(team, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveTeamMember removes a user from a team.
func (tt *TeamUtils) RemoveTeamMember(team int, user string) (*github.Response, error) {

	resp, err := tt.client.Organizations.RemoveTeamMember(team, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListTeamRepos lists the repositories that the specified team has access to.
func (tt *TeamUtils) ListTeamRepos(team int, opt *github.ListOptions) ([]github.Repository, *github.Response, error) {

	repos, resp, err := tt.client.Organizations.ListTeamRepos(team, opt)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil
}

// IsTeamRepo checks if a team manages the specified repository.
func (tt *TeamUtils) IsTeamRepo(team int, owner string, repo string) (bool, *github.Response, error) {

	feh, resp, err := tt.client.Organizations.IsTeamRepo(team, owner, repo)
	if err != nil {
		return false, resp, err
	}

	return feh, resp, nil
}

// AddTeamRepo adds a repository to be managed by the specified team.  The
// specified repository must be owned by the organization to which the team
// belongs, or a direct fork of a repository owned by the organization.
func (tt *TeamUtils) AddTeamRepo(team int, owner string, repo string) (*github.Response, error) {

	resp, err := tt.client.Organizations.AddTeamRepo(team, owner, repo)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveTeamRepo removes a repository from being managed by the specified
// team.  Note that this does not delete the repository, it just removes it
// from the team.
func (tt *TeamUtils) RemoveTeamRepo(team int, owner string, repo string) (*github.Response, error) {

	resp, err := tt.client.Organizations.RemoveTeamRepo(team, owner, repo)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetTeamMembership returns the membership status for a user in a team.
func (tt *TeamUtils) GetTeamMembership(team int, user string) (*github.Membership, *github.Response, error) {

	mship, resp, err := tt.client.Organizations.GetTeamMembership(team, user)
	if err != nil {
		return nil, resp, err
	}

	return mship, resp, nil
}

// AddTeamMembership adds or invites a user to a team.
//
// In order to add a membership between a user and a team, the authenticated
// user must have 'admin' permissions to the team or be an owner of the
// organization that the team is associated with.
//
// If the user is already a part of the team's organization (meaning they're on
// at least one other team in the organization), this endpoint will add the
// user to the team.
//
// If the user is completely unaffiliated with the team's organization (meaning
// they're on none of the organization's teams), this endpoint will send an
// invitation to the user via email. This newly-created membership will be in
// the "pending" state until the user accepts the invitation, at which point
// the membership will transition to the "active" state and the user will be
// added as a member of the team.
func (tt *TeamUtils) AddTeamMembership(team int, user string) (*github.Membership, *github.Response, error) {

	mship, resp, err := tt.client.Organizations.AddTeamMembership(team, user)
	if err != nil {
		return nil, resp, err
	}

	return mship, resp, nil
}

// RemoveTeamMembership removes a user from a team.
func (tt *TeamUtils) RemoveTeamMembership(team int, user string) (*github.Response, error) {

	resp, err := tt.client.Organizations.RemoveTeamMembership(team, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
