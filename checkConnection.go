package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	disconnected = "\033[1;31m ✖\033[0m"
	connected    = "\033[1;32m ✔\033[0m"
)

type server struct {
	alias, ip string
}


func (s server) getInfo() string {

	info := fmt.Sprintf("status: \033[1;36m%s\033[0m %s", s.alias, s.ip)
	stat := disconnected
	out, _ := exec.Command("ping", s.ip, "-t 1").Output()
	if !strings.Contains(string(out), "0 packets received") {
		stat = connected
	}
	return info + stat
}

func displayAll(storage []server) {
	for _, stored := range storage {
		fmt.Println(stored.getInfo())
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	// Sample serverinfos to test
	storedInfos := []server{
		server{alias: "Cloudflare", ip: "1.1.1.1"},
		server{alias: "Google DNS", ip: "8.8.8.8"},
		server{alias: "Home-Office", ip: "127.0.0.1"},
	}
	for {
		fmt.Print("gsCLI> ")
		text, _ := reader.ReadString('\n')
		// Convert CLRF to LF
		text = strings.Replace(text, "\n", "", -1)
		// TODO: rather use a interface for commands later on
		if strings.Compare("ls servers", text) == 0 {
			displayAll(storedInfos)
		}
		if strings.Compare("clear", text) == 0 {
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		}
		if strings.Compare("help", text) == 0 {

			fmt.Printf("To list the servers simply write 'ls servers'\n'clear' and 'exit' does work as well!\n")
		}
		if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		}
	}

}
