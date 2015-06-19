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
			a.Err = errors.New(msg)
		}
	}
	return a.Err
}

func (a *Validator) MinSliceLength(name string, expectedLength int, value []int) error {
	if a.Err == nil {
		if len(value) != expectedLength {
			a.Err = fmt.Errorf("Invalid parameter %s", name)
		}
	}
	return a.Err
}

func (a *Validator) GreaterThan(name string, minEclusive int, value int) error {
	if a.Err == nil {
		if value <= minEclusive {
			a.Err = fmt.Errorf("Invalid parameter %s", name)
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
			a.Err = fmt.Errorf("Invalid parameter %s", name)
		}
	}
	return a.Err
}

func (a *Validator) NoDuplicates(name string, values []int) error {
	if a.Err == nil {
		if len(uniq(values)) < len(values) {
			a.Err = fmt.Errorf("%s contains duplicates", name)
		}
	}
	return a.Err
}

func uniq(list []int) []int {
	unique_set := make(map[int]int, len(list))
	i := 0
	for _, x := range list {
		if _, there := unique_set[x]; !there {
			unique_set[x] = i
			i++
		}
	}
	result := make([]int, len(unique_set))
	for x, i := range unique_set {
		result[i] = x
	}
	return result
}
