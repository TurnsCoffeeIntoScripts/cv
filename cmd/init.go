/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cv/pkg/config"
	"cv/pkg/models"
	"cv/pkg/terminal"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Creates/updates the configuration file (%s.%s) in the home directory", config.FileName, config.FileType),
	Long: `TODO
TODO long desc...`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.Exists() {
			if proceed() {
				var conf models.Config
				data := config.Read()
				if err := yaml.Unmarshal(data, &conf); err != nil {
					terminal.ErrorMessage("Failed to unmarshall existing config (%s)", err.Error())
				}

				processSurveyAndWriteConfig(&conf)
			}
		} else {
			config.Create()
			processSurveyAndWriteConfig(nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func proceed() bool {
	p := false
	prompt := &survey.Confirm{
		Message: "Config already initialized. Continue anyway?",
	}
	_ = survey.AskOne(prompt, &p)

	return p
}

func processSurveyAndWriteConfig(conf *models.Config) {
	a := processSurvey(conf)
	if data, err := yaml.Marshal(models.InitAnwsersToConfig(a)); err != nil {
		terminal.ErrorMessage("Failed to marshal data (%s)", err.Error())
	} else {
		config.Write(data)
	}
}

func processSurvey(conf *models.Config) models.InitAnwsers {
	var defUsername string
	var defUrl string
	var defTeams []string
	if conf != nil {
		defUsername = conf.ConcourseConfig.Username
		defUrl = conf.ConcourseConfig.Url
		defTeams = conf.ConcourseConfig.Teams
	}
	var qs = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "Concourse username?",
				Default: defUsername,
			},
			Validate: survey.Required,
		},
		{
			Name: "concourseurl",
			Prompt: &survey.Input{
				Message: "Concourse URL?",
				Default: defUrl,
			},
		},
		{
			Name: "concourseteams",
			Prompt: &survey.MultiSelect{
				Message: "Select desired team(s):",
				Options: []string{"a", "b", "c"}, // TODO the values here should be fetch via the concourse API /api/v1/teams
				Default: defTeams,
			},
		},
	}

	answers := models.InitAnwsers{}

	if err := survey.Ask(qs, &answers); err != nil {
		fmt.Println(err.Error())
		return answers
	} else {
		return answers
	}
}
