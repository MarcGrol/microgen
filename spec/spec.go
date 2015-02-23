package spec

import (
	"fmt"
	"strings"
)

type Type int

const (
	TypeString Type = iota
	TypeBoolean
	TypeInt
	TypeTimestamp
)

type CardinalityType int

const (
	Optional CardinalityType = iota
	Mandatory
	Multiple
)

type MethodType int

const (
	Unknown MethodType = iota
	Get
	Put
	Post
	Delete
)

// application
type Application struct {
	Name     string
	Package  string
	Services []Service
}

func (a Application) NameToFirstLower() string {
	return strings.ToLower(fmt.Sprintf("%c", a.Name[0])) + a.Name[1:]
}

func (app Application) GetUniqueEvents() []Event {
	eventMap := make(map[string]Event)
	for _, service := range app.Services {
		for _, command := range service.Commands {
			for _, event := range command.ConsumesEvents {
				eventMap[event.Name] = event
			}
			for _, event := range command.ProducesEvents {
				eventMap[event.Name] = event
			}
		}
	}
	events := make([]Event, 0, 20)
	for _, event := range eventMap {
		events = append(events, event)
	}

	return events
}

func (a Application) ServiceNames() []string {
	nameList := make([]string, 0, 10)
	for _, entry := range a.Services {
		nameList = append(nameList, entry.Name)
	}
	return nameList
}

func (app Application) HasDateField() bool {
	status := false
	for _, service := range app.Services {
		if service.HasDateField() {
			status = true
			break
		}
	}
	return status
}

func (app Application) GetConsumingServiceNamesForEvent(eventName string) []string {
	serviceNames := make([]string, 0, 10)
	for _, service := range app.Services {
		events := service.GetConsumedEvents()
		for _, ev := range events {
			if ev.Name == eventName {
				serviceNames = append(serviceNames, ev.Name)
				break
			}
		}
	}
	return serviceNames
}

// 	CreateGamblerTeam -> ResultsGamblerTeamCreatedEventhandler [label="GamblerTeamCreated-event", style=dashed];

func (app Application) GraphvizEdgesForEvents() string {
	edges := make([]string, 0, 10)
	for _, service := range app.Services {
		for _, command := range service.Commands {
			for _, event := range command.ProducesEvents {
				for _, s := range app.ServicesThatConsumeEvent(event) {
					edge := fmt.Sprintf("\t\"%s%s\" -> \"%s%s\" [label=\"%s-event\", style=dashed]\n",
						service.Name,
						command.Name,
						s.Name,
						event.Name,
						event.Name)
					edges = append(edges, edge)
				}
			}
		}
	}
	var allEdgesAsString string
	for _, e := range edges {
		allEdgesAsString = allEdgesAsString + e
	}
	return allEdgesAsString
}

func (app Application) ServicesThatConsumeEvent(event Event) []Service {
	services := make([]Service, 0, 10)
	for _, service := range app.Services {
		for _, e := range service.GetConsumedEvents() {
			if event.Name == e.Name {
				services = append(services, service)
			}
		}
	}
	return services
}

// service
type Service struct {
	Name     string
	Commands []Command
}

func (s Service) FirstCommand() string {
	return s.Commands[0].Name
}

func (s Service) CommandNames() []string {
	nameList := make([]string, 0, 10)
	for _, entry := range s.Commands {
		nameList = append(nameList, entry.Name)
	}
	return nameList
}

func (service Service) GetProducedEvents() []Event {
	eventMap := make(map[string]Event)
	for _, command := range service.Commands {
		for _, event := range command.ProducesEvents {
			eventMap[event.Name] = event
		}
	}
	events := make([]Event, 0, 20)
	for _, event := range eventMap {
		events = append(events, event)
	}

	return events
}

func (service Service) GetConsumedEvents() []Event {
	externalEvents := make([]Event, 0, 20)
	for _, e := range service.GetAllEvents() {
		found := false
		for _, producedE := range service.GetProducedEvents() {
			if e.Name == producedE.Name {
				found = true
			}
		}
		if found == false {
			externalEvents = append(externalEvents, e)
		}
	}

	return externalEvents
}

func (serv Service) HasDateField() bool {
	status := false
	for _, command := range serv.Commands {
		for _, attr := range command.Input.Attributes {
			if attr.Type == TypeTimestamp {
				status = true
				break
			}
		}
	}
	return status
}

func (serv Service) NameToLower() string {
	return strings.ToLower(serv.Name)
}

func (service Service) GetAllEvents() []Event {
	eventMap := make(map[string]Event)
	for _, command := range service.Commands {
		for _, event := range command.ConsumesEvents {
			eventMap[event.Name] = event
		}
		for _, event := range command.ProducesEvents {
			eventMap[event.Name] = event
		}
	}
	events := make([]Event, 0, 20)
	for _, event := range eventMap {
		events = append(events, event)
	}

	return events
}

// command
type Command struct {
	Name           string
	Method         MethodType
	Url            string
	Input          Entity
	ConsumesEvents []Event
	ProducesEvents []Event
}

func (c Command) AttributeNames() []string {
	nameList := make([]string, 0, 10)
	for _, entry := range c.Input.Attributes {
		nameList = append(nameList, entry.Name)
	}
	return nameList
}

func (c Command) NameToFirstUpper() string {
	return strings.Title(c.Name)
}

func (c Command) IsQuery() bool {
	return c.Method == Get
}

// entity
type Entity struct {
	Name               string
	Attributes         []Attribute
	AggregateName      string
	AggregateFieldName string
}

func (e Entity) NameToFirstUpper() string {
	return strings.Title(e.Name)
}

func (e Entity) AttributeNames() []string {
	nameList := make([]string, 0, 10)
	for _, entry := range e.Attributes {
		nameList = append(nameList, entry.Name)
	}
	return nameList
}

// event
type Event Entity

func (e Event) AttributeNames() []string {
	nameList := make([]string, 0, 10)
	for _, entry := range e.Attributes {
		nameList = append(nameList, entry.Name)
	}
	return nameList
}

func (e Event) NameToFirstUpper() string {
	return strings.Title(e.Name)
}

func (e Event) NameToFirstLower() string {
	return strings.ToLower(fmt.Sprintf("%c", e.Name[0])) + e.Name[1:]
}

func (e Event) AggregateFieldNameToFirstUpper() string {
	return strings.Title(e.AggregateFieldName)
}

func (e Event) getAggregateType() Type {
	aggregateType := TypeString

	for _, attr := range e.Attributes {
		if attr.Name == e.AggregateFieldName {
			aggregateType = attr.Type
			break
		}
	}
	return aggregateType
}

func (e Event) HasAggregateFieldTypeInt() bool {
	if e.getAggregateType() == TypeInt {
		return true
	}

	return false
}

// attribute
type Attribute struct {
	Name           string
	Type           Type
	Cardinality    CardinalityType
	MustMatchRegex string
}

func (attr Attribute) TypeName() string {
	if attr.Type == TypeString {
		return "string"
	} else if attr.Type == TypeInt {
		return "int"
	} else if attr.Type == TypeBoolean {
		return "bool"
	} else if attr.Type == TypeTimestamp {
		return "time.Time"
	} else {
		return "unknown"
	}
}

func (attr Attribute) MultiplicityNsme() string {
	if attr.Cardinality == Multiple {
		return "[]"
	} else {
		return ""
	}
}

func (attr Attribute) NameToFirstUpper() string {
	return strings.Title(attr.Name)
}

func (attr Attribute) NameToFirstLower() string {
	return strings.ToLower(fmt.Sprintf("%c", attr.Name[0])) + attr.Name[1:]
}
