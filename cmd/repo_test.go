package cmd

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// TODO: mock things instead of integration

type MockedRepoObject struct {
	mock.Mock
}

func TestRepo(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"repo", "--help"}, false, false}, // help
	}

	testCliCommand(t, testCases)
}

func TestCreateRepo_Fail(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"repo", "create", "--help"}, false, false}, // help
		{[]string{"repo", "create", "--name", "hello2"}, true, true},
		{[]string{"repo", "create", "--name", "hello", "--owner", "someone"}, true, true}, // does not exist but we've got no permission
	}

	testCliCommand(t, testCases)
}

func TestCreateRepo_Success(t *testing.T) {
	t.Skip()
	// Mock these results
	testCases := []CliTestCase{
		// {[]string{"repo", "create", "--name", "ghcli", "--owner", "deka108"}, true, false},        // Repo already exist
		// {[]string{"repo", "create", "--name", "new-repo-cli", "--owner", "deka108"}, false, true}, // Repo not exist, will be created
	}
	testCliCommand(t, testCases)
}

func TestGetRepo(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"repo", "get", "--help"}, false, false},                                   // help
		{[]string{"repo", "get", "--name", "my-project"}, true, true},                       // Owner is unset
		{[]string{"repo", "get", "--name", "notexist", "--owner", "anybody"}, false, false}, // Repo doesn't exist
	}
	testCliCommand(t, testCases)
}

func TestGetRepo_Success(t *testing.T) {
	t.Skip()
	testCases := []CliTestCase{
		{[]string{"repo", "get", "--name", "ghcli", "--owner", "deka108"}, false, false}, // Repo exist
	}
	testCliCommand(t, testCases)
}
