package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tourCreatedX = Event{
		Name: "TourCreated",
		Attributes: []Attribute{
			{Name: "y", Type: TypeInt, Cardinality: Mandatory},
			{Name: "y", Type: TypeInt, Cardinality: Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}
	tourCreatedY = Event{
		Name: "TourCreated",
		Attributes: []Attribute{
			{Name: "y", Type: TypeInt, Cardinality: Mandatory},
			{Name: "y", Type: TypeInt, Cardinality: Mandatory},
		},
		AggregateName:      "tour",
		AggregateFieldName: "year",
	}

	testAplication = Application{
		Name: "tourApp",
		Services: []Service{
			{
				Name: "Tour",
				Commands: []Command{
					{
						Name:   "CreateTour",
						Method: Post,
						Url:    "/tour",
						Input: Entity{
							Attributes: []Attribute{
								{Name: "year", Type: TypeInt, Cardinality: Mandatory},
								{Name: "year", Type: TypeInt, Cardinality: Mandatory},
								{},
							},
						},
						ConsumesEvents: []Event{tourCreatedX, tourCreatedY},
						ProducesEvents: []Event{tourCreatedX, tourCreatedY},
					},
					{
						Name:   "CreateTour",
						Method: Post,
						Url:    "/tour/:year/cyclist",
						Input: Entity{
							Attributes: []Attribute{},
						},
						ConsumesEvents: []Event{tourCreatedX, tourCreatedX},
						ProducesEvents: []Event{tourCreatedY, tourCreatedY},
					},
					{},
				},
			},
			{
				Name: "Tour",
			},
			{},
		},
	}
)

func TestValidate(t *testing.T) {
	// write and close
	err := ValidateApplication(testAplication)
	assert.NotNil(t, err)
}
