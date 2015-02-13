package spec

import (
	"strings"
	"fmt"
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
	Get MethodType = iota
	Put
	Post
	Delete
)

type Application struct {
	Name     string
	Services []Service
}

type Service struct {
	Name     string
	Commands []Command
}

type Command struct {
	Name           string
	Method         MethodType
	Url            string
	Input          Entity
	ConsumesEvents []Event
	ProducesEvents []Event
}

type Entity struct {
	Name       string
	Attributes []Attribute
}
type Event Entity

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

func (e Entity) NameToFirstUpper() string {
	return strings.Title(e.Name)
}

func (e Event) NameToFirstUpper() string {
	return strings.Title(e.Name)
}

func (e Event) NameToFirstLower() string {
	return strings.ToLower(fmt.Sprintf("%c", e.Name[0])) + e.Name[1:]
}

func (c Command) NameToFirstUpper() string {
	return strings.Title(c.Name)
}

func (attr Attribute) NameToFirstUpper() string {
	return strings.Title(attr.Name)
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

func (service Service) GetConsumedEvents() []Event {
	eventMap := make(map[string]Event)
	for _, command := range service.Commands {
		for _, event := range command.ConsumesEvents {
			eventMap[event.Name] = event
		}
	}
	events := make([]Event, 0, 20)
	for _, event := range eventMap {
		events = append(events, event)
	}

	return events
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

func (app Application) GetEvents() []Event {
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
