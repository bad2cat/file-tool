/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fd/cmd/Impl"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var user, password, host, loc, remote string
var port int

// downfileCmd represents the downfile command
var downfileCmd = &cobra.Command{
	Use:   "downfile",
	Short: "Download file to computer",
	Long:  `downdile is tool to transport file from one compter or virtual machine to another`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := Impl.DownFileFromServer(user, password, host, port, loc, remote)
		out := os.Stdout
		if err != nil {
			fmt.Fprint(out,err)
			os.Exit(1)
		}
		fmt.Fprint(out,"download file success")
	},
}

func init() {
	rootCmd.AddCommand(downfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	downfileCmd.Flags().StringVarP(&user, "user", "U", "jx", "the remote user")
	downfileCmd.Flags().StringVarP(&password, "password", "P", "123456", "the password of remote")
	downfileCmd.Flags().StringVarP(&host, "host", "H", "localhost", "the remote host")
	downfileCmd.Flags().StringVarP(&remote, "remote", "R", "/home/"+user, "the remote path")
	downfileCmd.Flags().StringVarP(&loc, "loc", "L", "/home/"+user, "the local file")
	downfileCmd.Flags().IntVarP(&port, "port", "O", 80, "The port of remote host")

	flag.Parse()
}
