package api

import (
	"encoding/json"
	"net/http"
	"renthome/boiler"
	"time"

	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (api *APIController) NotificationHandler(w http.ResponseWriter, r *http.Request) (int, error) {

	return http.StatusCreated, nil
}

type GetNotificationsRequest struct {
	UserID string `json:"user_id"`
}

type NotificationState struct {
	Notification *boiler.Notification `json:"notification"`
	State        string               `json:"state"`
}

type GetNotificationsResponse struct {
	Notifications []*NotificationState `json:"notifications"`
}

func (api *APIController) GetNotifications(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetNotificationsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	userNotifications, err := boiler.UserNotifications(boiler.UserNotificationWhere.UserID.EQ(req.UserID), qm.WithDeleted()).All(api.Conn)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrBadRequest)
	}

	notificationsResponse := &GetNotificationsResponse{}

	for _, userNotification := range userNotifications {
		notification, err := boiler.FindNotification(api.Conn, userNotification.NotificationID)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		notificationState := &NotificationState{
			Notification: notification,
			State:        userNotification.State,
		}

		notificationsResponse.Notifications = append(notificationsResponse.Notifications, notificationState)
	}

	if err = json.NewEncoder(w).Encode(notificationsResponse); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
}

func (api *APIController) UpdateNotificationHandler(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &NotificationState{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	userNotification, err := boiler.FindUserNotification(api.Conn, user.ID, req.Notification.ID)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrBadRequest)
	}

	userNotification.State = req.State
	userNotification.UpdatedAt = time.Now()

	if _, err = userNotification.Update(api.Conn, boil.Infer()); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	res := &NotificationState{
		Notification: req.Notification,
		State:        req.State,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	return http.StatusOK, nil
}
