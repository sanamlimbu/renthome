package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"renthome/boiler"
	"time"

	"github.com/h2non/filetype"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PropertyType string

const (
	Unit      PropertyType = "Unit"
	House     PropertyType = "House"
	Villa     PropertyType = "Villa"
	Townhouse PropertyType = "Townhouse"
	Apartment PropertyType = "Apartment"
)

type AvailableDateCondition string

const (
	AvailableDateBefore AvailableDateCondition = "Before"
	AvailableDateAfter  AvailableDateCondition = "After"
	AvailableDateAt     AvailableDateCondition = "At"
)

type AvailableDate struct {
	Date      time.Time              `json:"date"`
	Condition AvailableDateCondition `json:"condition"`
}

type SearchFilter struct {
	PropertyTypes     []string      `json:"property_types"`
	PropertyTypesAny  bool          `json:"property_types_any"`
	PriceMin          int           `json:"price_min"`
	PriceMax          int           `json:"price_max"`
	PriceMinAny       bool          `json:"price_min_any"`
	PriceMaxAny       bool          `json:"price_max_any"`
	BedMin            int           `json:"bed_min"`
	BedMax            int           `json:"bed_max"`
	BedMinAny         bool          `json:"bed_min_any"`
	BedMaxAny         bool          `json:"bed_max_any"`
	BathroomCount     int           `json:"bathroom_count"`
	BathroomCountAny  bool          `json:"bathroom_count_any"`
	CarCount          int           `json:"car_count"`
	CarCountAny       bool          `json:"car_count_any"`
	AvailableDate     AvailableDate `json:"available_date"`
	AvailableDateAny  bool          `json:"available_date_any"`
	AvailableNow      bool          `json:"available_now"`
	IsFurnished       bool          `json:"is_furnished"`
	IsPetsConsidered  bool          `json:"is_pets_conisdereed"`
	HasAirConditioner bool          `json:"has_air_conditioner"`
	HasDishwasher     bool          `json:"has_dishwasher"`
}
type GetPropertiesRequest struct {
	Locations []string     `json:"locations"`
	Filter    SearchFilter `json:"filter"`
}

type GetPropertiesResponse struct {
	Properties []*boiler.Property `json:"properties"`
	Total      int                `json:"total"`
}

func (api *APIController) GetProperties(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	var properties boiler.PropertySlice
	var resp *GetPropertiesResponse
	var conditions []qm.QueryMod

	// Property type
	if req.Filter.PropertyTypesAny {
		propertyTypes := []string{
			string(Unit),
			string(House),
			string(Villa),
			string(Townhouse),
			string(Apartment),
		}
		conditions = append(conditions, boiler.PropertyWhere.Type.IN(propertyTypes))
	} else {
		conditions = append(conditions, boiler.PropertyWhere.Type.IN(req.Filter.PropertyTypes))
	}

	// Price
	if !req.Filter.PriceMinAny {
		conditions = append(conditions, boiler.PropertyWhere.Price.GTE(req.Filter.PriceMin))

	}
	if !req.Filter.PriceMaxAny {
		conditions = append(conditions, boiler.PropertyWhere.Price.LTE(req.Filter.PriceMax))
	}

	// Bedroom
	if !req.Filter.BedMinAny {
		conditions = append(conditions, boiler.PropertyWhere.BedCount.GTE(req.Filter.BedMin))
	}
	if !req.Filter.BedMaxAny {
		conditions = append(conditions, boiler.PropertyWhere.BedCount.LTE(req.Filter.BedMax))
	}

	//Bathroom
	if !req.Filter.BathroomCountAny {
		conditions = append(conditions, boiler.PropertyWhere.BathCount.GTE(req.Filter.BathroomCount))
	}

	// Car
	if !req.Filter.CarCountAny {
		conditions = append(conditions, boiler.PropertyWhere.CarCount.GTE(req.Filter.CarCount))
	}

	// Available date
	if req.Filter.AvailableNow {
		conditions = append(conditions, boiler.PropertyWhere.IsAvailableNow.EQ(true))
	} else {
		if !req.Filter.AvailableDateAny {
			if req.Filter.AvailableDate.Condition == AvailableDateBefore {
				conditions = append(conditions, qm.Where("date_trunc('day', ?) < date_trunc('day', ?)", boiler.PropertyColumns.AvailableAt, req.Filter.AvailableDate.Date))
			} else if req.Filter.AvailableDate.Condition == AvailableDateAfter {
				conditions = append(conditions, qm.Where("date_trunc('day', ?) > date_trunc('day', ?)", boiler.PropertyColumns.AvailableAt, req.Filter.AvailableDate.Date))
			} else {
				conditions = append(conditions, qm.Where("date_trunc('day', ?) = date_trunc('day', ?)", boiler.PropertyColumns.AvailableAt, req.Filter.AvailableDate.Date))
			}
		}
	}

	// Furnished
	conditions = append(conditions, boiler.PropertyWhere.IsFurnished.EQ(req.Filter.IsFurnished))

	// Pets
	conditions = append(conditions, boiler.PropertyWhere.IsPetsConsidered.EQ(req.Filter.IsPetsConsidered))

	// Aircon
	conditions = append(conditions, boiler.PropertyWhere.HasAircon.EQ(req.Filter.HasAirConditioner))

	// Dishwasher
	conditions = append(conditions, boiler.PropertyWhere.HasDishwasher.EQ(req.Filter.HasAirConditioner))

	// Postcode
	conditions = append(conditions, boiler.PropertyWhere.Postcode.IN(req.Locations))

	// Location
	conditions = append(conditions, boiler.PropertyWhere.Location.IN(req.Locations))

	properties, err = boiler.Properties(conditions...).All(api.Conn)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	resp = &GetPropertiesResponse{
		Properties: properties,
		Total:      len(properties),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
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
	Type             string    `json:"type" validate:"required"`
	Category         string    `json:"category" validate:"required"`
	Street           string    `json:"street" validate:"required"`
	Suburb           string    `json:"suburb" validate:"required"`
	Postcode         string    `json:"postcode" validate:"required"`
	State            string    `json:"state" validate:"required"`
	BedCount         int       `json:"bed_count" validate:"required"`
	BathCount        int       `json:"bath_count" validate:"required"`
	CarCount         int       `json:"car_count" validate:"required"`
	HasAirCon        bool      `json:"has_aircon" validate:"required"`
	IsFurnished      bool      `json:"is_furnished" validate:"required"`
	IsPetsConsidered bool      `json:"is_pets_considered" validate:"required"`
	AvailableAt      null.Time `json:"available_at" validate:""`
	OpenAt           null.Time `json:"open_at" validate:""`
	Price            int       `json:"price" validate:"required"`
}

type CreatePropertyResponse struct {
	Property *boiler.Property `json:"property"`
}

func (api *APIController) CreateProperty(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &CreatePropertyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	// check of zero value

	// if req.AvailableAt.Valid {
	// 	availableAt, err = time.Parse(time.RFC3339, req.AvailableAt)
	// 	if err != nil {
	// 		return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	// 	}
	// }

	// openAt, err := time.Parse(time.RFC3339, req.AvailableAt.String())
	// if err != nil {
	// 	return http.StatusBadRequest, terror.Error(err, "Invalid date.")
	// }

	// propertyID, err := uuid.NewV4()
	// if err != nil {
	// 	return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	// }

	fmt.Println("available", req.AvailableAt)
	fmt.Println("Open", req.OpenAt)

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
		AvailableAt:      req.AvailableAt,
		OpenAt:           req.OpenAt,
		Price:            req.Price,
		ManagerID:        "90b71c18-c836-421b-9e17-0bb119019baa",
		AgencyID:         "5d621a17-6ea0-430a-98f0-ea419097c751",
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	err = property.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to create property.")
	}

	// for _, image := range images {
	// 	image.PropertyID = propertyID.String()
	// 	image.UploaderID = propertyID.String()

	// 	image.Insert(tx, boil.Infer())
	// 	if err != nil {
	// 		return http.StatusInternalServerError, terror.Error(err, "Unable to create property.")
	// 	}
	// }

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	createPropertyResponse := &CreatePropertyResponse{
		Property: property,
	}

	if err = json.NewEncoder(w).Encode(createPropertyResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusCreated, nil
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

// parseImages will read a multipart form request and returns slice of Image
func parseImages(r *http.Request) ([]*boiler.Image, error) {

	var images []*boiler.Image

	multipartReader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, err
		}

		// handle file
		if part.FormName() == "file" {
			// get mime type
			kind, err := filetype.Match(data)
			if err != nil {
				return nil, err
			}

			mimeType := kind.MIME.Value
			extension := kind.Extension

			if kind == filetype.Unknown {
				return nil, err
			}

			image := &boiler.Image{
				FileSizeBytes: int64(len(data)),
				Extension:     extension,
				MimeType:      mimeType,
			}

			images = append(images, image)

		}
	}
	return images, nil
}
