package prov

import (
	"log"
	"time"

	"github.com/MarcGrol/microgen/tourApp/client"
	"github.com/MarcGrol/microgen/tourApp/tour"
)

func provision2015(targetHost string) error {
	year := 2015

	client := client.NewClient(targetHost)

	client.CreateTour(year)

	client.CreateEtappe(year, 1, date(year, time.July, 4), "Utrecht", "Utrecht", 14, tour.TimeTrial)
	client.CreateEtappe(year, 2, date(year, time.July, 5), "Utrecht", "Neeltje Jans", 166, tour.Flat)
	client.CreateEtappe(year, 3, date(year, time.July, 6), "Antwerpen", "Muur van Hoei", 159, tour.Hilly)
	client.CreateEtappe(year, 4, date(year, time.July, 7), "Seraing", "Cambrai", 223, tour.Flat)
	client.CreateEtappe(year, 5, date(year, time.July, 8), "Arras", "Amiens", 189, tour.Flat)
	client.CreateEtappe(year, 6, date(year, time.July, 9), "Abbeville", "Le Havre", 191, tour.Flat)
	client.CreateEtappe(year, 7, date(year, time.July, 10), "Livarot", "Fougères", 190, tour.Flat)
	client.CreateEtappe(year, 8, date(year, time.July, 11), "Rennes", "Mûr-de-Bretagne", 181, tour.Hilly)
	client.CreateEtappe(year, 9, date(year, time.July, 12), "Vannes", "Plumelec", 28, tour.TeamTimeTrial)
	//ma 13-7 	rustdag in Pau
	client.CreateEtappe(year, 10, date(year, time.July, 14), "Tarbes", "Arette La Pierre Saint Martin", 167, tour.Hilly)
	client.CreateEtappe(year, 11, date(year, time.July, 15), "Pau", "Cauterets", 188, tour.Mountains)
	client.CreateEtappe(year, 12, date(year, time.July, 16), "Lannemezan", "Plateau de Beille", 195, tour.Mountains)
	client.CreateEtappe(year, 13, date(year, time.July, 17), "Muret", "Rodez", 198, tour.Hilly)
	client.CreateEtappe(year, 14, date(year, time.July, 18), "Rodez", "Mende", 178, tour.Hilly)
	client.CreateEtappe(year, 15, date(year, time.July, 19), "Mende", "Valence", 183, tour.Hilly)
	client.CreateEtappe(year, 16, date(year, time.July, 20), "Bourg-De-Péage", "Gap", 201, tour.Hilly)
	//di 21-7 	rustdag in Gap
	client.CreateEtappe(year, 17, date(year, time.July, 22), "Digne-les-Bains", "Pra Loup", 161, tour.Mountains)
	client.CreateEtappe(year, 18, date(year, time.July, 23), "Gap", "Saint-Jean-de-Maurienne", 186, tour.Mountains)
	client.CreateEtappe(year, 19, date(year, time.July, 24), "Saint-Jean-de-Maurienne", "La Toussuire", 138, tour.Mountains)
	client.CreateEtappe(year, 20, date(year, time.July, 25), "Modane", "L’Alpe d’Huez", 110, tour.Mountains)
	client.CreateEtappe(year, 21, date(year, time.July, 26), "Sèvres", "Parijs / Champs-Élysées", 109, tour.Flat)

	team := "Astana Pro Team"
	client.CreateCyclist(year, 1, "Vincenzo Nibali", team)
	client.CreateCyclist(year, 2, "Lars Boom", team)
	client.CreateCyclist(year, 3, "Jakob Fuglsang", team)
	client.CreateCyclist(year, 4, "Andriy Grivko", team)
	client.CreateCyclist(year, 5, "Dmitriy Gruzdev", team)
	client.CreateCyclist(year, 6, "Tanel Kangert", team)
	client.CreateCyclist(year, 7, "Michele Scarponi", team)
	client.CreateCyclist(year, 8, "Rein Taaramäe", team)
	client.CreateCyclist(year, 9, "Lieuwe Westra", team)

	team = "AG2R La Mondiale"
	client.CreateCyclist(year, 11, "Jean-Christophe Peraud", team)
	client.CreateCyclist(year, 12, "Jan Bakelants", team)
	client.CreateCyclist(year, 13, "Romain Bardet", team)
	client.CreateCyclist(year, 14, "Mikael Cherel", team)
	client.CreateCyclist(year, 15, "Ben Gastauer", team)
	client.CreateCyclist(year, 16, "Damien Gaudin", team)
	client.CreateCyclist(year, 17, "Christophe Riblon", team)
	client.CreateCyclist(year, 18, "Johan Vansummeren", team)
	client.CreateCyclist(year, 19, "Alexis Vuillermoz", team)

	team = "FDJ"
	client.CreateCyclist(year, 21, "Thibaut Pinot", team)
	client.CreateCyclist(year, 22, "William Bonnet", team)
	client.CreateCyclist(year, 23, "Sébastien Chavanel", team)
	client.CreateCyclist(year, 24, "Arnaud Démare", team)
	client.CreateCyclist(year, 25, "Alexandre Geniez", team)
	client.CreateCyclist(year, 26, "Matthieu Ladagnous", team)
	client.CreateCyclist(year, 27, "Steve Morabito", team)
	client.CreateCyclist(year, 28, "Jérémy Roy", team)
	client.CreateCyclist(year, 29, "Benoit Vaugrenard", team)

	team = "Team Sky"
	client.CreateCyclist(year, 31, "Chris Froome", team)
	client.CreateCyclist(year, 32, "Peter Kennaugh", team)
	client.CreateCyclist(year, 33, "Leopold König", team)
	client.CreateCyclist(year, 34, "Wout Poels", team)
	client.CreateCyclist(year, 35, "Richie Porte", team)
	client.CreateCyclist(year, 36, "Nicolas Roche", team)
	client.CreateCyclist(year, 37, "Luke Rowe", team)
	client.CreateCyclist(year, 38, "Ian Stannard", team)
	client.CreateCyclist(year, 39, "Geraint Thomas", team)

	team = "Tinkoff – Saxo"
	client.CreateCyclist(year, 41, "Alberto Contador", team)
	client.CreateCyclist(year, 42, "Ivan Basso", team)
	client.CreateCyclist(year, 43, "Daniele Bennati", team)
	client.CreateCyclist(year, 44, "Roman Kreuziger", team)
	client.CreateCyclist(year, 45, "Rafał Majka", team)
	client.CreateCyclist(year, 46, "Michael Rogers", team)
	client.CreateCyclist(year, 47, "Peter Sagan", team)
	client.CreateCyclist(year, 48, "Matteo Tosatto", team)
	client.CreateCyclist(year, 49, "Michael Valgren", team)

	team = "Movistar Team"
	client.CreateCyclist(year, 51, "Nairo Quintana", team)
	client.CreateCyclist(year, 52, "Winner Anacona", team)
	client.CreateCyclist(year, 53, "Jonathan Castroviejo", team)
	client.CreateCyclist(year, 54, "Alex Dowsett", team)
	client.CreateCyclist(year, 55, "Imanol Erviti", team)
	client.CreateCyclist(year, 56, "José Herrada", team)
	client.CreateCyclist(year, 57, "Gorka Izagirre", team)
	client.CreateCyclist(year, 58, "Adriano Malori", team)
	client.CreateCyclist(year, 59, "Alejandro Valverde", team)

	team = "BMC Racing Team"
	client.CreateCyclist(year, 61, "Tejay van Garderen", team)
	client.CreateCyclist(year, 62, "Damiano Caruso", team)
	client.CreateCyclist(year, 63, "Rohan Dennis", team)
	client.CreateCyclist(year, 64, "Daniel Oss", team)
	client.CreateCyclist(year, 65, "Manuel Quinziato", team)
	client.CreateCyclist(year, 66, "Samuel Sánchez", team)
	client.CreateCyclist(year, 67, "Michael Schar", team)
	client.CreateCyclist(year, 68, "Greg van Avermaet", team)
	client.CreateCyclist(year, 69, "Danilo Wyss", team)

	team = "Lotto Soudal"
	client.CreateCyclist(year, 71, "Tony Gallopin", team)
	client.CreateCyclist(year, 72, "Lars Bak", team)
	client.CreateCyclist(year, 73, "Thomas De Gendt", team)
	client.CreateCyclist(year, 74, "Jens Debusschere", team)
	client.CreateCyclist(year, 75, "André Greipel", team)
	client.CreateCyclist(year, 76, "Adam Hansen", team)
	client.CreateCyclist(year, 77, "Gregory Henderson", team)
	client.CreateCyclist(year, 78, "Marcel Sieberg", team)
	client.CreateCyclist(year, 79, "Tim Wellens", team)

	team = "Giant – Alpecin"
	client.CreateCyclist(year, 81, "John Degenkolb", team)
	client.CreateCyclist(year, 82, "Warren Barguil", team)
	client.CreateCyclist(year, 83, "Roy Curvers", team)
	client.CreateCyclist(year, 84, "Koen de Kort", team)
	client.CreateCyclist(year, 85, "Tom Dumoulin", team)
	client.CreateCyclist(year, 86, "Simon Geschke", team)
	client.CreateCyclist(year, 87, "Georg Preidler", team)
	client.CreateCyclist(year, 88, "Ramon Sinkeldam", team)
	client.CreateCyclist(year, 89, "Albert Timmer", team)

	team = "Team Katusha"
	client.CreateCyclist(year, 91, "Joaquim Rodriguez", team)
	client.CreateCyclist(year, 92, "Giampaolo Caruso", team)
	client.CreateCyclist(year, 93, "Jacopo Guarnieri", team)
	client.CreateCyclist(year, 94, "Marco Haller", team)
	client.CreateCyclist(year, 95, "Dmitry Kozontchuk", team)
	client.CreateCyclist(year, 96, "Alexander Kristoff", team)
	client.CreateCyclist(year, 97, "Alberto Losada", team)
	client.CreateCyclist(year, 98, "Tiago Machado", team)
	client.CreateCyclist(year, 99, "Luca Paolini", team)

	team = "Orica GreenEDGE"
	client.CreateCyclist(year, 101, "Simon Gerrans", team)
	client.CreateCyclist(year, 102, "Michael Albasini", team)
	client.CreateCyclist(year, 103, "Luke Durbridge", team)
	client.CreateCyclist(year, 104, "Daryl Impey", team)
	client.CreateCyclist(year, 105, "Michael Matthews", team)
	client.CreateCyclist(year, 106, "Svein Tuft", team)
	client.CreateCyclist(year, 107, "Pieter Weening", team)
	client.CreateCyclist(year, 108, "Adam Yates", team)
	client.CreateCyclist(year, 109, "Simon Yates", team)

	team = "Etixx – Quick Step"
	client.CreateCyclist(year, 111, "Michal Kwiatkowski", team)
	client.CreateCyclist(year, 112, "Mark Cavendish", team)
	client.CreateCyclist(year, 113, "Michal Golas", team)
	client.CreateCyclist(year, 114, "Tony Martin", team)
	client.CreateCyclist(year, 115, "Mark Renshaw", team)
	client.CreateCyclist(year, 116, "Zdenek Stybar", team)
	client.CreateCyclist(year, 117, "Matteo Trentin", team)
	client.CreateCyclist(year, 118, "Rigoberto Uran", team)
	client.CreateCyclist(year, 119, "Julien Vermote", team)

	team = "Team Europcar"
	client.CreateCyclist(year, 121, "Pierre Rolland", team)
	client.CreateCyclist(year, 122, "Bryan Coquard", team)
	client.CreateCyclist(year, 123, "Cyril Gautier", team)
	client.CreateCyclist(year, 124, "Yohann Gene", team)
	client.CreateCyclist(year, 125, "Bryan Nauleau", team)
	client.CreateCyclist(year, 126, "Perrig Quemeneur", team)
	client.CreateCyclist(year, 127, "Romain Sicard", team)
	client.CreateCyclist(year, 128, "Angelo Tulik", team)
	client.CreateCyclist(year, 129, "Thomas Voeckler", team)

	team = "LottoNL-Jumbo"
	client.CreateCyclist(year, 131, "Wilco Kelderman", team)
	client.CreateCyclist(year, 132, "Robert Gesink", team)
	client.CreateCyclist(year, 133, "Steven Kruijswijk", team)
	client.CreateCyclist(year, 134, "Tom Leezer", team)
	client.CreateCyclist(year, 135, "Paul Martens", team)
	client.CreateCyclist(year, 136, "Bram Tankink", team)
	client.CreateCyclist(year, 137, "Laurens ten Dam", team)
	client.CreateCyclist(year, 138, "Jos van Emden", team)
	client.CreateCyclist(year, 139, "Sep Vanmarcke", team)

	team = "Trek Factory Racing"
	client.CreateCyclist(year, 141, "Bauke Mollema", team)
	client.CreateCyclist(year, 142, "Julián Arredondo", team)
	client.CreateCyclist(year, 143, "Fabian Cancellara", team)
	client.CreateCyclist(year, 144, "Stijn Devolder", team)
	client.CreateCyclist(year, 145, "Laurent Didier", team)
	client.CreateCyclist(year, 146, "Markel Irizar", team)
	client.CreateCyclist(year, 147, "Bob Jungels", team)
	client.CreateCyclist(year, 148, "Gregory Rast", team)
	client.CreateCyclist(year, 149, "Haimar Zubeldia", team)

	team = "Lampre – Merida"
	client.CreateCyclist(year, 151, "Rui Costa", team)
	client.CreateCyclist(year, 152, "Matteo Bono", team)
	client.CreateCyclist(year, 153, "Davide Cimolai", team)
	client.CreateCyclist(year, 154, "Kristijan Durasek", team)
	client.CreateCyclist(year, 155, "Nelson Oliveira", team)
	client.CreateCyclist(year, 156, "Rubén Plaza", team)
	client.CreateCyclist(year, 157, "Filippo Pozzato", team)
	client.CreateCyclist(year, 158, "José Serpa", team)
	client.CreateCyclist(year, 159, "Rafael Valls", team)

	team = "Cannondale – Garmin"
	client.CreateCyclist(year, 161, "Andrew Talansky", team)
	client.CreateCyclist(year, 162, "Jack Bauer", team)
	client.CreateCyclist(year, 163, "Nathan Haas", team)
	client.CreateCyclist(year, 164, "Ryder Hesjedal", team)
	client.CreateCyclist(year, 165, "Kristijan Koren", team)
	client.CreateCyclist(year, 166, "Sebastian Langeveld", team)
	client.CreateCyclist(year, 167, "Daniel Martin", team)
	client.CreateCyclist(year, 168, "Ramunas Navardauskas", team)
	client.CreateCyclist(year, 169, "Dylan van Baarle", team)

	team = "Cofidis, Solutions Crédits"
	client.CreateCyclist(year, 171, "Nacer Bouhanni", team)
	client.CreateCyclist(year, 172, "Nicolas Edet", team)
	client.CreateCyclist(year, 173, "Christophe Laporte", team)
	client.CreateCyclist(year, 174, "Luis Ángel Maté", team)
	client.CreateCyclist(year, 175, "Daniel Navarro", team)
	client.CreateCyclist(year, 176, "Florian Senechal", team)
	client.CreateCyclist(year, 177, "Julien Simon", team)
	client.CreateCyclist(year, 178, "Geoffrey Soupe", team)
	client.CreateCyclist(year, 179, "Kenneth Vanbilsen", team)

	team = "IAM Cycling"
	client.CreateCyclist(year, 181, "Mathias Frank", team)
	client.CreateCyclist(year, 182, "Matthias Brändle", team)
	client.CreateCyclist(year, 183, "Sylvain Chavanel", team)
	client.CreateCyclist(year, 184, "Stef Clement", team)
	client.CreateCyclist(year, 185, "Jérome Coppel", team)
	client.CreateCyclist(year, 186, "Martin Elmiger", team)
	client.CreateCyclist(year, 187, "Reto Hollenstein", team)
	client.CreateCyclist(year, 188, "Jarlinson Pantano", team)
	client.CreateCyclist(year, 189, "Marcel Wyss", team)

	team = "Bora-Argon 18"
	client.CreateCyclist(year, 191, "Emanuel Buchmann", team)
	client.CreateCyclist(year, 192, "Jan Bárta", team)
	client.CreateCyclist(year, 193, "Sam Bennett", team)
	client.CreateCyclist(year, 194, "Zakkari Dempster", team)
	client.CreateCyclist(year, 195, "Bartosz Huzarski", team)
	client.CreateCyclist(year, 196, "José Mendes", team)
	client.CreateCyclist(year, 197, "Dominik Nerz", team)
	client.CreateCyclist(year, 198, "Andreas Schillinger", team)
	client.CreateCyclist(year, 199, "Paul Voss", team)

	team = "Bretagne – Séché Environnement"
	client.CreateCyclist(year, 201, "Eduardo Sepúlveda", team)
	client.CreateCyclist(year, 202, "Frederic Brun", team)
	client.CreateCyclist(year, 203, "Anthony Delaplace", team)
	client.CreateCyclist(year, 204, "Pierrick Fedrigo", team)
	client.CreateCyclist(year, 205, "Brice Feillu", team)
	client.CreateCyclist(year, 206, "Armindo Fonseca", team)
	client.CreateCyclist(year, 207, "Arnaud Gerard", team)
	client.CreateCyclist(year, 208, "Pierre-Luc Perichon", team)
	client.CreateCyclist(year, 209, "Florian Vachon", team)

	team = "MTN – Qhubeka"
	client.CreateCyclist(year, 211, "Daniel Teklehaimanot", team)
	client.CreateCyclist(year, 212, "Edvald Boasson Hagen", team)
	client.CreateCyclist(year, 213, "Stephen Cummings", team)
	client.CreateCyclist(year, 214, "Tyler Farrar", team)
	client.CreateCyclist(year, 215, "Jacques J, v, Rensburg", team)
	client.CreateCyclist(year, 216, "Reinardt J, v, Rensburg", team)
	client.CreateCyclist(year, 217, "Merhawi Kudus", team)
	client.CreateCyclist(year, 218, "Louis Meintjes", team)
	client.CreateCyclist(year, 219, "Serge Pauwels", team)

	client.CreateGamblerTeam(year, "marc", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11})
	client.CreateGamblerTeam(year, "eva", []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21})
	client.CreateGamblerTeam(year, "pien", []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 31})
	client.CreateGamblerTeam(year, "tijl", []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 41})
	client.CreateGamblerTeam(year, "freek", []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 51})

	client.CreateNewsItem(year, date(year, time.July, 4), "Marcus", "De tour gaat weer beginnen")
	client.CreateNewsItem(year, date(year, time.July, 5), "Grol", "Echt waar")

	if client.Err != nil {

		log.Printf("Error provisioning data: %s", client.Err)
		return client.Err
	}
	return nil
}
