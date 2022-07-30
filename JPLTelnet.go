package main

import {
	"../Telnet/Telnet
}

func main() {
	var string response
	var string website = "horizons.jpl.nasa.gov"
	var int portnumber = 6775
	
	connection = Telnet(website, portnumber)

}
