package main

import (
	"log"
	"os/exec"
)

func main() {
	cmdStartGodoc := exec.Command("godoc", "-http=:6060")
	errStartGodoc := cmdStartGodoc.Run()
	if errStartGodoc != nil {
		log.Println("failed to start godoc:", errStartGodoc)
	}
	cmdOpenBrowser := exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:6060/doc")
	errOpenBrowser := cmdOpenBrowser.Run()
	if errOpenBrowser != nil {
		log.Println("failed to open browser:", errOpenBrowser)
	}
}
