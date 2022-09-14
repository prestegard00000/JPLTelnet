package main

import (
	"fmt"
	"github.com/prestegard00000/telnet_interface_go"
	"error"
	//"example.com/telnet_interface.go"

)

func main() {
	var JPLProfile Telnet_profile
	JPLProfile.portnumber = 6775
	JPLProfile.site = "horizons.jpl.nasa.gov"
	//JPLProfile := communications.Telnet_profile {
//		portNumber: 6775,
//		site: "horizons.jpl.nasa.gov",
//	}

	var response string = ""
	var response_error error
	//var website string = "horizons.jpl.nasa.gov"
	//var portnumber int = 6775
	response, response_error = JPLProfile.communications.Telnet_connect()

	if(response_error != nil) {
	
		fmt.Sprintf("%s",response)
	
	}
}
