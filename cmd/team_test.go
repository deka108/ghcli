package cmd

import (
	"testing"
)

func TestTeam(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"team", "--help"}, false, false}, // help
	}

	testCliCommand(t, testCases)
}

func TestListTeams(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"team", "list", "--help"}, false, false}, // help
		{[]string{"team", "list"}, true, true},             // No org flag
	}

	testCliCommand(t, testCases)
}

func TestListTeams_Success(t *testing.T) {
	t.Skip()
	testCases := []CliTestCase{
		{[]string{"team", "list", "--org", "real-org"}, false, false}, // mock this
	}

	testCliCommand(t, testCases)
}

func TestGetTeamFromName(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"team", "getTeamFromName", "--help"}, false, false},            // help
		{[]string{"team", "getTeamFromName", "--name", "repo-name"}, true, true}, // No org flag
		{[]string{"team", "getTeamFromName", "--org", "some-org"}, true, true},   // No name flag
	}

	testCliCommand(t, testCases)
}

func TestGetTeamFromName_Success(t *testing.T) {
	t.Skip()
	testCases := []CliTestCase{
		{[]string{"team", "getTeamFromName", "--org", "realteam", "--name", "team-doesnt-exist"}, false, false}, // mock this
		// {[]string{"team", "getTeamFromName", "--org", "realteam", "--name", "team-exist"}, false, false},    // mock this
	}

	testCliCommand(t, testCases)
}

func TestGetTeamFromId(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"team", "getTeamFromId", "--help"}, false, false}, // help
		{[]string{"team", "getTeamFromId"}, true, true},             // No org flag
	}

	testCliCommand(t, testCases)
}

func TestGetTeamFromId_Success(t *testing.T) {
	t.Skip()
	testCases := []CliTestCase{
		{[]string{"team", "getTeamFromId", "--id", "-1"}, true, false},    // mock this
		{[]string{"team", "getTeamFromId", "--id", "1234"}, false, false}, // mock this
	}

	testCliCommand(t, testCases)
}

func TestAddTeamRepo(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"team", "addTeamRepo", "--help"}, false, false}, // help
		{[]string{"team", "addTeamRepo"}, true, true},             // No id, owner, and repo flag
	}

	testCliCommand(t, testCases)
}

func TestAddTeamRepo_Success(t *testing.T) {
	t.Skip()
	testCases := []CliTestCase{
		{[]string{"team", "addTeamRepo", "--id", "teamId", "--owner", "real-owner",
			"--repo", "real-repo", "--permission", "admin"}, false, false}, // mock this
	}

	testCliCommand(t, testCases)
}
