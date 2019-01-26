# ghcli - A GitHub Cli

## About
CLI + GitHub = 20% Automation + 10% Productivity

In spirit of CI/CD, automation and version control system is the heart of things. I need a headless way of automating a VCS related actions, in this case is GitHub.

This cli will serve as a headless and non-programmatic way of performing authenticated GitHub actions via REST API adhering the official [guideline](https://developer.github.com/v3/) 

I am very well aware of the official GitHub's cli, [hub](https://github.com/github/hub). However, the commands offered in hub is limited, so I need to write my own GitHub CLI to perform these specialized tasks.

## TODO
This CLI will provide the following action items. These are tasks that I need, will be added as I require more:
[x] Get GitHub Repository (Public, Private, Organization)
[x] Creation of GitHub Repository (Public, Private, Organization)
[] Teams Listing
[] Adding members to a GitHub Repository
[] Adding a GitHub Repository to a team or vice versa

Add unit tests + mock instead of integration tests

## Thanks
Powered by Cobra and Go-GitHub libraries
