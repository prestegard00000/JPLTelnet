module example.com/JPLTelnet

go 1.18


require example.com/telnet_interface.go v0.0.0-00010101000000-000000000000

replace example.com/telnet_interface.go => ../Telnet

require (
	github.com/reiver/go-oi v1.0.0 // indirect
	github.com/reiver/go-telnet v0.0.0-20180421082511-9ff0b2ab096e // indirect
)
