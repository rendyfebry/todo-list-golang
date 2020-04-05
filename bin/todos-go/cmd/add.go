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

	"github.com/rendyfebry/todo-list-golang/lib/todos"
	"github.com/spf13/cobra"
)

var (
	tNew todos.Task

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add task",
		Long:  "Add new task",
		RunE:  addTask,
	}
)

func init() {
	addCmd.ResetFlags()
	addCmd.PersistentFlags().StringVarP(&tNew.Text, "text", "t", "", "task content")
	addCmd.MarkPersistentFlagRequired("text")

	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) error {
	fmt.Println("\nAdd Task")
	fmt.Println("==========================")

	newDoc, err := todosSvc.Add(tNew.Text)
	if err != nil {
		fmt.Println("Add Item Failed!")
		fmt.Println(err)

		return nil
	}

	fmt.Println("New task added!")
	fmt.Println(fmt.Sprintf("%+v", newDoc))

	return nil
}
