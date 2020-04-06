package main

import "fmt"
//inspired by https://gist.github.com/ik5  THANK YOU!
const (
	InfoColor    = "\033[1;37m%s\033[0m"
	MessageColor = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	OnlineColor  = "\033[1;32m%s\033[0m"
	IdleColor    = "\033[1;35m%s\033[0m"
  	JustBlack    = "\033[1;30m%s\033[0m"
)

func main() {
  // Just to display the colors
	fmt.Printf(InfoColor, "Info / General")
	fmt.Printf(MessageColor, "Message / Update")
	fmt.Printf(WarningColor, "Warning / Alert")
	fmt.Printf(ErrorColor, "Error / Offline")
	fmt.Printf(OnlineColor, "Online")
	fmt.Printf(IdleColor, "Idle / AFK / Bussy")
  	fmt.Printf(JustBlack, "Just Black")
}
