package utils

import "github.com/google/go-github/github"

// OrgUtils handles activities around GitHub organizations.
// This code is basically just a wrapper around the library that google has graciously
// put together.
//
// TODO: Methods implemented on this struct should provide generic return values (Not, the go-github structs)
type OrgUtils struct {
	*Utils
}

// List the organizations for a user.  Passing the empty string will list
// organizations for the authenticated user.
func (ot *OrgUtils) List(user string, opt *github.ListOptions) ([]github.Organization, *github.Response, error) {

	orgout, resp, err := ot.client.Organizations.List(user, opt)
	if err != nil {
		return nil, resp, err
	}

	return orgout, resp, nil
}

// Get fetches an organization by name.
func (ot *OrgUtils) Get(org string) (*github.Organization, *github.Response, error) {

	orgout, resp, err := ot.client.Organizations.Get(org)
	if err != nil {
		return nil, resp, err
	}

	return orgout, resp, nil
}

// Edit an organization.
func (ot *OrgUtils) Edit(name string, org *github.Organization) (*github.Organization, *github.Response, error) {

	orgout, resp, err := ot.client.Organizations.Edit(name, org)
	if err != nil {
		return nil, resp, err
	}

	return orgout, resp, nil

}

// ListMembers lists the members for an organization.  If the authenticated
// user is an owner of the organization, this will return both concealed and
// public members, otherwise it will only return public members.
func (ot *OrgUtils) ListMembers(org string, opt *github.ListMembersOptions) ([]github.User, *github.Response, error) {

	users, resp, err := ot.client.Organizations.ListMembers(org, opt)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// IsMember checks if a user is a member of an organization.
func (ot *OrgUtils) IsMember(org, user string) (bool, *github.Response, error) {

	feh, resp, err := ot.client.Organizations.IsMember(org, user)
	if err != nil {
		return false, resp, err
	}

	return feh, resp, nil
}

// IsPublicMember checks if a user is a public member of an organization.
func (ot *OrgUtils) IsPublicMember(org, user string) (bool, *github.Response, error) {

	feh, resp, err := ot.client.Organizations.IsPublicMember(org, user)
	if err != nil {
		return false, resp, err
	}

	return feh, resp, nil
}

// RemoveMember removes a user from all teams of an organization.
func (ot *OrgUtils) RemoveMember(org, user string) (*github.Response, error) {

	resp, err := ot.client.Organizations.RemoveMember(org, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// PublicizeMembership publicizes a user's membership in an organization.
func (ot *OrgUtils) PublicizeMembership(org, user string) (*github.Response, error) {

	resp, err := ot.client.Organizations.PublicizeMembership(org, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ConcealMembership conceals a user's membership in an organization.
func (ot *OrgUtils) ConcealMembership(org, user string) (*github.Response, error) {

	resp, err := ot.client.Organizations.ConcealMembership(org, user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListOrgMemberships lists the organization memberships for the authenticated user.
func (ot *OrgUtils) ListOrgMemberships(opt *github.ListOrgMembershipsOptions) ([]github.Membership, *github.Response, error) {

	membership, resp, err := ot.client.Organizations.ListOrgMemberships(opt)
	if err != nil {
		return nil, resp, err
	}

	return membership, resp, nil
}

// GetOrgMembership gets the membership for the authenticated user for the
// specified organization.
func (ot *OrgUtils) GetOrgMembership(org string) (*github.Membership, *github.Response, error) {

	membership, resp, err := ot.client.Organizations.GetOrgMembership(org)
	if err != nil {
		return nil, resp, err
	}

	return membership, resp, nil
}

// EditOrgMembership edits the membership for the authenticated user for the
// specified organization.
func (ot *OrgUtils) EditOrgMembership(org string, membership *github.Membership) (*github.Membership, *github.Response, error) {

	membershp, resp, err := ot.client.Organizations.EditOrgMembership(org, membership)
	if err != nil {
		return nil, resp, err
	}

	return membershp, resp, nil
}
