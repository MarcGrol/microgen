package main

import (
	"github.com/MarcGrol/microgen/spec"
)

var (
	tourCreated = spec.Event{
		Name: "TourCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	cyclistCreated = spec.Event{
		Name: "CyclistCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "cyclistId", Type: spec.TypeInt},
			{Name: "cyclistName", Type: spec.TypeString},
			{Name: "cyclistTeam", Type: spec.TypeString},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	etappeCreated = spec.Event{
		Name: "EtappeCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "etappeId", Type: spec.TypeInt},
			{Name: "etappeDate", Type: spec.TypeTimestamp},
			{Name: "etappeStartLocation", Type: spec.TypeString},
			{Name: "etappeFinishLocation", Type: spec.TypeString},
			{Name: "etappeLength", Type: spec.TypeInt},
			{Name: "etappeKind", Type: spec.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	gamblerCreated = spec.Event{
		Name: "GamblerCreated",
		Attributes: []spec.Attribute{
			{Name: "gamblerUid", Type: spec.TypeString, Cardinality: spec.Mandatory},
			{Name: "gamblerName", Type: spec.TypeString},
			{Name: "gamblerEmail", Type: spec.TypeString},
			{Name: "gamblerImageIUrl", Type: spec.TypeString},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	gamblerTeamCreated = spec.Event{
		Name: "GamblerTeamCreated",
		Attributes: []spec.Attribute{
			{Name: "gamblerUid", Type: spec.TypeString, Cardinality: spec.Mandatory},
			{Name: "year", Type: spec.TypeInt},
			{Name: "gamblerCyclists", Type: spec.TypeInt, Cardinality: spec.Multiple},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	tourStarted = spec.Event{
		Name: "TourStarted",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	scoringRulesAvailable = spec.Event{
		Name: "scoringRulesAvailable",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	etappeStarted = spec.Event{
		Name: "EtappeStarted",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "etappeId", Type: spec.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	etappeResultsCreated = spec.Event{
		Name: "EtappeResultsAvailable",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "bestDayCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestAllrondersCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestSprintersCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestClimberCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	cyclistScoreCalculated = spec.Event{
		Name: "CyclistScoreCalculated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "cyclistId", Type: spec.TypeInt},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "newScore", Type: spec.TypeInt},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	gamblerScoreCalculated = spec.Event{
		Name: "GamblerScoreCalculated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
			{Name: "gamblerUid", Type: spec.TypeString},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "newScore", Type: spec.TypeInt},
		},
		AggregateName:      "gambler",
		AggregateFieldName: "gamblerUid",
	}

	application = spec.Application{
		Name: "tourApp",
		Services: []spec.Service{
			{
				Name: "Tour",
				Commands: []spec.Command{
					{
						Name:   "CreateTour",
						Method: spec.Post,
						Url:    "/tour",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
							},
						},
						ConsumesEvents: []spec.Event{},
						ProducesEvents: []spec.Event{tourCreated},
					},
					{
						Name:   "CreateCyclist",
						Method: spec.Post,
						Url:    "/tour/:year/cyclist",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "id", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "name", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "team", Type: spec.TypeString, Cardinality: spec.Mandatory},
							},
						},
						ConsumesEvents: []spec.Event{tourCreated},
						ProducesEvents: []spec.Event{cyclistCreated},
					},
					{
						Name:   "CreateEtappe",
						Method: spec.Post,
						Url:    "/tour/:year/etappe",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "id", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "date", Type: spec.TypeTimestamp, Cardinality: spec.Mandatory},
								{Name: "startLocation", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "finishLocation", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "length", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "kind", Type: spec.TypeInt, Cardinality: spec.Mandatory},
							},
						},
						ConsumesEvents: []spec.Event{tourCreated},
						ProducesEvents: []spec.Event{etappeCreated},
					},
				},
			},
			{
				Name: "Gambler",
				Commands: []spec.Command{
					{
						Name:   "CreateGambler",
						Method: spec.Post,
						Url:    "/gambler",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "gamblerUid", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "name", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "email", Type: spec.TypeString, Cardinality: spec.Mandatory},
							},
						},
						ConsumesEvents: []spec.Event{tourCreated},
						ProducesEvents: []spec.Event{gamblerCreated},
					},
					{
						Name:   "CreateGamblerTeam",
						Method: spec.Post,
						Url:    "/gambler/:gamblerUid/team",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "gamblerUid", Type: spec.TypeString, Cardinality: spec.Mandatory},
								{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "cyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
							},
						},
						ConsumesEvents: []spec.Event{tourCreated, cyclistCreated},
						ProducesEvents: []spec.Event{gamblerTeamCreated},
					},
				},
			},
			{
				Name: "Results",
				Commands: []spec.Command{
					{
						Name:   "CreateDayResults",
						Method: spec.Post,
						Url:    "/tour/:year/etappe/:etappeId/results",
						Input: spec.Entity{
							Attributes: []spec.Attribute{
								{Name: "year", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "id", Type: spec.TypeInt, Cardinality: spec.Mandatory},
								{Name: "bestDayCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
								{Name: "bestAllroundCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
								{Name: "bestClimbCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
								{Name: "bestSprintCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
							},
						},
						ConsumesEvents: []spec.Event{tourCreated, etappeCreated, cyclistCreated, gamblerCreated, gamblerTeamCreated},
						ProducesEvents: []spec.Event{etappeResultsCreated, cyclistScoreCalculated, gamblerScoreCalculated},
					},
				},
			},
		},
	}
)
