package main

import (
	"fmt"

	"gitee.com/os4top_admin/cloud-station-g7/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}