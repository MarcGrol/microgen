package main

import (
	"github.com/MarcGrol/microgen/tool/dsl"
)

var (
	tourCreated = dsl.Event{
		Id:   1,
		Name: "TourCreated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	cyclistCreated = dsl.Event{
		Id:   2,
		Name: "CyclistCreated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "cyclistId", Type: dsl.TypeInt},
			{Name: "cyclistName", Type: dsl.TypeString},
			{Name: "cyclistTeam", Type: dsl.TypeString},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	etappeCreated = dsl.Event{
		Id:   3,
		Name: "EtappeCreated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "etappeId", Type: dsl.TypeInt},
			{Name: "etappeDate", Type: dsl.TypeTimestamp},
			{Name: "etappeStartLocation", Type: dsl.TypeString},
			{Name: "etappeFinishLocation", Type: dsl.TypeString},
			{Name: "etappeLength", Type: dsl.TypeInt},
			{Name: "etappeKind", Type: dsl.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	gamblerCreated = dsl.Event{
		Id:   4,
		Name: "GamblerCreated",
		Attributes: []dsl.Attribute{
			{Name: "gamblerUid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
			{Name: "gamblerName", Type: dsl.TypeString},
			{Name: "gamblerEmail", Type: dsl.TypeString},
			{Name: "gamblerImageIUrl", Type: dsl.TypeString},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	gamblerTeamCreated = dsl.Event{
		Id:   5,
		Name: "GamblerTeamCreated",
		Attributes: []dsl.Attribute{
			{Name: "gamblerUid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
			{Name: "year", Type: dsl.TypeInt},
			{Name: "gamblerCyclists", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	etappeStarted = dsl.Event{
		Id:   6,
		Name: "EtappeStarted",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "etappeId", Type: dsl.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	etappeResultsCreated = dsl.Event{
		Id:   7,
		Name: "EtappeResultsCreated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "lastEtappeId", Type: dsl.TypeInt},
			{Name: "bestDayCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
			{Name: "bestAllrondersCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
			{Name: "bestSprintersCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
			{Name: "bestClimberCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	cyclistScoreCalculated = dsl.Event{
		Id:   8,
		Name: "CyclistScoreCalculated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "cyclistId", Type: dsl.TypeInt},
			{Name: "lastEtappeId", Type: dsl.TypeInt},
			{Name: "newScore", Type: dsl.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	gamblerScoreCalculated = dsl.Event{
		Id:   9,
		Name: "GamblerScoreCalculated",
		Attributes: []dsl.Attribute{
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "gamblerUid", Type: dsl.TypeString},
			{Name: "lastEtappeId", Type: dsl.TypeInt},
			{Name: "newScore", Type: dsl.TypeInt},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	newsItemCreated = dsl.Event{
		Id:   10,
		Name: "NewsItemCreated",
		Attributes: []dsl.Attribute{
			{Name: "uid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
			{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
			{Name: "timestamp", Type: dsl.TypeTimestamp, Cardinality: dsl.Mandatory},
			{Name: "message", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
			{Name: "sender", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
		},
		AggregateName:      "news",
		AggregateFieldName: "uid",
	}

	application = dsl.Application{
		Name:    "tourApp",
		Package: "github.com/MarcGrol/microgen",
		Services: []dsl.Service{
			{
				Name: "Tour",
				Commands: []dsl.Command{
					{
						Name:   "CreateTour",
						Method: dsl.Post,
						Url:    "/tour",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{},
						ProducesEvents: []dsl.Event{tourCreated},
					},
					{
						Name:   "CreateCyclist",
						Method: dsl.Post,
						Url:    "/tour/:year/cyclist",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "id", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "name", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "team", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated},
						ProducesEvents: []dsl.Event{cyclistCreated},
					},
					{
						Name:   "CreateEtappe",
						Method: dsl.Post,
						Url:    "/tour/:year/etappe",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "id", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "date", Type: dsl.TypeTimestamp, Cardinality: dsl.Mandatory},
								{Name: "startLocation", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "finishLocation", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "length", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "kind", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated},
						ProducesEvents: []dsl.Event{etappeCreated},
					},
					{
						Name:   "CreateEtappeResults",
						Method: dsl.Post,
						Url:    "/tour/:year/etappe/:etappeId/results",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "etappeId", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "bestDayCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
								{Name: "bestAllroundCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
								{Name: "bestClimbCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
								{Name: "bestSprintCyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated, etappeCreated, cyclistCreated},
						ProducesEvents: []dsl.Event{etappeResultsCreated},
					},
					{
						Name:   "GetTour",
						Method: dsl.Get,
						Url:    "/tour/:year",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						OutputName:     "*Tour",
						ConsumesEvents: []dsl.Event{tourCreated, etappeCreated, cyclistCreated, etappeResultsCreated},
						ProducesEvents: []dsl.Event{},
					},
				},
			},
			{
				Name: "Gambler",
				Commands: []dsl.Command{
					{
						Name:   "CreateGambler",
						Method: dsl.Post,
						Url:    "/gambler",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "gamblerUid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "name", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "email", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated},
						ProducesEvents: []dsl.Event{gamblerCreated},
					},
					{
						Name:   "CreateGamblerTeam",
						Method: dsl.Post,
						Url:    "/gambler/:gamblerUid/team",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "gamblerUid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "cyclistIds", Type: dsl.TypeInt, Cardinality: dsl.Multiple},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated, cyclistCreated},
						ProducesEvents: []dsl.Event{gamblerTeamCreated},
					},
					{
						Name:   "GetGambler",
						Method: dsl.Get,
						Url:    "/gambler/:gamblerUid",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "gamblerUid", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						OutputName:     "*Gambler",
						ConsumesEvents: []dsl.Event{gamblerCreated, gamblerTeamCreated},
						ProducesEvents: []dsl.Event{},
					},
					{
						Name:   "GetResults",
						Method: dsl.Get,
						Url:    "/results",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						OutputName:     "*Results",
						ConsumesEvents: []dsl.Event{tourCreated, cyclistCreated, etappeCreated, gamblerCreated, gamblerTeamCreated, etappeResultsCreated},
						ProducesEvents: []dsl.Event{},
					},
				},
			},
			{
				Name: "News",
				Commands: []dsl.Command{
					{
						Name:   "CreateNewsItem",
						Method: dsl.Post,
						Url:    "/news",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
								{Name: "timestamp", Type: dsl.TypeTimestamp, Cardinality: dsl.Mandatory},
								{Name: "message", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
								{Name: "sender", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated, etappeCreated, cyclistCreated, etappeResultsCreated},
						ProducesEvents: []dsl.Event{newsItemCreated},
					},
					{
						Name:   "GetNews",
						Method: dsl.Get,
						Url:    "/news",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "year", Type: dsl.TypeInt, Cardinality: dsl.Mandatory},
							},
						},
						OutputName:     "*News",
						ConsumesEvents: []dsl.Event{tourCreated, etappeCreated, cyclistCreated, etappeResultsCreated, newsItemCreated},
						ProducesEvents: []dsl.Event{},
					},
				},
			},
			{
				Name: "Notification",
				Commands: []dsl.Command{
					{
						Name:   "SubscribeToNotifications",
						Method: dsl.Post,
						Url:    "/notification",
						Input: dsl.Entity{
							Attributes: []dsl.Attribute{
								{Name: "email", Type: dsl.TypeString, Cardinality: dsl.Mandatory},
							},
						},
						ConsumesEvents: []dsl.Event{tourCreated, cyclistCreated, etappeCreated, etappeResultsCreated, newsItemCreated},
						ProducesEvents: []dsl.Event{},
					},
				},
			},
		},
	}
)
