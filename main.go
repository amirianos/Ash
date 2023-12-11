package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// TODO LIST
// 1 - add config file like bashrc file for bash
// 2 - fix real time out put problem . ex ping 8.8.8.8 => fixed
// 3 - A lot of things :)
// 4 - fix ctrl + c exit problem

func main() {

	for {
		fmt.Print("$ ")
		Reader := bufio.NewReader(os.Stdin)
		cmd_string, err := Reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd_string = cmd_string[0 : len(cmd_string)-1]
		if cmd_string == "exit" {
			fmt.Println("GoodBye")
			break
		}
		cmd_spilited_string := strings.Split(cmd_string, " ")
		cmd_first_part := cmd_spilited_string[0]
		cmd_second_part := cmd_spilited_string[1:]
		if cmd_first_part == "cd" {
			err = os.Chdir(cmd_second_part[0])
			continue
		}
		cmd := exec.Command(cmd_first_part, cmd_second_part...)

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigCh
			if cmd.Process != nil {
				_ = cmd.Process.Kill() // Kill the command process

			}
		}()

		var stdoutBuf, stderrBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
		outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
		if errStr == "" {
			fmt.Errorf(errStr)
		}
		_ = outStr

	}
}
