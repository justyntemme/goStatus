package main

import "github.com/justyntemme/goStatus/web"

func main() {
	err := 0
	for err == 0 {
		returnCode := web.SendRequest("http://nextwavesolutions.io")
		if returnCode != 200 {
			err = 1
		}
	}
	returnCode
}
