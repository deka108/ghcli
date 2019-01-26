package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TODO: mock things instead of integration

type MockedObject struct {
	mock.Mock
}

func createNewCommand() *cobra.Command {
	return CreateGhcliCommand()
}

type CliTestCase struct {
	args       []string
	isError    bool
	hideOutput bool
}

func testCliCommand(t *testing.T, testCases []CliTestCase) {
	for _, tc := range testCases {
		runCliCommand(t, tc)
	}
}

func runCliCommand(t *testing.T, tc CliTestCase) {
	rootCmd := createNewCommand()
	if tc.hideOutput {
		rootCmd.SetOutput(ioutil.Discard)
	}
	rootCmd.SetArgs(tc.args)
	if tc.isError {
		require.Error(t, rootCmd.Execute())
	} else {
		require.NoError(t, rootCmd.Execute())
	}
}

func TestCreateRepo(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"repo", "create", "--name", "hello2"}, true, true},
		{[]string{"repo", "create", "--name", "hello", "--owner", "ooo"}, false, true},
		{[]string{"repo", "create", "--name", "ghcli", "--owner", "deka108"}, true, false}, // Repo already exist
	}

	testCliCommand(t, testCases)
}

func TestCreateNewRepo(t *testing.T) {
	t.Skip()
	testCase := CliTestCase{[]string{"repo", "create", "--name", "new-repo-cli", "--owner", "deka108"}, false, true} // Mock this result
	runCliCommand(t, testCase)
}

func TestGetRepo(t *testing.T) {
	testCases := []CliTestCase{
		{[]string{"repo", "get", "--name", "my-project"}, true, true},                       // Owner is unset
		{[]string{"repo", "get", "--name", "notexist", "--owner", "anybody"}, false, false}, // Repo doesn't exist
		{[]string{"repo", "get", "--name", "ghcli", "--owner", "deka108"}, false, false},    // Repo exist
	}
	testCliCommand(t, testCases)
}
