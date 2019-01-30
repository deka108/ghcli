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
	"encoding/json"
	"fmt"

	"github.com/deka108/ghcli/ghutil"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createRepoFlagItems = []struct {
	field    string
	optional bool
}{
	{"name", false},
	{"description", true},
	{"gitignoreTemplate", true},
	{"teamId", true},
	{"private", false},
	{"autoInit", false},
}

func repositoryFromCmd(cmd *cobra.Command) *github.Repository {
	data := make(map[string]interface{})
	flags := cmd.Flags()
	for _, flagItem := range createRepoFlagItems {
		field := flagItem.field
		scField := ghutil.ToSnakeCase(field)
		if flagItem.optional {
			if flags.Changed(field) {
				data[scField] = viper.Get(field)
			}
		} else {
			data[scField] = viper.Get(field)
		}
	}
	jsonStr, _ := json.Marshal(data)
	var repo *github.Repository
	json.Unmarshal(jsonStr, &repo)

	return repo
}

// CreateRepoCommand creates the cli for GitHub repo related commands
func CreateRepoCommand() *cobra.Command {
	// repoCmd represents the repo command
	var repoCmd = &cobra.Command{
		Use:   "repo [operations]",
		Short: "Perform github operations related to repository",
	}

	var createRepoCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a repository if it does not exist",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			owner, name, orgName := viper.GetString("owner"), viper.GetString("name"), viper.GetString("orgName")
			client := getGithubClient()
			_, resp, _ := client.Repositories.Get(context.Background(), owner, name)
			if resp.StatusCode == 404 {
				authClient, err := getAuthClient()
				newRepo := repositoryFromCmd(cmd)
				if err != nil {
					return err
				}
				if orgName != "" {
					_, _, err = authClient.Repositories.Create(context.Background(), orgName, newRepo)
				} else {
					_, _, err = authClient.Repositories.Create(context.Background(), "", newRepo)
				}
				if err != nil {
					return err
				}
				fmt.Printf("Repository %s is successfully created\n", name)
			} else {
				return fmt.Errorf("repo %s/%s already exist on GitHub", owner, name)
			}
			return nil
		},
	}

	var getRepoCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a repository",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		Run: func(cmd *cobra.Command, args []string) {
			client := getGithubClient()
			repo, _, _ := client.Repositories.Get(context.Background(), viper.GetString("owner"), viper.GetString("name"))
			ghutil.PrettyPrint(repo)
		},
	}

	// var addCollaboratorCmd = &cobra.Command{
	// 	Use:   "addCollaborator",
	// 	Short: "Adds user as collaborator to repository",
	// 	PreRun: func(cmd *cobra.Command, args []string) {
	// 		viper.BindPFlags(cmd.Flags())
	// 	},
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		// client := getGithubClient()
	// 		// repo, _, _ := client.Repositories.AddCollaborator(context.Background(),
	// 		// 	viper.GetString("owner"), viper.GetString("repo"), viper.GetString("user"))
	// 		// ghutil.PrettyPrint(repo)
	// 	},
	// }

	// Common repo flags
	repoCmd.PersistentFlags().String("name", "", "Repo Name (Required)")
	repoCmd.PersistentFlags().String("owner", "", "Repo's Owner (Required)")
	repoCmd.MarkPersistentFlagRequired("name")
	repoCmd.MarkPersistentFlagRequired("owner")

	// Create Repo Flags
	createRepoCmd.PersistentFlags().String("description", "", "Repo's Description")
	createRepoCmd.PersistentFlags().String("gitignoreTemplate", "", "Gitignore template")
	createRepoCmd.PersistentFlags().Int("teamId", -1, "Team ID")
	createRepoCmd.PersistentFlags().Bool("private", false, "Option to indicate if repo is private")
	createRepoCmd.PersistentFlags().Bool("autoInit", false, "Option to autoinit repo")
	createRepoCmd.PersistentFlags().String("orgName", "", "Add organization name for the repo (optional)")

	// addCollaboratorCmd.Flags().String("user", "", "GitHub username")
	// addCollaboratorCmd.MarkFlagRequired("user")

	repoCmd.AddCommand(createRepoCmd)
	repoCmd.AddCommand(getRepoCmd)
	// repoCmd.AddCommand(addCollaboratorCmd)

	return repoCmd
}
