package prov

import ()

func Start(targetHost string) error {
	return provision2012(targetHost)
}
