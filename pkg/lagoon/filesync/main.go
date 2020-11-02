package filesync

import (
"fmt"
"github.com/zloylos/grsync"
	"os"
	"time"
)

type SyncEnvironmentsRemote struct {
	Username string
	Hostname string
	FilePattern string
}

type SyncEnvironmentLocal struct {
	FilePattern string
}


type SyncEnvironments struct {
	Remote SyncEnvironmentsRemote
	Local SyncEnvironmentLocal
	FilePattern string
	Rsh string
	Deets map[string]string
}


func SyncFiles(sshConfig SyncEnvironments) string {

	task := grsync.NewTask(
		fmt.Sprintf("%s@%s:%s", sshConfig.Remote.Username, sshConfig.Remote.Hostname, sshConfig.Remote.FilePattern), //"amazeelabsv4-com-dev@ssh.lagoon.amazeeio.cloud:/tmp",
		sshConfig.Local.FilePattern,
		grsync.RsyncOptions{
			Rsh: sshConfig.Rsh,
		},
	)

	go func() {
		for {
			state := task.State()
			fmt.Printf(
				"progress: %.2f / rem. %d / tot. %d / sp. %s \n",
				state.Progress,
				state.Remain,
				state.Total,
				state.Speed,
			)
			time.Sleep(time.Second)
		}
	}()

	if err := task.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(task.Log())

	return sshConfig.FilePattern
}
