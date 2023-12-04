package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// TODO LIST
// 1 - add config file like bashrc file for bash
// 3 - fix real time out put problem . ex ping 8.8.8.8
// 3 - A lot of things :)

func main() {
	working_directory_cmd := exec.Command("pwd")
	working_directory_path, err := working_directory_cmd.Output()
	if err != nil {
		fmt.Println("I cant get working directory from OS")
		fmt.Println(err.Error())
	}
	fmt.Println("i am here and path is", string(working_directory_path))
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
		std_out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(std_out))

	}
}
