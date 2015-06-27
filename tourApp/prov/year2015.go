package prov

import (
	"log"
	"time"

	"github.com/MarcGrol/microgen/tourApp/client"
	"github.com/MarcGrol/microgen/tourApp/tour"
)

func provision2012(targetHost string) error {
	year := 2012

	client := client.NewClient(targetHost)

	client.CreateTour(year)

	client.CreateEtappe(year, 1, date(year, time.June, 30), "Liège", "Liège", 6, tour.TimeTrial)
	client.CreateEtappe(year, 2, date(year, time.July, 1), "Liège", "Seraing", 198, tour.Flat)
	client.CreateEtappe(year, 3, date(year, time.July, 2), "Visé", "Tournai", 207, tour.Flat)
	client.CreateEtappe(year, 4, date(year, time.July, 3), "Orchies", "Boulogne-sur-Mer", 200, tour.Hilly)
	client.CreateEtappe(year, 5, date(year, time.July, 4), "Abbeville", "Rouen", 214, tour.Flat)
	client.CreateEtappe(year, 6, date(year, time.July, 5), "Rouen", "Saint-Quentin", 196, tour.Flat)
	client.CreateEtappe(year, 7, date(year, time.July, 6), "Épernay", "Metz", 207, tour.Flat)
	client.CreateEtappe(year, 8, date(year, time.July, 7), "Tomblaine", "La Planche des Belles Filles", 199, tour.Hilly)
	client.CreateEtappe(year, 9, date(year, time.July, 8), "Belfort", "Porrentruy", 157, tour.Hilly)
	client.CreateEtappe(year, 10, date(year, time.July, 9), "Arc-et-Senans", "Besançon", 41, tour.TimeTrial)
	client.CreateEtappe(year, 11, date(year, time.July, 11), "Mâcon", "Bellegarde-sur-Valserine", 194, tour.Mountains)
	client.CreateEtappe(year, 12, date(year, time.July, 12), "Albertville", "La Toussuire - Les Sybelles", 148, tour.Mountains)
	client.CreateEtappe(year, 13, date(year, time.July, 13), "Saint-Jean-de-Maurienne", "Annonay Davézieux", 226, tour.Hilly)
	client.CreateEtappe(year, 14, date(year, time.July, 14), "Saint-Paul-Trois-Châteaux", "Le Cap d’Agde", 217, tour.Flat)
	client.CreateEtappe(year, 15, date(year, time.July, 15), "Limoux", "Foix", 191, tour.Mountains)
	client.CreateEtappe(year, 16, date(year, time.July, 16), "Samatan", "Pau", 158, tour.Flat)
	client.CreateEtappe(year, 17, date(year, time.July, 18), "Pau", "Bagnères-de-Luchon", 197, tour.Mountains)
	client.CreateEtappe(year, 18, date(year, time.July, 19), "Bagnères-de-Luchon", "Peyragudes", 143, tour.Mountains)
	client.CreateEtappe(year, 19, date(year, time.July, 20), "Blagnac", "Brive-la-Gaillarde", 222, tour.Mountains)
	client.CreateEtappe(year, 20, date(year, time.July, 21), "Bonneval", "Chartres", 53, tour.TimeTrial)
	client.CreateEtappe(year, 21, date(year, time.July, 22), "Rambouillet", "Paris Champs-Élysées", 120, tour.Flat)

	client.CreateCyclist(year, 9, "VAN GARDEREN Tejay", "BMC")
	client.CreateCyclist(year, 8, "SCH€R Michael", "BMC")
	client.CreateCyclist(year, 7, "QUINZIATO Manuel", "BMC")
	client.CreateCyclist(year, 6, "MOINARD Ama‘l", "BMC")
	client.CreateCyclist(year, 5, "HINCAPIE George", "BMC")
	client.CreateCyclist(year, 4, "GILBERT Philippe", "BMC")
	client.CreateCyclist(year, 3, "CUMMINGS Stephen", "BMC")
	client.CreateCyclist(year, 2, "BURGHARDT Marcus", "BMC")
	client.CreateCyclist(year, 1, "EVANS Cadel", "BMC")

	client.CreateCyclist(year, 19, "ZUBELDIA Haimar", "RNT")
	client.CreateCyclist(year, 18, "VOIGT Jens", "RNT")
	client.CreateCyclist(year, 17, "POPOVYCH Yaroslav", "RNT")
	client.CreateCyclist(year, 16, "MONFORT Maxime", "RNT")
	client.CreateCyclist(year, 15, "KL…DEN AndrŽas", "RNT")
	client.CreateCyclist(year, 14, "HORNER Christopher", "RNT")
	client.CreateCyclist(year, 13, "GALLOPIN Tony", "RNT")
	client.CreateCyclist(year, 12, "CANCELLARA Fabian", "RNT")
	client.CreateCyclist(year, 11, "SCHLECK Frank", "RNT")

	client.CreateCyclist(year, 29, "ROLLAND Pierre", "EUC")
	client.CreateCyclist(year, 28, "MALACARNE Davide", "EUC")
	client.CreateCyclist(year, 27, "KERN Christophe", "EUC")
	client.CreateCyclist(year, 26, "JEROME Vincent", "EUC")
	client.CreateCyclist(year, 25, "GENE Yohann", "EUC")
	client.CreateCyclist(year, 24, "GAUTIER Cyril", "EUC")
	client.CreateCyclist(year, 23, "BERNAUDEAU Giovanni", "EUC")
	client.CreateCyclist(year, 22, "ARASHIRO Yukiya", "EUC")
	client.CreateCyclist(year, 21, "VOECKLER Thomas", "EUC")

	client.CreateCyclist(year, 39, "VERDUGO Gorka", "EUS")
	client.CreateCyclist(year, 38, "URTASUN PEREZ Pablo", "EUS")
	client.CreateCyclist(year, 37, "TXURRUKA Amets", "EUS")
	client.CreateCyclist(year, 36, "PEREZ MORENO Ruben", "EUS")
	client.CreateCyclist(year, 35, "MARTINEZ Egoi", "EUS")
	client.CreateCyclist(year, 34, "IZAGUIRRE INSAUSTI Gorka", "EUS")
	client.CreateCyclist(year, 32, "ASTARLOZA Mikel", "EUS")
	client.CreateCyclist(year, 31, "SANCHEZ Samuel", "EUS")

	client.CreateCyclist(year, 49, "VIGANO Davide", "LAM")
	client.CreateCyclist(year, 48, "STORTONI Simone", "LAM")
	client.CreateCyclist(year, 47, "PETACCHI Alessandro", "LAM")
	client.CreateCyclist(year, 46, "MARZANO Marco", "LAM")
	client.CreateCyclist(year, 45, "LLOYD Matthew", "LAM")
	client.CreateCyclist(year, 43, "HONDO Danilo", "LAM")
	client.CreateCyclist(year, 42, "BOLE Grega", "LAM")
	client.CreateCyclist(year, 41, "SCARPONI Michele", "LAM")

	client.CreateCyclist(year, 59, "VANOTTI Alessandro", "LIQ")
	client.CreateCyclist(year, 58, "SZMYD Sylvester", "LIQ")
	client.CreateCyclist(year, 57, "SAGAN Peter", "LIQ")
	client.CreateCyclist(year, 56, "OSS Daniel", "LIQ")
	client.CreateCyclist(year, 55, "NERZ Dominik", "LIQ")
	client.CreateCyclist(year, 54, "KOREN Kristijan", "LIQ")
	client.CreateCyclist(year, 53, "CANUTI Federico", "LIQ")
	client.CreateCyclist(year, 52, "BASSO Ivan", "LIQ")
	client.CreateCyclist(year, 51, "NIBALI Vincenzo", "LIQ")

	client.CreateCyclist(year, 69, "ZABRISKIE David", "GRS")
	client.CreateCyclist(year, 68, "VANDE VELDE Christian", "GRS")
	client.CreateCyclist(year, 67, "VAN SUMMEREN Johan", "GRS")
	client.CreateCyclist(year, 66, "MILLAR David", "GRS")
	client.CreateCyclist(year, 65, "MARTIN Daniel", "GRS")
	client.CreateCyclist(year, 64, "HUNTER Robert", "GRS")
	client.CreateCyclist(year, 63, "FARRAR Tyler", "GRS")
	client.CreateCyclist(year, 62, "DANIELSON Thomas", "GRS")
	client.CreateCyclist(year, 61, "HESJEDAL Ryder", "GRS")

	client.CreateCyclist(year, 79, "ROCHE Nicolas", "ALM")
	client.CreateCyclist(year, 78, "RIBLON Christophe", "ALM")
	client.CreateCyclist(year, 77, "MINARD SŽbastien", "ALM")
	client.CreateCyclist(year, 76, "KADRI Blel", "ALM")
	client.CreateCyclist(year, 75, "HINAULT SŽbastien", "ALM")
	client.CreateCyclist(year, 74, "DUPONT Hubert", "ALM")
	client.CreateCyclist(year, 73, "CHEREL Mikael", "ALM")
	client.CreateCyclist(year, 72, "BOUET Maxime", "ALM")
	client.CreateCyclist(year, 71, "PERAUD Jean-Christophe", "ALM")

	client.CreateCyclist(year, 89, "ZINGLE Romain", "COF")
	client.CreateCyclist(year, 88, "MONCOUTIE David", "COF")
	client.CreateCyclist(year, 87, "MATE MARDONES Luis Angel", "COF")
	client.CreateCyclist(year, 86, "GHYSELINCK Jan", "COF")
	client.CreateCyclist(year, 85, "FOUCHARD Julien", "COF")
	client.CreateCyclist(year, 84, "EDET Nicolas", "COF")
	client.CreateCyclist(year, 83, "DUMOULIN Samuel", "COF")
	client.CreateCyclist(year, 82, "DI GREGORIO RŽmy", "COF")
	client.CreateCyclist(year, 81, "TAARAMAE Rein", "COF")

	client.CreateCyclist(year, 99, "SIMON Julien", "SAU")
	client.CreateCyclist(year, 98, "MARINO Jean Marc", "SAU")
	client.CreateCyclist(year, 97, "LEVARLET Guillaume", "SAU")
	client.CreateCyclist(year, 96, "LEMOINE Cyril", "SAU")
	client.CreateCyclist(year, 95, "JEANDESBOZ Fabrice", "SAU")
	client.CreateCyclist(year, 94, "FEILLU Brice", "SAU")
	client.CreateCyclist(year, 93, "ENGOULVENT Jimmy", "SAU")
	client.CreateCyclist(year, 92, "DELAPLACE Anthony", "SAU")
	client.CreateCyclist(year, 91, "COPPEL JŽr™me", "SAU")

	client.CreateCyclist(year, 109, "SIVTSOV Kanstantsin", "SKY")
	client.CreateCyclist(year, 108, "ROGERS Michael", "SKY")
	client.CreateCyclist(year, 107, "PORTE Richie", "SKY")
	client.CreateCyclist(year, 106, "KNEES Christian", "SKY")
	client.CreateCyclist(year, 105, "FROOME Christopher", "SKY")
	client.CreateCyclist(year, 104, "EISEL Bernhard", "SKY")
	client.CreateCyclist(year, 103, "CAVENDISH Mark", "SKY")
	client.CreateCyclist(year, 102, "BOASSON HAGEN Edvald", "SKY")
	client.CreateCyclist(year, 101, "WIGGINS Bradley", "SKY")

	client.CreateCyclist(year, 119, "VANENDERT Jelle", "LTB")
	client.CreateCyclist(year, 118, "SIEBERG Marcel", "LTB")
	client.CreateCyclist(year, 117, "ROELANDTS Jurgen", "LTB")
	client.CreateCyclist(year, 116, "HENDERSON Gregory", "LTB")
	client.CreateCyclist(year, 115, "HANSEN Adam", "LTB")
	client.CreateCyclist(year, 114, "GREIPEL AndrŽ", "LTB")
	client.CreateCyclist(year, 113, "DE GREEF Francis", "LTB")
	client.CreateCyclist(year, 112, "BAK Lars", "LTB")
	client.CreateCyclist(year, 111, "VAN DEN BROECK Jurgen", "LTB")

	client.CreateCyclist(year, 129, "VAN HUMMEL Kenny Robert", "VCD")
	client.CreateCyclist(year, 128, "VALLS FERRI Rafael", "VCD")
	client.CreateCyclist(year, 127, "RUIJGH Rob", "VCD")
	client.CreateCyclist(year, 126, "POELS Wouter", "VCD")
	client.CreateCyclist(year, 125, "MARCATO Marco", "VCD")
	client.CreateCyclist(year, 124, "LARSSON Gustav", "VCD")
	client.CreateCyclist(year, 123, "HOOGERLAND Johnny", "VCD")
	client.CreateCyclist(year, 122, "BOECKMANS Kris", "VCD")
	client.CreateCyclist(year, 121, "WESTRA Lieuwe", "VCD")

	client.CreateCyclist(year, 139, "VORGANOV Eduard", "KAT")
	client.CreateCyclist(year, 138, "TROFIMOV Yury", "KAT")
	client.CreateCyclist(year, 137, "PAOLINI Luca", "KAT")
	client.CreateCyclist(year, 136, "KUCHYNSKI Aliaksandr", "KAT")
	client.CreateCyclist(year, 135, "HORRACH Joan", "KAT")
	client.CreateCyclist(year, 134, "GUSEV Vladimir", "KAT")
	client.CreateCyclist(year, 133, "FREIRE Oscar", "KAT")
	client.CreateCyclist(year, 132, "CARUSO Gianpaolo", "KAT")
	client.CreateCyclist(year, 131, "MENCHOV Denis", "KAT")

	client.CreateCyclist(year, 149, "VICHOT Arthur", "FDJ")
	client.CreateCyclist(year, 148, "ROY JŽrŽmy", "FDJ")
	client.CreateCyclist(year, 147, "ROUX Anthony", "FDJ")
	client.CreateCyclist(year, 146, "PINOT Thibaut", "FDJ")
	client.CreateCyclist(year, 145, "PINEAU Cedric", "FDJ")
	client.CreateCyclist(year, 144, "LADAGNOUS Matthieu", "FDJ")
	client.CreateCyclist(year, 143, "HUTAROVICH Yauheni", "FDJ")
	client.CreateCyclist(year, 142, "FEDRIGO Pierrick", "FDJ")
	client.CreateCyclist(year, 141, "CASAR Sandy", "FDJ")

	client.CreateCyclist(year, 159, "WYNANTS Maarten", "RAB")
	client.CreateCyclist(year, 158, "TJALLINGII Maarten", "RAB")
	client.CreateCyclist(year, 157, "TEN DAM Laurens", "RAB")
	client.CreateCyclist(year, 156, "TANKINK Bram", "RAB")
	client.CreateCyclist(year, 155, "SANCHEZ Luis-Leon", "RAB")
	client.CreateCyclist(year, 154, "RENSHAW Mark", "RAB")
	client.CreateCyclist(year, 153, "MOLLEMA Bauke", "RAB")
	client.CreateCyclist(year, 152, "KRUIJSWIJK Steven", "RAB")
	client.CreateCyclist(year, 151, "GESINK Robert", "RAB")

	client.CreateCyclist(year, 169, "ROJAS Jose Joaquin", "MOV")
	client.CreateCyclist(year, 168, "PLAZA MOLINA Ruben", "MOV")
	client.CreateCyclist(year, 167, "KIRYIENKA Vasili", "MOV")
	client.CreateCyclist(year, 166, "KARPETS Vladimir", "MOV")
	client.CreateCyclist(year, 165, "GUTIERREZ JosŽ Ivan", "MOV")
	client.CreateCyclist(year, 164, "ERVITI Imanol", "MOV")
	client.CreateCyclist(year, 163, "COSTA Rui Alberto", "MOV")
	client.CreateCyclist(year, 162, "COBO ACEBO Juan Jose", "MOV")
	client.CreateCyclist(year, 161, "VALVERDE Alejandro", "MOV")

	client.CreateCyclist(year, 179, "SORENSEN Nicki", "STB")
	client.CreateCyclist(year, 178, "SORENSEN Chris Anker", "STB")
	client.CreateCyclist(year, 176, "NUYENS Nick", "STB")
	client.CreateCyclist(year, 175, "MORKOV Michael", "STB")
	client.CreateCyclist(year, 174, "LUND Anders", "STB")
	client.CreateCyclist(year, 173, "KROON Karsten", "STB")
	client.CreateCyclist(year, 172, "HAEDO Juan Jose", "STB")

	client.CreateCyclist(year, 189, "VINOKOUROV Alexandre", "AST")
	client.CreateCyclist(year, 188, "KISERLOVSKI Robert", "AST")
	client.CreateCyclist(year, 187, "KESSIAKOFF Fredrik", "AST")
	client.CreateCyclist(year, 186, "KASHECHKIN Andrey", "AST")
	client.CreateCyclist(year, 185, "IGLINSKIY Maxim", "AST")
	client.CreateCyclist(year, 184, "GRIVKO Andriy", "AST")
	client.CreateCyclist(year, 183, "FOFONOV Dmitriy", "AST")
	client.CreateCyclist(year, 182, "BOZIC Borut", "AST")
	client.CreateCyclist(year, 181, "BRAJKOVIC Janez", "AST")

	client.CreateCyclist(year, 199, "VELITS Peter", "OPQ")
	client.CreateCyclist(year, 198, "VELITS Martin", "OPQ")
	client.CreateCyclist(year, 197, "PINEAU JŽr™me", "OPQ")
	client.CreateCyclist(year, 196, "MARTIN Tony", "OPQ")
	client.CreateCyclist(year, 195, "GRABSCH Bert", "OPQ")
	client.CreateCyclist(year, 194, "DEVENYNS Dries", "OPQ")
	client.CreateCyclist(year, 193, "DE WEERT Kevin", "OPQ")
	client.CreateCyclist(year, 192, "CHAVANEL Sylvain", "OPQ")
	client.CreateCyclist(year, 191, "LEIPHEIMER Levi", "OPQ")

	client.CreateCyclist(year, 209, "WEENING Pieter", "OGE")
	client.CreateCyclist(year, 208, "OÕGRADY Stuart", "OGE")
	client.CreateCyclist(year, 207, "LANGEVELD Sebastian", "OGE")
	client.CreateCyclist(year, 206, "LANCASTER Brett Daniel", "OGE")
	client.CreateCyclist(year, 205, "IMPEY Daryl", "OGE")
	client.CreateCyclist(year, 204, "GOSS Matthew Harley", "OGE")
	client.CreateCyclist(year, 203, "COOKE Baden", "OGE")
	client.CreateCyclist(year, 202, "ALBASINI Michael", "OGE")
	client.CreateCyclist(year, 201, "GERRANS Simon", "OGE")

	client.CreateCyclist(year, 219, "VEELERS Tom", "ARG")
	client.CreateCyclist(year, 218, "TIMMER Albert", "ARG")
	client.CreateCyclist(year, 217, "SPRICK Matthieu", "ARG")
	client.CreateCyclist(year, 216, "HUGUET Yann", "ARG")
	client.CreateCyclist(year, 215, "GRETSCH Patrick", "ARG")
	client.CreateCyclist(year, 214, "FR…HLINGER Johannes", "ARG")
	client.CreateCyclist(year, 213, "DE KORT Koen", "ARG")
	client.CreateCyclist(year, 212, "CURVERS Roy", "ARG")
	client.CreateCyclist(year, 211, "KITTEL Marcel", "ARG")

	client.CreateGambler("marc", "Marc Grol", "marc.grol@gmail.com")
	client.CreateGamblerTeam(year, "marc", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11})

	client.CreateGambler("eva", "Eva Berkhout", "eva.marc@hetnet.com")
	client.CreateGamblerTeam(year, "eva", []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 21})

	client.CreateGambler("pien", "Pien Grol", "pien.grol@gmail.com")
	client.CreateGamblerTeam(year, "pien", []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 31})

	client.CreateGambler("tijl", "Tijl Grol", "tijl.grol@gmail.com")
	client.CreateGamblerTeam(year, "tijl", []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 41})

	client.CreateGambler("freek", "Freek Grol", "freek.grol@gmail.com")
	client.CreateGamblerTeam(year, "freek", []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 51})

	if client.Err != nil {

		log.Printf("Error provisioning data: %s", client.Err)
		return client.Err
	}
	return nil
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 9, 0, 0, 0, time.Local)
}
