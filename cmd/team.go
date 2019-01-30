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
	"context"
	"fmt"

	"github.com/deka108/ghcli/ghutil"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CreateTeamCommand creates cli for GitHub team command
func CreateTeamCommand() *cobra.Command {
	var teamCmd = &cobra.Command{
		Use:   "team [operations]",
		Short: "Perform GitHub team actions",
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all the available teams",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getGithubClient()
			teams, _, err := client.Teams.ListTeams(context.Background(), viper.GetString("org"), nil)
			if err != nil {
				return err
			}
			ghutil.PrettyPrint(teams)
			return nil
		},
	}

	var getTeamFromNameCmd = &cobra.Command{
		Use:   "getTeamFromName",
		Short: "Gets the Team of an organization from a team name",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getGithubClient()
			teamName := viper.GetString("name")
			var resTeam *github.Team
			teams, _, err := client.Teams.ListTeams(context.Background(), viper.GetString("org"), nil)
			if err != nil {
				return err
			}
			for _, team := range teams {
				if *team.Name == teamName {
					resTeam = team
					break
				}
			}
			ghutil.PrettyPrint(resTeam)
			return nil
		},
	}

	var getTeamFromIdCmd = &cobra.Command{
		Use:   "getTeamFromId",
		Short: "Gets the Team from a team ID",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getGithubClient()
			teamId := viper.GetInt64("id")
			team, resp, err := client.Teams.GetTeam(context.Background(), teamId)
			if resp.StatusCode == 404 {
				return fmt.Errorf("teamId: %d doesn't exist", teamId)
			}
			if err != nil {
				return err
			}
			ghutil.PrettyPrint(team)
			return nil
		},
	}

	var addTeamRepoCmd = &cobra.Command{
		Use:   "addTeamRepo",
		Short: "Add team's repository",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := getGithubClient()
			teamID, owner, repo := viper.GetInt64("id"), viper.GetString("owner"), viper.GetString("repo")
			opt := &github.TeamAddTeamRepoOptions{Permission: viper.GetString("permission")}
			resp, err := client.Teams.AddTeamRepo(context.Background(), teamID, owner, repo, opt)
			if err != nil {
				return err
			}
			if resp.StatusCode == 404 {
				return fmt.Errorf("teamId: %d doesn't exist", teamID)
			}
			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
				fmt.Printf("Updating team's access to repo %s is successful!\n", repo)
			}
			return nil
		},
	}

	listCmd.Flags().String("org", "", "Organization Name (required)")
	listCmd.MarkFlagRequired("org")

	getTeamFromNameCmd.Flags().String("name", "", "Team Name (required)")
	getTeamFromNameCmd.MarkFlagRequired("name")
	getTeamFromNameCmd.Flags().String("org", "", "Organization Name (required)")
	getTeamFromNameCmd.MarkFlagRequired("org")

	getTeamFromIdCmd.Flags().Int64("id", -1, "Team ID (required)")
	getTeamFromIdCmd.MarkFlagRequired("id")

	addTeamRepoCmd.Flags().Int64("id", -1, "Team ID (required)")
	addTeamRepoCmd.MarkFlagRequired("id")
	addTeamRepoCmd.Flags().String("owner", "", "Owner (required)")
	addTeamRepoCmd.MarkFlagRequired("owner")
	addTeamRepoCmd.Flags().String("repo", "", "Repo (required)")
	addTeamRepoCmd.MarkFlagRequired("repo")
	addTeamRepoCmd.Flags().String("permission", "push", "Team's Repo Permission. Allowed values: push,admin,read")

	teamCmd.AddCommand(listCmd)
	teamCmd.AddCommand(getTeamFromIdCmd)
	teamCmd.AddCommand(getTeamFromNameCmd)
	teamCmd.AddCommand(addTeamRepoCmd)

	return teamCmd
}
