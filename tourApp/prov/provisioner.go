package prov

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarcGrol/microgen/tourApp/gambler"
	"github.com/MarcGrol/microgen/tourApp/tour"
	"log"
	"net/http"
	"time"
)

type EtappeKind int

const (
	Flat      = 1
	Hilly     = 2
	Mountains = 3
	TimeTrial = 4
)

type Provisioner struct {
	hostname string
	err      error
}

func NewProvisioner(hostname string) *Provisioner {
	p := new(Provisioner)
	p.hostname = hostname
	return p
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 9, 0, 0, 0, time.Local)
}

type Response struct {
	Status bool             `json:"status"`
	Error  *ErrorDescriptor `json:"error"`
}

type ErrorDescriptor struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (p *Provisioner) CreateTour(year int) error {
	if p.err != nil {
		log.Printf("Create tour %d", year)
		command := &tour.CreateTourCommand{Year: year}

		url := fmt.Sprintf("http://%s/api/tour/%d", p.hostname, year)

		p.err = doPost(url, command)
	}
	return p.err
}

func (p *Provisioner) CreateCyclist(year int, number int, name string, team string) error {
	if p.err == nil {
		log.Printf("Create cyclist %s", name)
		command := &tour.CreateCyclistCommand{
			Year: year,
			Id:   number,
			Name: name,
			Team: team}

		url := fmt.Sprintf("http://%s/api/tour/%d/cyclist", p.hostname, year)

		p.err = doPost(url, command)
	}
	return p.err
}

func (p *Provisioner) CreateEtappe(year int, number int, timestamp time.Time, start string, end string, length int, kind int) error {
	if p.err == nil {
		log.Printf("Create etappe to %s", end)
		command := &tour.CreateEtappeCommand{
			Year:           year,
			Id:             number,
			Date:           timestamp,
			StartLocation:  start,
			FinishLocation: end,
			Length:         length,
			Kind:           kind}

		url := fmt.Sprintf("http://%s/api/tour/%d/etappe", p.hostname, year)

		p.err = doPost(url, command)

	}
	return p.err
}

func (p *Provisioner) CreateGambler(gamblerUid string, name string, email string) error {
	if p.err == nil {
		log.Printf("Create gambler %s", name)
		command := &gambler.CreateGamblerCommand{
			GamblerUid: gamblerUid,
			Name:       name,
			Email:      email}

		url := fmt.Sprintf("http://%s/api/gambler", p.hostname)

		p.err = doPost(url, command)
	}
	return p.err
}

func (p *Provisioner) CreateGamblerTeam(year int, gamblerUid string, cyclistsIds []int) error {
	if p.err == nil {
		log.Printf("Create team for gambler %s", gamblerUid)
		command := &gambler.CreateGamblerTeamCommand{
			GamblerUid: gamblerUid,
			Year:       year,
			CyclistIds: cyclistsIds}
		// gambler/:gamblerUid/year/:year/team

		url := fmt.Sprintf("http://%s/api/gambler/%s/year/%d/team", p.hostname, gamblerUid, year)

		p.err = doPost(url, command)

	}
	return p.err
}

func doPost(url string, command interface{}) error {
	// serialize request to json
	requestBody, err := encodeRequest(command)
	if err != nil {
		log.Printf("Error marshalling request %+v", err)
		return err
	}

	log.Printf("Posting on url %s:%v", url, requestBody)

	// perform http call
	httpResponse, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Printf("Error posting request %+v", err)
		return err
	}

	// evaluate response
	if httpResponse.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Http error creating tour: %d", httpResponse.Status)
		log.Printf(errMsg)
		err = errors.New(errMsg)
		return err
	}

	// decode response
	dec := json.NewDecoder(httpResponse.Body)
	var applicationResponse Response
	err = dec.Decode(applicationResponse)
	if err != nil {
		log.Printf("Error unmarshalling response %+v", err)
		return err
	}

	// evaluate application status
	if applicationResponse.Status == false {
		errMsg := fmt.Sprintf("Error creating tour: %s", applicationResponse.Error.Message)
		log.Printf(errMsg)
		err = errors.New(errMsg)
		return err
	}

	log.Printf("Succesfully perform POST on url %s", url)

	return nil
}

func encodeRequest(command interface{}) (*bytes.Buffer, error) {
	var requestBody bytes.Buffer
	enc := json.NewEncoder(&requestBody)
	err := enc.Encode(command)
	if err != nil {
		return nil, err
	}
	return &requestBody, nil
}
