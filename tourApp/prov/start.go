package prov

func Start(targetHost string) error {
	//return provision2012(targetHost)
	return provision2015(targetHost)
}
