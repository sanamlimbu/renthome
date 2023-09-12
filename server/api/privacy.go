package api

import (
	"encoding/json"
	"net/http"
	"renthome/boiler"
	"time"

	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GetPrivaciesRequest struct {
	UserID string `json:"user_id"`
}

type PrivacyState struct {
	Privacy *boiler.Privacy `json:"privacy"`
	State   string          `json:"state"`
}

type GetPrivaciesResponse struct {
	Privacies []*PrivacyState `json:"privacies"`
}

func (api *APIController) GetPrivacies(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetPrivaciesRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	privaciesResponse := &GetPrivaciesResponse{}

	userPrivacies, err := boiler.UserPrivacies(boiler.UserPrivacyWhere.UserID.EQ(req.UserID)).All(api.Conn)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrBadRequest)
	}

	for _, userPrivacy := range userPrivacies {
		privacy, err := boiler.FindPrivacy(api.Conn, userPrivacy.PrivacyID)
		if err != nil {
			return http.StatusBadRequest, terror.Error(err, ErrBadRequest)
		}

		privacyState := &PrivacyState{
			Privacy: privacy,
			State:   userPrivacy.State,
		}

		privaciesResponse.Privacies = append(privaciesResponse.Privacies, privacyState)
	}

	if err = json.NewEncoder(w).Encode(privaciesResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
}

func (api *APIController) UpdatePrivacyHandler(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &PrivacyState{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	userPrivacy, err := boiler.FindUserPrivacy(api.Conn, user.ID, req.Privacy.ID)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrBadRequest)
	}

	userPrivacy.State = req.State
	userPrivacy.UpdatedAt = time.Now()

	if _, err = userPrivacy.Update(api.Conn, boil.Infer()); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	res := &PrivacyState{
		Privacy: req.Privacy,
		State:   req.State,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
}
