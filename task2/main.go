package main

import (
	"fmt"
	"main/client"
	"main/server"
	"time"
)

func main() {
	go server.RunServer()
	time.Sleep(1 * time.Second)

	address := "http://127.0.0.1:8080"
	cli := client.Client{address}

	version, err := cli.GetVersion()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(version)
	}

	origStr := "Let's Go"
	decoded, err := cli.Decode(origStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(decoded)
	}

	isFinished, code, err := cli.CallHardOp()
	if err != nil {
		fmt.Println(err.Error())
	} else if code == -1 {
		fmt.Println(isFinished)
	} else {
		fmt.Println(isFinished, code)
	}
}
