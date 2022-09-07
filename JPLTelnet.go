package main

import (
	"fmt"

	"example.com/telnet_interface.go"

)

func main() {
	var response string = ""
	var response_error error
	var website string = "horizons.jpl.nasa.gov"
	var portnumber int = 6775

	response, response_error = communications.Telnet_connect(website, portnumber)

	if(response_error == "") {
	
		fmt.Sprintf("%s",response)
	
	}
}
