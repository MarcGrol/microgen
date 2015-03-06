package prov

import (
	"log"
)

func Start(targetHost string) error {
	log.Printf("Start provisioning")
	return provision2012(targetHost)
}
