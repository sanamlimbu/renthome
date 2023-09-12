package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"renthome/boiler"
	"strconv"
	"strings"

	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetLocationsResponse struct {
	Locations []string `json:"locations"`
	Total     int      `json:"total"`
}

func (api *APIController) GetLocations(w http.ResponseWriter, r *http.Request) (int, error) {
	_locations := []string{}
	isPostcodeSearch := false
	isTotalResultLessThanSeven := false

	searchTerm := strings.ToLower(r.URL.Query().Get("search_term"))
	if searchTerm == "" {
		return 200, nil
	}

	if _, err := strconv.Atoi(searchTerm); err == nil {
		isPostcodeSearch = true
	}

	if isPostcodeSearch {
		locations, err := boiler.Locations(qm.Where("postcode LIKE ?", fmt.Sprintf("%s%%", searchTerm)),
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
		locations, err := boiler.Locations(qm.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)),
			qm.Limit(7)).All(api.Conn)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, location := range locations {
			_locations = append(_locations, location.Description)
		}
	}

	if isTotalResultLessThanSeven {
		locations, err := boiler.Locations(qm.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)),
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
