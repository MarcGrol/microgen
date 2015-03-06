package client

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

type Client struct {
	hostname string
	Err      error
}

func NewClient(hostname string) *Client {
	c := new(Client)
	c.hostname = hostname
	return c
}

type Response struct {
	Status bool             `json:"status"`
	Error  *ErrorDescriptor `json:"error"`
}

type ErrorDescriptor struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) CreateTour(year int) error {
	if c.Err == nil {
		log.Printf("Create tour %d", year)
		command := &tour.CreateTourCommand{Year: year}

		url := fmt.Sprintf("http://%s/api/tour", c.hostname)

		c.Err = doPost(url, command)
	}
	return c.Err
}

func (c *Client) CreateCyclist(year int, number int, name string, team string) error {
	if c.Err == nil {
		log.Printf("Create cyclist %s", name)
		command := &tour.CreateCyclistCommand{
			Year: year,
			Id:   number,
			Name: name,
			Team: team}

		url := fmt.Sprintf("http://%s/api/tour/%d/cyclist", c.hostname, year)

		c.Err = doPost(url, command)
	}
	return c.Err
}

func (c *Client) CreateEtappe(year int, number int, timestamp time.Time, start string, end string, length int, kind int) error {
	if c.Err == nil {
		log.Printf("Create etappe to %s", end)
		command := &tour.CreateEtappeCommand{
			Year:           year,
			Id:             number,
			Date:           timestamp,
			StartLocation:  start,
			FinishLocation: end,
			Length:         length,
			Kind:           kind}

		url := fmt.Sprintf("http://%s/api/tour/%d/etappe", c.hostname, year)

		c.Err = doPost(url, command)

	}
	return c.Err
}

func (c *Client) CreateGambler(gamblerUid string, name string, email string) error {
	if c.Err == nil {
		log.Printf("Create gambler %s", name)
		command := &gambler.CreateGamblerCommand{
			GamblerUid: gamblerUid,
			Name:       name,
			Email:      email}

		url := fmt.Sprintf("http://%s/api/gambler", c.hostname)

		c.Err = doPost(url, command)
	}
	return c.Err
}

func (c *Client) CreateGamblerTeam(year int, gamblerUid string, cyclistsIds []int) error {
	if c.Err == nil {
		log.Printf("Create team for gambler %s", gamblerUid)
		command := &gambler.CreateGamblerTeamCommand{
			GamblerUid: gamblerUid,
			Year:       year,
			CyclistIds: cyclistsIds}
		// gambler/:gamblerUid/year/:year/team

		url := fmt.Sprintf("http://%s/api/gambler/%s/year/%d/team", c.hostname, gamblerUid, year)

		c.Err = doPost(url, command)

	}
	return c.Err
}

func doPost(url string, command interface{}) error {
	// serialize request to json
	requestBody, err := encodeRequest(command)
	if err != nil {
		log.Printf("Error marshalling request %+v", err)
		return err
	}

	log.Printf("Posting on url '%s' <- %v", url, requestBody)

	// perform httc call
	httpResponse, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Printf("Error posting request %+v", err)
		return err
	}

	// evaluate response
	if httpResponse.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Http error: %d (%+v)", httpResponse.StatusCode, httpResponse.Status)
		log.Printf(errMsg)
		// err = errors.New(errMsg)
		// return err
	}

	// decode response
	dec := json.NewDecoder(httpResponse.Body)
	var applicationResponse Response
	err = dec.Decode(&applicationResponse)
	if err != nil {
		log.Printf("Error unmarshalling response: %+v", err)
		return err
	}

	// evaluate application status
	if applicationResponse.Status == false {
		errMsg := fmt.Sprintf("Applicative error: %s", applicationResponse.Error.Message)
		log.Printf(errMsg)
		err = errors.New(errMsg)
		return err
	}

	log.Printf("Succesfully performed POST on url %s", url)

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
