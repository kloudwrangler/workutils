/*
Package cmd
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"path/filepath"

	"github.com/aldorperez1/workutils/task"
	"github.com/aldorperez1/workutils/utils"
	//"path/filepath"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		fmt.Printf("%v\n", args)
	},
}
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "createProjectCommand creates a work project",
	Long: `createProjectCommand creates a work project. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: createProjectCommand,
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
		t := task.Task{Name: args[0]}
		err := t.CreateTask()
		if err != nil {
			fmt.Println("Unable to create task")
		}

	},
}
var effortCmd = &cobra.Command{
	Use:   "effort",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
		t := task.Effort{Name: args[0]}
		err := t.CreateEffort()
		if err != nil {
			fmt.Println("Unable to create task")
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(projectCmd)
	createCmd.AddCommand(taskCmd)
	createCmd.AddCommand(effortCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectCmd.Flags().String("prefix", "", "Add this number to the front of the project")

}

// createProjectCommand creates a work project
func createProjectCommand(cmd *cobra.Command, args []string) {
	var projBaseName string
	if len(args) == 0 {
		fmt.Println("Must provide name of project")
		return
	} else {
		projBaseName = args[0]
	}
	fmt.Println("Creating a project called: ", args[0])

	// TODO: validate directory name

	// Make sure that home directory exists
	projHome := "/Users/aldoperez/Projects"
	if err := createProjectHomeDir(projHome); err != nil {
		fmt.Println(err)
		return
	}
	prefix, err := cmd.Flags().GetString("prefix")
	if err != nil {
		panic("ahh")
	}
	fmt.Println(prefix)

	if len(prefix) == 0 {
		fmt.Println("Empty Prefix. Generating one")
		// Get the name of the project based on the contents of the home directory
		var err error
		prefix, err = utils.GeneratePrefix(projHome)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	projName := fmt.Sprintf("%s-%s", prefix, projBaseName)
	fmt.Println(projName)
	fmt.Println("making a project called: ", projName)
	projNameFullPath := filepath.Join(projHome, projName)
	// Make sure that there is not a project with the same name
	if err := os.Mkdir(projNameFullPath, 0755); err != nil {
		fmt.Println("Error: Problems creating the project")
		return
	}
	fmt.Println("Created project: ", projName, " at ", projNameFullPath)

	if err := os.Chdir(projNameFullPath); err != nil {
		fmt.Println(err)
		return
	}
	// Make all the directories needed
	dirs := []string{"Admin", "Docs", "Notes", "Tasks", "Utils", "src"}
	for i, name := range dirs {
		newName := fmt.Sprintf("%02d-%s", i, name)
		fmt.Println(newName)
		if err := os.Mkdir(newName, 0755); err != nil {
			fmt.Println("Error creating the subdirs")
			return
		}

	}
}

//func generatePrefix(directory string) (string, error) {
//	dirItems, err := os.ReadDir(directory)
//	if err != nil {
//		fmt.Println("We had problems reading the home dir")
//		return "", err
//	}
//	// How many dirs
//	count := 0
//	for _, file := range dirItems {
//		//fmt.Println(file.Name())
//		if file.IsDir(){
//			//fmt.Println("Its a directory")
//			count++
//		}
//	}
//	fmt.Println("We have ", count, " Directories")
//	prefix := fmt.Sprintf("%02d", count+1)
//	return prefix, nil
//}
func createProjectHomeDir(projHome string) error {
	_, err := os.Stat(projHome)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist
			fmt.Println("A project home folder does not exist. Creating at ", projHome)
			if err := os.Mkdir(projHome, 0755); err != nil {
				fmt.Println("Something Happened: directory unable to be created")
				return err
			}
		} else {
			// file may or may not exist. See err for details
			fmt.Println("Something Happened: file may or may not exist")
			return err
		}
	}
	fmt.Println("Project Home exists at: ", projHome)
	return nil
}
