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

	couchdb "github.com/leesper/couchdb-golang"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

// Task ...
type Task struct {
	ID   string
	Text string
	Done bool
}

var (
	db   *couchdb.Database
	tNew Task

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

	connectDB()
}

func connectDB() {
	var err error
	dbString := fmt.Sprintf("http://%s:%s@%s:5984/%s_rendyfebry", dbRemoteUser, dbRemotePassword, dbRemoteHost, dbName)

	db, err = couchdb.NewDatabase(dbString)
	if err != nil {
		fmt.Println(err)
	}
}

func addTask(cmd *cobra.Command, args []string) error {
	fmt.Println("\nAdd Task")
	fmt.Println("==========================")

	newDoc := map[string]interface{}{
		"_id":  uuid.NewV4().String(),
		"text": tNew.Text,
		"done": tNew.Done,
	}

	_, _, err := db.Save(newDoc, nil)
	if err != nil {
		fmt.Println("Add Item Failed!")
		fmt.Println(err)

		return nil
	}

	fmt.Println("New task added!")
	fmt.Println(fmt.Sprintf("%+v", newDoc))

	return nil
}
