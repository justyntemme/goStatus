package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/justyntemme/goStatus/web"
)

func main() {
	//CMDLINEOPTIONS
	var alertScript string
	var url string
	flag.StringVar(&alertScript, "exec", "", "script to execute if non-200 response is recieved. returnCode is passed as argv[1]")
	flag.StringVar(&url, "url", "", "Url to watch")
	flag.Parse()
	err := 0
	for err == 0 {
		returnCode := web.SendRequest(url)
		fmt.Println(returnCode)
		if returnCode != 200 {
			exec.Command(alertScript, string(returnCode))

		}
	}
}
