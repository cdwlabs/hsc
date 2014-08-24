# HSC

HSC (Hack Some Code) is an opinionated command line utility designed to get you writing code....quickly!

## Why?
If you are new to software development, the whole process of choosing a lanuage, framework, design pattern, tooling, etc. can seem daunting.  Solving that problem or implementing that new idea quickly becomes unapproachable.

And even if you're an experienced engineer, do you really want to spend an hour writing boilerplate before you can implement that proof-of-concept?

## Background 
The original premise was pretty simple.  I wanted to create a collection of small starter projects that engineers could use to quickly implement an idea or solve a problem.  With the technologies currently available (e.g. [Vagrant](http://www.vagrantup.com/), [VirtualBox](https://www.virtualbox.org/), and [Docker](https://www.docker.com/)), I felt it was feasible to create a developer environment that was consistent across operating systems, [polyglot programmer-friendly](http://radar.oreilly.com/2013/11/polyglot-programming-what-is-it-and-why-should-you-be-using-it.html), and not burdensome to use.

After getting a couple of starter projects created, I next set out to start writing some documentation around how to use them.  Writing good technical documentation has it's own unique set of challenges (which I won't get into here). As I started organizing my thoughts around what to write, I asked myself if writing a bunch of documentation was really worth the effort?  I mean, regardless of my efforts in writing quality documentation, there would be a segment of my target audience that would either not understand what I wrote or would not even bother reading it. 

So it occured to me that maybe I should consider a different approach. What if I could automate a large portion of the workflow?  How much documentation would I need to write if the prescribed workflow was baked into a small, easily accessible app?

Let the experiment begin...

## Concepts
* The workflow is very much centered around using [Git](http://git-scm.com/) and at least initially, GitHub.  So much of what HSC sets out to accomplish is around automating tasks around managing GitHub repositories, their issues, teams, collaborators, etc. etc.   
* HSC was conceived to support engineers collaborating within a [GitHub Organization](https://github.com/blog/674-introducing-organizations).  But since the GitHub [API](https://developer.github.com/v3/) obviously supports both individual and organization account types, then HSC should as well.
* Even if I could automate everything you need to do with Git, I'm not sure I'd want to.  So bottom line, you need a basic understanding of version control software and in particular, Git.  Quite frankly I don't think that's too much to ask.   

## Usage
HSC is written in Go and the plan is to cross-compile to run on Linux, Windows, and OS X.  As a command line application, its workflow is implemented as a collection of sub-commands:     

**init**: First time installing HSC   
**new**: Start a new project   
**fork**: Collaborate on an existing project   
**issues**: Manage project ideas, features, stories, and bugs   
**version**: Display version of your local HSC install   
**check**: Check for available updates for HSC   
**download**: Updates to latest HSC version   
**help**: Show this help

#### Sub-Commands (In Depth)
Let's explore each of the sub-commands.   

##### Init
This sub-command initializes a user's local installation with necessary configuration to streamline processing.  Storing a configuation locally will also eliminate the need to query for input that may be redundant across sub-commands.  GitHub repository information would be one example of the type of data that may be redundant across sub-commands.

##### New
This sub-command creates a new project for a user.  The user would be able to select from a list of pre-defined starter projects ideally implemented in multiple languages.  A good example would be the customary "hello world" application implemented as a [microservice](http://martinfowler.com/articles/microservices.html).  

##### Fork
If you find an existing project you'd like to contribute to, this is the sub-command for you.  

##### Issues
If you have a bug, enhancement, feature, etc. that you want to create for an existing project and don't feel like going to GitHub.  Additionally, this sub-command will allow you to "import" feature files as issues in GitHub. 

##### Version 
Display version of your local HSC install.  This includes both the semantic version number as well as Git information.

##### Check
Automagically checks the ether for a new version of this utility.

##### Download
Updates to latest HSC version.

##### Help
Display some help message.

### Status
No real work has been completed on this yet.  Just identifiying high-level functionality.

### To-Do's

##### init
* Decide on file format for config (e.g. yaml, toml, json, ini).
* Determine what to capture during init.

##### new
* Select from available starter projects.
* Provide option for creating repository in GitHub.

##### fork
* Fork existing GitHub repository.
* Clone locally and set up upstream remote for consistent fetch/merge workflow.

##### issues
* Create a generic GitHub issue (with label support).
* Import *.feature files.  Includes naming convention to prevent duplicates.
* Evaluate ideas around git hook support.

##### version 
* Determine what to expose as version info.

##### check
* Develop scheme to simplify version checking.
* Leverage Bintray API?

##### download
* Just an extension of "check" sub-command (ie call check first).
* Develop a scheme to determine if backwards compatability is broken.

##### help
* Should be easy...right?
