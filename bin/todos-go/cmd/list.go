/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

const (
	dbName           = "todos"
	dbRemoteUser     = "admin"
	dbRemotePassword = "iniadmin"
	dbRemoteHost     = "13.250.43.79"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  `List all task.`,
	RunE:  listTask,
}

func init() {
	rootCmd.AddCommand(listCmd)

	connectDB()
}

func listTask(cmd *cobra.Command, args []string) error {
	fmt.Println("\nTask list")
	fmt.Println("==========================")

	docs, err := db.QueryJSON(`{"selector": {}, "limit": 1000}`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(docs) == 0 {
		fmt.Println("Empty!")
		return nil
	}

	for _, doc := range docs {
		doneMark := " "
		if doc["done"] != nil && doc["done"].(bool) {
			doneMark = "X"
		}

		fmt.Println(fmt.Sprintf("[%s] %s - %s", doneMark, doc["_id"], doc["text"]))
	}

	return nil
}