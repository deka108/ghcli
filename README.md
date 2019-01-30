# ghcli - A GitHub Cli

## About

CLI + GitHub = 20% Automation + 10% Productivity

In spirit of CI/CD, automation and version control system is the heart of things. I need a headless way of automating VCS related actions, which in my case is using GitHub.

This `ghcli` will serve as a headless and non-programmatic way of performing authenticated GitHub actions via REST API adhering to the official [guideline](https://developer.github.com/v3/).

I am well aware of the official GitHub's cli, [hub](https://github.com/github/hub). However, the commands offered in hub is limited, so I need to write my own GitHub CLI to perform some of customized GitHub actions.

## Setup

To install

### If you have Go installed

Run `go get github.com/deka108/ghcli`. `ghcli` should now will be available from your cli

### If you have Docker installed

Run `docker run -it deka108/ghcli:latest [commands]`.
To see the available commands, you can run `docker run -it deka108/ghcli:latest --help`

## How to Use

`ghcli --help`

### Repositories

|   Action   |   Command  |
| --- | --- |
| Gets a repository | `ghcli repo get --help` |
| Creates a repository | `ghcli repo create --help` |

### Teams

|   Action   |   Command  |
| --- | --- |
| List available org's teams | `ghcli team list --help` |
| Gets a team from an organization based on team's name | `ghcli team getTeamFromName --help` |
| Gets a team from an team's ID | `ghcli team getTeamFromId --help` |

## TODO

This CLI will provide the following action items:

- [x] Get GitHub Repository (Public, Private, Organization)
- [x] Creation of GitHub Repository (Public, Private, Organization)
- [x] Teams Listing, Gets a Team from Name
- [x] Adding members to a GitHub Repository
- [x] Adding a GitHub Repository to a team or vice versa
- [x] Use tests
- [ ] Mocking the unit tests instead of integration tests

These are tasks that I need, will be added as I require more

## Thanks

Powered by Cobra and Go-GitHub
