// Copyright © 2019 Deka Auliya deka108@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/deka108/ghcli/ghutil"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// isAuth checks if the current cli operation is authenticated
func isAuth() bool {
	return os.Getenv("GITHUB_TOKEN") != "" || viper.GetString("token") != ""
}

// getAuthClient creates authenticated client from flag or env variable
func getAuthClient() (*github.Client, error) {
	if viper.GetString("token") != "" {
		return ghutil.NewAuthorizedClient(viper.GetString("token")), nil
	} else if os.Getenv("GITHUB_TOKEN") != "" {
		return ghutil.NewAuthorizedClient(os.Getenv("GITHUB_TOKEN")), nil
	}
	return nil, fmt.Errorf("error: one of the following, --token or GITHUB_TOKEN must be set")
}

func getGithubClient() *github.Client {
	if isAuth() {
		authClient, err := getAuthClient()
		if err == nil {
			return authClient
		}
	}
	return github.NewClient(nil)
}

// CreateGhcliCommand creates the root command for ghcli
func CreateGhcliCommand() *cobra.Command {
	ghcliCmd := &cobra.Command{
		Use:   "ghcli",
		Short: "Headless GitHub operations via cli",
	}
	repoCmd := CreateRepoCommand()
	ghcliCmd.AddCommand(repoCmd)
	teamCmd := CreateTeamCommand()
	ghcliCmd.AddCommand(teamCmd)
	return ghcliCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := CreateGhcliCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
