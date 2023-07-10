package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"renthome/boiler"
	"time"

	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Filter struct {
	All               bool   `json:"all"`
	House             bool   `json:"house"`
	Apartment         bool   `json:"apartment"`
	Unit              bool   `json:"unit"`
	TownHouse         bool   `json:"town_house"`
	Villa             bool   `json:"villa"`
	PriceMin          string `json:"price_min"`
	PriceMax          string `json:"price_max"`
	BedMin            string `json:"bed_min"`
	BedMax            string `json:"bed_max"`
	BathroomCount     string `json:"bathroom_count"`
	CarCount          string `json:"car_count"`
	AvailableDate     string `json:"available_date"`
	IsFurnished       bool   `json:"is_furnished"`
	IsPetsConsidered  bool   `json:"is_pets_conisdereed"`
	HasAirConditioner bool   `json:"has_air_conditioner"`
	HasDishwasher     bool   `json:"has_dishwasher"`
}
type GetPropertiesRequest struct {
	Search string `json:"search"`
	Filter Filter `json:"filter"`
}

func (api *APIController) GetProperties(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

func (api *APIController) GetProperty(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

type CreatePropertyRequest struct {
	Type             string `json:"type" validate:"required"`
	Category         string `json:"category" validate:"required"`
	Street           string `json:"street" validate:"required"`
	Suburb           string `json:"suburb" validate:"required"`
	Postcode         int    `json:"postcode" validate:"required,gte=0"`
	State            string `json:"state" validate:"required"`
	BedCount         int    `json:"bed_count" validate:"required"`
	BathCount        int    `json:"bath_count" validate:"required"`
	CarCount         int    `json:"car_count" validate:"required"`
	HasAirCon        bool   `json:"has_aircon" validate:"required"`
	IsFurnished      bool   `json:"is_furnished" validate:"required"`
	IsPetsConsidered bool   `json:"is_pets_considered" validate:"required"`
	AvailableAt      string `json:"available_at" validate:"required"`
	OpenAt           string `json:"open_at" validate:"required"`
	Price            int    `json:"price" validate:"required"`
}

func (api *APIController) CreateProperty(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &CreatePropertyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	availableAt, err := time.Parse(time.RFC3339, req.AvailableAt)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	}

	openAt, err := time.Parse(time.RFC3339, req.AvailableAt)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	}

	property := &boiler.Property{
		Type:             req.Type,
		Category:         req.Category,
		Street:           req.Street,
		Suburb:           req.Suburb,
		Postcode:         req.Postcode,
		State:            req.State,
		BedCount:         req.BedCount,
		BathCount:        req.BathCount,
		CarCount:         req.CarCount,
		HasAircon:        req.HasAirCon,
		IsFurnished:      req.IsFurnished,
		IsPetsConsidered: req.IsPetsConsidered,
		AvailableAt:      null.TimeFrom(availableAt),
		OpenAt:           null.TimeFrom(openAt),
		Price:            req.Price,
	}

	err = property.Insert(api.Conn, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to create property.")
	}

	return 200, nil
}

func (api *APIController) UpdateProperty(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

func (api *APIController) DeleteProperty(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	return 200, nil

}

type TestRequest struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

func (api *APIController) Test(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &TestRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, err.Error())
	}

	err = api.Validator.Validate.Struct(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, err.Error())
	}

	fmt.Println(req)

	return 200, nil

}
