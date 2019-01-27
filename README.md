# ghcli - A GitHub Cli

## About
CLI + GitHub = 20% Automation + 10% Productivity

In spirit of CI/CD, automation and version control system is the heart of things. I need a headless way of automating VCS related actions, which in my case is using GitHub.

This `ghcli` will serve as a headless and non-programmatic way of performing authenticated GitHub actions via REST API adhering to the official [guideline](https://developer.github.com/v3/) 

I am well aware of the official GitHub's cli, [hub](https://github.com/github/hub). However, the commands offered in hub is limited, so I need to write my own GitHub CLI to perform some of customized GitHub actions.

## Setup
To install, make sure you have go installed. Then, run `go get github.com/deka108/ghcli`. `ghcli` should now be available from your cli

## How to Use

`ghcli --help`

### Repositories
|   Action   |   Command  |
| --- | --- |
| Gets a repository | `ghcli repo get --help` |
| Creates a repository | `ghcli repo create --help` |


## TODO
This CLI currently will provide the following action items, these tasks will be added as I require more tasks:
- [x] Get GitHub Repository (Public, Private, Organization)
- [x] Creation of GitHub Repository (Public, Private, Organization)
- [ ] Teams Listing
- [ ] Adding members to a GitHub Repository
- [ ] Adding a GitHub Repository to a team or vice versa
- [ ] Use unit tests + mocking the unit tests instead of integration tests

## Thanks
Powered by Cobra and Go-GitHub
