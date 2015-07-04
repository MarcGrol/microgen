package prov

import "github.com/MarcGrol/microgen/tourApp/client"

func Start(targetHost string) error {

	client := client.NewClient(targetHost)

	client.CreateGambler("marc", "Marc Grol", "marc.grol@gmail.com")
	client.CreateGambler("eva", "Eva Berkhout", "eva.marc@hetnet.com")
	client.CreateGambler("pien", "Pien Grol", "pien.grol@gmail.com")
	client.CreateGambler("tijl", "Tijl Grol", "tijl.grol@gmail.com")
	client.CreateGambler("freek", "Freek Grol", "freek.grol@gmail.com")

	provision2012(targetHost)
	return provision2015(targetHost)
}
