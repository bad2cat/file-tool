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

// upfileCmd represents the upfile command
var upfileCmd = &cobra.Command{
	Use:   "upfile",
	Short: "upload file",
	Long:  `upfile is tool for upload local file to remote`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Impl.UpFileFromServer(user, password, host, port, loc, remote)
		out := os.Stdout
		if err != nil {
			fmt.Fprint(out, err)
			os.Exit(1)
		}
		fmt.Fprint(out, "upload file success")
	},
}

func init() {
	rootCmd.AddCommand(upfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	upfileCmd.Flags().StringVarP(&user, "user", "U", "jx", "the remote user")
	upfileCmd.Flags().StringVarP(&password, "password", "P", "123456", "the password of remote")
	upfileCmd.Flags().StringVarP(&host, "host", "H", "localhost", "the remote host")
	upfileCmd.Flags().StringVarP(&remote, "remote", "R", "/home/"+user, "the remote path")
	upfileCmd.Flags().StringVarP(&loc, "loc", "L", "/home/"+user, "the local file")
	upfileCmd.Flags().IntVarP(&port, "port", "O", 80, "The port of remote host")

	flag.Parse()
}
