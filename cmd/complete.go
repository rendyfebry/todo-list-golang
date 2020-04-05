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

var (
	tEdit Task

	completeCmd = &cobra.Command{
		Use:   "complete",
		Short: "Complete a task",
		Long:  `Mark a task as completed/done.`,
		RunE:  editTask,
	}
)

func init() {
	completeCmd.ResetFlags()
	completeCmd.PersistentFlags().StringVarP(&tEdit.ID, "id", "i", "", "task id")
	completeCmd.MarkPersistentFlagRequired("id")

	rootCmd.AddCommand(completeCmd)

	connectDB()
}

func editTask(cmd *cobra.Command, args []string) error {
	fmt.Println("\nMark Item")
	fmt.Println("==========================")

	doc, err := db.Get(tEdit.ID, nil)
	if err != nil {
		fmt.Println("Mark Item Failed!")
		fmt.Println(err)

		return nil
	}

	updatedDoc := map[string]interface{}{
		"_id":  doc["_id"],
		"_rev": doc["_rev"],
		"text": doc["text"],
		"done": true,
	}

	_, _, err = db.Save(updatedDoc, nil)
	if err != nil {
		fmt.Println("Mark Item Failed!")
		fmt.Println(err)

		return nil
	}

	fmt.Println("Successfully mark!")
	fmt.Println(fmt.Sprintf("%+v", updatedDoc))

	return nil
}
