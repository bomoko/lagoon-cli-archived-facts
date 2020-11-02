package cmd

import (
	"fmt"
	"github.com/amazeeio/lagoon-cli/pkg/lagoon/filesync"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/spf13/viper"
	"os"
)

var syncFilesCommand = &cobra.Command{
	Use:   "sync",
	Short: "Syncs files from a remote environment to the local environment",
	Run: func(cmd *cobra.Command, args []string) {
		//validateToken(viper.GetString("current")) // get a new token if the current one is invalid

		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}

		sshConfig := map[string]string{
			"hostname": viper.GetString("lagoons." + cmdLagoon + ".hostname"),
			"port":     viper.GetString("lagoons." + cmdLagoon + ".port"),
			"username": cmdProjectName + "-" + cmdProjectEnvironment,
			//"rsh": "ssh -o LogLevel=ERROR -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -p 32222",
		}

		//Okay - so package import seems to work. Now we have to set up our types
		//fmt.Println(filesync.SyncFiles())
		x := filesync.SyncEnvironment{
			Rsh: "ssh -o LogLevel=ERROR -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -p 32222",
			FilePattern: "testing this out here",
			Deets: sshConfig,
		}
		fmt.Println(filesync.SyncFiles(x))

	},
}

func init() {
	rootCmd.AddCommand(syncFilesCommand)
}