package ccc

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Create(fun func()) {
	appName := getSelfFileName()
	//fmt.Println("Executor:", appName)
	var daemon bool
	var cmdStart = &cobra.Command{
		Use:   "start",
		Short: "start " + appName,
		Run: func(cc *cobra.Command, args []string) {
			if daemon {
				var srvPid = getPid(appName)
				if srvPid != "" {
					fmt.Printf("%v<%v> already started\n", appName, srvPid)
					return
				}

				cmd := exec.Command("./"+appName, "start")
				err := cmd.Start()
				if err != nil {
					fmt.Printf("%v start fail, error:%v\n", appName, err.Error())
					return
				}

				fmt.Printf(appName+" start, [PID] %d running...\n", cmd.Process.Pid)

				daemon = false
				os.Exit(0)
			}
			fun()
		},
	}
	cmdStart.Flags().BoolVarP(&daemon, "daemon", "d", false, "is daemon?")

	var cmdStop = &cobra.Command{
		Use:   "stop",
		Short: "Stop " + appName,
		Run: func(cc *cobra.Command, args []string) {
			var srvPid = getPid(appName)
			if srvPid == "" {
				fmt.Printf("%v already stopped\n", appName)
				return
			}
			err := exec.Command("kill", srvPid).Start()
			if err != nil {
				fmt.Printf("kill pid<%v> error:%v\n", srvPid, err.Error())
				return
			}
			fmt.Printf("%v<%v> stop\n", appName, srvPid)
		},
	}

	var cmdStatus = &cobra.Command{
		Use:   "status",
		Short: "status of" + appName,
		Run: func(cc *cobra.Command, args []string) {
			var srvPid = getPid(appName)
			if srvPid != "" {
				fmt.Printf("%v<%v> is running\n", appName, srvPid)
			} else {
				fmt.Printf("%v is stopped\n", appName)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: appName}
	rootCmd.AddCommand(cmdStart, cmdStop, cmdStatus)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func getSelfFileName() string {
	fullPath, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	var sep = "/"
	if strings.Contains(fullPath, "\\") {
		sep = "\\"
	}
	list := strings.Split(fullPath, sep)
	return list[len(list)-1]
}

func getPid(appName string) string {
	out, err := exec.Command("pidof", appName).Output()
	if err != nil {
		fmt.Printf("get pid of %v fail:%v\n", appName, err.Error())
		return ""
	}
	s := string(out)
	list := strings.Split(strings.Replace(s, "\n", "", -1), " ")
	for _, v := range list {
		if v != strconv.Itoa(os.Getpid()) {
			return v
		}
	}
	return ""
}
