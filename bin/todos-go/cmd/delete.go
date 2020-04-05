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
	tDel Task

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete task",
		Long:  `Delete single task.`,
		RunE:  deleteTask,
	}
)

func init() {
	deleteCmd.ResetFlags()
	deleteCmd.PersistentFlags().StringVarP(&tDel.ID, "id", "i", "", "task id")
	deleteCmd.MarkPersistentFlagRequired("id")

	rootCmd.AddCommand(deleteCmd)

	connectDB()
}

func deleteTask(cmd *cobra.Command, args []string) error {
	fmt.Println("\nDelete Item")
	fmt.Println("==========================")

	err := db.Delete(tDel.ID)
	if err != nil {
		fmt.Println("Delete Failed!")
		fmt.Println(err)
		return nil
	}

	fmt.Println("Sucessfully delete Item!")

	return nil
}
