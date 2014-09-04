                           Desktop          Github

1. Engineer requests to be added to org.
   |
   |__ Request is opened as issue in Org's GitHub "alpha" repo.
   
1b. GitHub Owner invites Engineer   

init
new
 -github (defaults to public repo for now)
 -name=""
 -description=""
 -team=""
 -type=microservice
 -language=perl
 
team
 - join
 - leave
 - create
 - status
 - members
 - repos
 - list
info
 - starterprojects

### Workflow

#### New User
* - as you initialize your install, check w/sys to see if you're member of organization
	if not; update config to indicate non-membership
	if so, update config to indicate membership

#### New Project
Create a new repository in this organization. The authenticated user must be a member of the specified organization (if not org member, repo is public personal).   

* GH should allow you to create a repo whether in you're an org member or not
  if an org member the repo defaults to public

#### Team Repository
In order to add a repository to a team, the authenticated user must be an owner of the org that the team is associated with. Also, the repository must be owned by the organization, or a direct fork of a repository owned by the organization.  

* Move repo to team
  If Team Lead, automatically move repo to team. Team Lead decides if repo should become private or remain public.
  If non-Team member, request that repo be moved.  Request goes to Team Lead and Org Owner (or created as an issue in A1 repo).

#### New Team
In order to create a team, the authenticated user must be an owner of org.

* Anyone can request a team be created.  The person requesting the new team becomes the Team Lead unless they nominate someone else. 
