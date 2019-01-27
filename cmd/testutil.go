package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func createNewCommand() *cobra.Command {
	return CreateGhcliCommand()
}

// CliTestCase is used to store args for cli test
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
