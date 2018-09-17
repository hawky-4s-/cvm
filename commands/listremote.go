// Copyright © 2016 Christian Lipphardt <christian.lipphardt@camunda.com>
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

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:   "list-remote",
	Short: "Lists remote CamBPM versions.",
	Long: `Fetch available CamBPM versions remotely from Camunda Nexus.
Depending on supplied basic auth credentials, if will also show the CamBPM Enterprise edition versions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		snapshots, err := cmd.Flags().GetBool("snapshots")
		if err != nil {
			return err
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "list-remote called with %#v, %s, %s. Args: %s", snapshots, username, password, args)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listRemoteCmd)

	listRemoteCmd.Flags().BoolP("snapshots", "s", false, "Also list snapshot versions")
	listRemoteCmd.Flags().StringP("username", "u", "", "Basic auth username")
	listRemoteCmd.Flags().StringP("password", "p", "", "Basic auth password")
}
