package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"renthome/boiler"
	"strconv"

	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetLocationsRequest struct {
	SearchTerm  string   `json:"search_term"`
	IgnoreTerms []string `json:"ignore_terms"`
}

type GetLocationsResponse struct {
	Locations []string `json:"locations"`
	Total     int      `json:"total"`
}

func (api *APIController) GetLocations(w http.ResponseWriter, r *http.Request) (int, error) {
	_locations := []string{}
	isPostcodeSearch := false
	isTotalResultLessThanSeven := false

	req := &GetLocationsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	if req.SearchTerm == "" {
		return http.StatusBadRequest, fmt.Errorf("no search term provided")
	}

	if _, err := strconv.Atoi(req.SearchTerm); err == nil {
		isPostcodeSearch = true
	}

	if isPostcodeSearch {
		locations, err := boiler.Locations(qm.Where("postcode LIKE ?", fmt.Sprintf("%s%%", req.SearchTerm)),
			boiler.LocationWhere.Postcode.NIN(req.IgnoreTerms),
			boiler.LocationWhere.Description.NIN(req.IgnoreTerms),
			qm.Limit(7)).All(api.Conn)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, location := range locations {
			_locations = append(_locations, location.Postcode)
		}

		if len(_locations) < 7 {
			isTotalResultLessThanSeven = true
		}
	} else {
		locations, err := boiler.Locations(qm.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", req.SearchTerm)),
			boiler.LocationWhere.Postcode.NIN(req.IgnoreTerms),
			boiler.LocationWhere.Description.NIN(req.IgnoreTerms),
			qm.Limit(7)).All(api.Conn)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, location := range locations {
			_locations = append(_locations, location.Description)
		}
	}

	if isTotalResultLessThanSeven {
		locations, err := boiler.Locations(qm.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", req.SearchTerm)),
			boiler.LocationWhere.Postcode.NIN(req.IgnoreTerms),
			boiler.LocationWhere.Description.NIN(req.IgnoreTerms),
			qm.Limit(7)).All(api.Conn)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, location := range locations {
			if len(_locations) < 7 {
				_locations = append(_locations, location.Description)
			}
		}
	}

	locationsResponse := &GetLocationsResponse{
		Locations: _locations,
		Total:     len(_locations),
	}

	if err := json.NewEncoder(w).Encode(locationsResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
}
