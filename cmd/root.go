// Copyright © 2018 Thrashing Code Team <thrashingcode@thrashingcode.com>
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
	"fmt"
	"os"

	"github.com/gocql/gocql"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "colligere",
	Short: "A brief description of your application",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Things happen here.")

		cluster := gocql.NewCluster("127.0.0.1")
		cluster.Keyspace = "mungedcolligere"
		session, _ := cluster.CreateSession()

		defer session.Close()

		fmt.Print(session)

		// docker run --name colligere-cassandra -d cassandra:3.11.3
		// docker exec -it colligere-cassandra bash

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.colligere.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

}
