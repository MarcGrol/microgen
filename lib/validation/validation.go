package validation

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Validator struct {
	Err error
}

func (a *Validator) NotEmpty(name string, value string) error {
	if a.Err == nil {
		if value == "" {
			msg := fmt.Sprintf("Missing parameter %s", name)
			//log.Print(msg)
			a.Err = errors.New(msg)
		}
	}
	return a.Err
}

func (a *Validator) SliceLength(name string, expectedLength int, value []int) error {
	if a.Err == nil {
		if len(value) != expectedLength {
			msg := fmt.Sprintf("Invalid parameter %s", name)
			//log.Print(msg)
			a.Err = errors.New(msg)
		}
	}
	return a.Err
}

func (a *Validator) GreaterThan(name string, minEclusive int, value int) error {
	if a.Err == nil {
		if value <= minEclusive {
			msg := fmt.Sprintf("Invalid parameter %s", name)
			//log.Print(msg)
			a.Err = errors.New(msg)
		}
	}
	return a.Err
}

func (a *Validator) After(name string, minExclusive string, value time.Time) error {
	if a.Err == nil {
		minTime, err := time.Parse(time.RFC3339, minExclusive)
		if err != nil {
			log.Fatalf("**** Invalid min-exlusive date %s", err)
		}
		if value.Before(minTime) {
			msg := fmt.Sprintf("Invalid parameter %s", name)
			//log.Print(msg)
			a.Err = errors.New(msg)
		}
	}
	return a.Err
}
