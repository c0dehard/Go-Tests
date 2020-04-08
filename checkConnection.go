package main

import (
	"fmt"
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

func main() {

	storedInfos := []server{
		server{alias: "Cloudflare", ip: "1.1.1.1"},
		server{alias: "Google DNS", ip: "8.8.8.8"},
		server{alias: "Home-Office", ip: "127.0.0.1"},
	}

	for _, stored := range storedInfos {
		fmt.Println(stored.getInfo())
	}

}
