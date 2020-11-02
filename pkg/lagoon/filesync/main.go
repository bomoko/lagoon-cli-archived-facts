package filesync

import (
"fmt"
"github.com/zloylos/grsync"
	"os"
	"time"
)

type SyncEnvironment struct {
	FilePattern string
	Rsh string
	Deets map[string]string
}

//func SyncEnvironmentFactory(filePatternString string) syncEnvironment {
//	return syncEnvironment{filePattern: filePatternString}
//}

func SyncFiles(sshConfig SyncEnvironment) string {

	task := grsync.NewTask(
		"amazeelabsv4-com-dev@ssh.lagoon.amazeeio.cloud:/tmp",
		"/tmp/testing",
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

	fmt.Println("here2")
	fmt.Println("well done")
	fmt.Println(task.Log())

	return sshConfig.FilePattern
}
