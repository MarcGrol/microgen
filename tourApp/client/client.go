package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MarcGrol/microgen/lib/myerrors"
	"github.com/MarcGrol/microgen/tourApp/gambler"
	"github.com/MarcGrol/microgen/tourApp/tour"
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

		url := fmt.Sprintf("%s/api/tour", c.hostname)

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

		url := fmt.Sprintf("%s/api/tour/%d/cyclist", c.hostname, year)

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

		url := fmt.Sprintf("%s/api/tour/%d/etappe", c.hostname, year)

		c.Err = doPost(url, command)

	}
	return c.Err
}

func (c *Client) CreateEtappeResults(year int, etappeId int, bestDayCyclistIds []int,
	bestAllrondersCyclistIds []int, bestSprintersCyclistIds []int, bestClimberCyclistIds []int) error {
	if c.Err == nil {
		log.Printf("Create etappe results for etappe %d", etappeId)
		command := &tour.CreateEtappeResultsCommand{
			Year:                   year,
			EtappeId:               etappeId,
			BestDayCyclistIds:      bestDayCyclistIds,
			BestAllroundCyclistIds: bestAllrondersCyclistIds,
			BestClimbCyclistIds:    bestClimberCyclistIds,
			BestSprintCyclistIds:   bestSprintersCyclistIds}

		url := fmt.Sprintf("%s/api/tour/%d/etappe/%d", c.hostname, year, etappeId)

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

		url := fmt.Sprintf("%s/api/gambler", c.hostname)

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

		url := fmt.Sprintf("%s/api/gambler/%s/year/%d/team", c.hostname, gamblerUid, year)

		c.Err = doPost(url, command)

	}
	return c.Err
}

func doPost(url string, command interface{}) error {
	// serialize request to json
	requestBody, err := encodeRequest(command)
	if err != nil {
		return fmt.Errorf("Error unmarshalling request: %s", err.Error())
	}

	log.Printf("Posting on url '%s' <- %v", url, requestBody)

	// perform http call
	httpResponse, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return fmt.Errorf("Error performing http POST: %s", err.Error())
	}

	// decode response
	applicationResponse, err := decodeResponse(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("Error unmarshalling response: %s", err.Error())
	}

	// evaluate application status
	if httpResponse.StatusCode != http.StatusOK {
		errMsg := "unknown error"
		if applicationResponse.Error != nil {
			errMsg = fmt.Sprintf(applicationResponse.Error.Message)
		}
		if httpResponse.StatusCode == http.StatusInternalServerError {
			return myerrors.NewInternalError(errors.New(errMsg))
		} else if httpResponse.StatusCode == http.StatusNotFound {
			return myerrors.NewNotFoundError(errors.New(errMsg))
		} else if httpResponse.StatusCode == http.StatusBadRequest {
			return myerrors.NewInvalidInputError(errors.New(errMsg))
		} else if httpResponse.StatusCode == http.StatusUnauthorized {
			return myerrors.NewNotAuthorizedError(errors.New(errMsg))
		} else {
			return errors.New(errMsg)
		}
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

func decodeResponse(responseBody io.ReadCloser) (*Response, error) {
	dec := json.NewDecoder(responseBody)
	var applicationResponse Response
	err := dec.Decode(&applicationResponse)
	if err != nil {
		return nil, err
	}
	return &applicationResponse, nil
}
