package main

import (
	"SynAck/internal/app"
	"os"
)

func main() {
	addr, grt := "scanme.nmap.org", "8"
	switch true {
	case len(os.Args) == 2:
		addr = os.Args[1]
		break
	case len(os.Args) >= 3:
		addr, grt = os.Args[1], os.Args[2]
	default:
		break
	}
	app.Run(addr, grt)
}
