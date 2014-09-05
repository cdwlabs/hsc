## Workflow Concepts

### New User

#### GitHub Notes
To add someone as a member to an org, you must add them to a team.  

In order to add a membership between a user and a team, the authenticated user must have ‘admin’ permissions to the team or be an owner of the organization that the team is associated with.   

If the user is already a part of the team’s organization (meaning they’re on at least one other team in the organization), this endpoint will add the user to the team.   

If the user is completely unaffiliated with the team’s organization (meaning they’re on none of the organization’s teams), this endpoint will send an invitation to the user via email. This newly-created membership will be in the “pending” state until the user accepts the invitation, at which point the membership will transition to the “active” state and the user will be added as a member of the team.   

#### HSC
As the user initializes their installation of HSC, the system should:     
* Assuming all data submitted is valid; create a .hsc config file   
* Check if the user is a member of the organization.      
* If the user is a current member of the organization; then there's nothing else the system needs to do.    
* If the user is NOT a current member of the organization; then the system sends a message to hsc message queue.   
* The message contains the user's GitHub username, the GitHub organization they wish to join, and the date of their request.   
* Once the message enters the queue, the system will create a GitHub issue in the organizations admin, issues-only repository.   
* The GitHub organization's owner will be notified that someone wants to join their organization.   
* If the owner declines the user's request, they should contact the user about their reasons for declining request and then close the issue in GitHub.   
* Accepted requests are completed by adding the user to the organization's default team (typically a team with the same name as the organization).  GitHub will send an invitation to the user to join the team. The organization owner will close the ticket and all is good.   


### Default Team
All users to the GitHub organization need to belong to at least one GitHub Team.  The HSC workflow, by convention, will put all HSC users into the team that has the same name as the organization name (e.g. cdwlabs).  As of this writing, this default team will have admin permissions.  This allows team members to create (and delete) their own respositories.  The potential risk is that a user deletes a repository they have no business messing around with.  For this reason, the general guideline is to rarely (if ever) delete a repository.  Additionally, the HSC workflow should encourage the creation of other GitHub Teams that are more domain focused.  These domain-specific Teams are curated by "team leads" and most users will only have read (ie pull) permissions.  See New Teams below.


### New Teams

#### GitHub Notes
In order to create a team, the authenticated user must be an owner of org.   

#### HSC
The purpose of additional teams from a HSC perspective is to support a "Team Lead" or "Committer" role. Given the nature of these roles, the number of members in these teams should be limited.  The goal should be only enough members to get the job done...which is primarily that of managing pull requests. With this context, when a HSC user requests a new team it is for the sole purpose of becoming the team lead of the new team.    

Assuming the user is already a member of the organization, when they submit a request to create a team, the system should:    
* Check that the team doesn't already exist.    
* If the team doesn't exist; then the system sends a message to hsc message queue.   
* The message contains the user's GitHub username, the GitHub organization, the team name, and the date of their request.   
* Once the message enters the queue, the system will create a GitHub issue in the organizations admin, issues-only repository.   
* The GitHub organization's owner will be notified that someone wants to create a team.   
* If the owner declines the user's request, they should contact the user about their reasons for declining request and then close the issue in GitHub.   
* Accepted requests are completed by creating the new team (with admin permissions) and adding the user to the new team.  GitHub should notify the user of the team's creation.   The organization owner will closs the ticket and all is good.   


#### New Project

#### GitHub Notes
Create a new repository in this organization. The authenticated user must be a member of the specified organization (if not org member, repo is public personal).   

### HSC
By default, new projects will get created in the organization's default team.  But the user should be able to specify which team they would like to create the repo in.  Once the repo is created, a new branch called "initialdev" is also created.  After the branch has been created, the user will be given the example command for cloning the branch of the new project.

git clone -b initialdev https://github.com/cdwlabs/yippee.git


#### Move Repository

#### GitHub Notes
In order to add a repository to a team, the authenticated user must be an owner of the org that the team is associated with. Also, the repository must be owned by the organization, or a direct fork of a repository owned by the organization.     

### HSC
* If Team Lead, automatically move repo to team. Team Lead decides if repo should become private or remain public.
* If non-Team member, request that repo be moved.  Request goes to Team Lead and Org Owner (or created as an issue in A1 repo).
