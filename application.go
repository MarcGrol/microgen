package main

import (
	"github.com/xebia/microgen/spec"
)

var (
	tourCreated = spec.Event{
		Name: "TourCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
		},
	}

	cyclistCreated = spec.Event{
		Name: "CyclistCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "cyclistId", Type: spec.TypeInt},
			{Name: "cyclistName", Type: spec.TypeString},
			{Name: "cyclistTeam", Type: spec.TypeString},
		},
	}

	etappeCreated = spec.Event{
		Name: "EtappeCreated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "etaopeId", Type: spec.TypeInt},
			{Name: "etappeDate", Type: spec.TypeTimestamp},
			{Name: "etappeStartLocation", Type: spec.TypeString},
			{Name: "etappeFinishtLocation", Type: spec.TypeString},
			{Name: "etappeLength", Type: spec.TypeInt},
			{Name: "etappeKind", Type: spec.TypeInt},
		},
	}

	gamblerCreated = spec.Event{
		Name: "GamblerCreated",
		Attributes: []spec.Attribute{
			{Name: "gamblerUid", Type: spec.TypeString},
			{Name: "gamblerName", Type: spec.TypeString},
			{Name: "gamblerEmail", Type: spec.TypeString},
			{Name: "gamblerImageIUrl", Type: spec.TypeString},
		},
	}

	gamblerTeamCreated = spec.Event{
		Name: "GamblerTeamCreated",
		Attributes: []spec.Attribute{
			{Name: "gamblerUid", Type: spec.TypeString},
			{Name: "year", Type: spec.TypeInt},
			{Name: "gamblerCyclists", Type: spec.TypeInt, Cardinality: spec.Multiple},
		},
	}

	tourStarted = spec.Event{
		Name: "TourStarted",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
		},
	}

	scoringRulesAvailable = spec.Event{
		Name: "scoringRulesAvailable",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
		},
	}

	etappeStarted = spec.Event{
		Name: "EtappeStarted",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "etappeId", Type: spec.TypeInt},
		},
	}

	etappeResultsCreated = spec.Event{
		Name: "EtappeResultsAvailable",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "bestDayCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestAllrondersCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestSprintersCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
			{Name: "bestClimberCyclistIds", Type: spec.TypeInt, Cardinality: spec.Multiple},
		},
	}

	cyclistScoreCalculated = spec.Event{
		Name: "CyclistScoreCalculated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "cyclistId", Type: spec.TypeInt},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "newScore", Type: spec.TypeInt},
		},
	}

	gamblerScoreCalculated = spec.Event{
		Name: "GamblerScoreCalculated",
		Attributes: []spec.Attribute{
			{Name: "year", Type: spec.TypeInt},
			{Name: "gamblerUid", Type: spec.TypeString},
			{Name: "lastEtappeId", Type: spec.TypeInt},
			{Name: "newScore", Type: spec.TypeInt},
		},
	}

	application = spec.Application{
		Name: "Tour de france",
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
						ConsumesEvents: []spec.Event{},
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
						ConsumesEvents: []spec.Event{},
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
						ConsumesEvents: []spec.Event{},
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
						ConsumesEvents: []spec.Event{tourCreated},
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
