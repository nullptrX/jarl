/*
Copyright © 2019 Reijhanniel Jearl Campos <devcsrj@apache.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/devcsrj/jarl"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)
// Package is called aw
import "github.com/deanishe/awgo"

// Workflow is the main API
var wf *aw.Workflow
var repo *jarl.Mvnrepository

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jarl",
	Short: "Your trusty 'Jar locator'",
	Long: `Locate jar coordinates right from your terminal.

Example:
$ jarl gson

The ending coordinates is automatically copied to the clipboard.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			wf.NewWarningItem(":( requires a search term.", "error")
			wf.SendFeedback()
			return errors.New("requires a search term")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = wf.ClearData()
		err := handleAction(args)
		if err != nil {
			wf.WarnEmpty(err.Error(), "error")
		}
		wf.SendFeedback()
	},
}

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
	repo = new(jarl.Mvnrepository)
	repo.Init("https://mvnrepository.com")
}

func handleAction(args []string) error {
	action := args[0]

	switch action {
	case "search":
		data := strings.Split(args[1], " ")
		if len(data) > 2 {
			wf.NewWarningItem(":( Too many args.", "error")
			return nil
		}
		var query = data[0]
		if len(query) == 0 {
			wf.NewWarningItem(":( Empty query.", "error")
			return nil
		}
		var page = 1
		if len(data) == 2 {
			var err error
			page, err = strconv.Atoi(data[1])
			if err != nil {
				wf.NewWarningItem(":( Page num should be int.", "error")
				return nil
			}
		}
		searchRepo(query, page)
		break
	case "artifact":
		data := strings.Split(args[1], ":")
		if len(data) < 2 {
			wf.NewWarningItem(":( Searching artifact need 2 arguments at least.", "error")
			return nil
		}
		style := os.Getenv("STYLE")
		if len(data) > 2 {
			style = data[2]
		} else if len(style) == 0 {
			style = "gradle"
		}
		group := data[0]
		id := data[1]
		searchArtifacts(group, id, style)
		break
	case "style":
		query := args[1]
		searchStyles(query)
		break
	}
	return nil
}

func searchRepo(query string, page int) {
	results := repo.SearchArtifacts(query, page)
	artifacts := results.Artifacts
	dir, _ := os.Getwd()

	for _, artifact := range artifacts {
		title := fmt.Sprintf("%s:%s", artifact.Group, artifact.Id)
		item := wf.NewItem(title)
		item.Icon(&aw.Icon{
			Value: dir + "/icon.png",
			Type:  aw.IconTypeImage,
		})
		item.Arg(title)
		item.Subtitle(strings.ReplaceAll(artifact.Description, "\n", ""))
		item.Valid(true)
	}
	if artifacts == nil || len(artifacts) == 0 {
		wf.WarnEmpty("Repo not found.", "warning")
	}
}

func searchArtifacts(group, id, style string) {
	details := repo.GetArtifactDetails(group, id)
	versions := details.Versions
	dir, _ := os.Getwd()
	artifact := jarl.Artifact{
		Group: group,
		Id:    id,
	}
	for _, version := range versions {
		item := wf.NewItem(version.Value)
		item.Icon(&aw.Icon{
			Value: dir + "/icon.png",
			Type:  aw.IconTypeImage,
		})
		item.Subtitle(version.Date)
		item.Valid(true)
		var value string
		switch style {
		case "maven":
			value = new(jarl.MavenImportStyle).Apply(artifact, version)
			break
		case "gradle":
			value = new(jarl.GradleImportStyle).Apply(artifact, version)
			break
		case "sbt":
			value = new(jarl.SbtImportStyle).Apply(artifact, version)
			break
		case "ivy":
			value = new(jarl.IvyImportStyle).Apply(artifact, version)
			break
		case "grape":
			value = new(jarl.GrapeImportStyle).Apply(artifact, version)
			break
		case "leiningen":
			value = new(jarl.LeiningenImportStyle).Apply(artifact, version)
			break
		case "buildr":
			value = new(jarl.BuildrImportStyle).Apply(artifact, version)
			break
		}
		item.Arg(value)
	}
	if len(versions) == 0 {
		wf.WarnEmpty("Artifact not found.", "warning")
	}
}

func searchStyles(query string) {
	styles := []string{"maven", "gradle", "sbt", "ivy", "grape", "leiningen", "buildr"}
	for _, style := range styles {
		item := wf.NewItem(style)
		arg := fmt.Sprintf("%s:%s", query, style)
		item.Arg(arg)
		item.Valid(true)
	}
}

// Your workflow starts here
func run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	wf.Run(run)
}
