/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"CreateAd",
			strings.ToUpper("Post"),
			"/ads",
			c.CreateAd,
		},
		{
			"CreateClient",
			strings.ToUpper("Post"),
			"/client",
			c.CreateClient,
		},
		{
			"CreateEvent",
			strings.ToUpper("Post"),
			"/events",
			c.CreateEvent,
		},
		{
			"CreateEventType",
			strings.ToUpper("Post"),
			"/event_types",
			c.CreateEventType,
		},
		{
			"DeleteClient",
			strings.ToUpper("Delete"),
			"/client",
			c.DeleteClient,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/user",
			c.DeleteUser,
		},
		{
			"Filter",
			strings.ToUpper("Post"),
			"/filter",
			c.Filter,
		},
		{
			"GetAds",
			strings.ToUpper("Get"),
			"/ads",
			c.GetAds,
		},
		{
			"GetClient",
			strings.ToUpper("Get"),
			"/client",
			c.GetClient,
		},
		{
			"GetClientStats",
			strings.ToUpper("Get"),
			"/stats",
			c.GetClientStats,
		},
		{
			"GetClients",
			strings.ToUpper("Get"),
			"/clients",
			c.GetClients,
		},
		{
			"GetEventTypes",
			strings.ToUpper("Get"),
			"/event_types",
			c.GetEventTypes,
		},
		{
			"GetUser",
			strings.ToUpper("Get"),
			"/user",
			c.GetUser,
		},
		{
			"GetUserByID",
			strings.ToUpper("Get"),
			"/user/id",
			c.GetUserByID,
		},
		{
			"GetUsers",
			strings.ToUpper("Get"),
			"/users",
			c.GetUsers,
		},
		{
			"GrantUserAdmin",
			strings.ToUpper("Put"),
			"/user",
			c.GrantUserAdmin,
		},
		{
			"Login",
			strings.ToUpper("Post"),
			"/login",
			c.Login,
		},
		{
			"Register",
			strings.ToUpper("Post"),
			"/register",
			c.Register,
		},
		{
			"UserInfo",
			strings.ToUpper("Get"),
			"/users/me",
			c.UserInfo,
		},
	}
}

// CreateAd - Создание рекламной рассылки
func (c *DefaultApiController) CreateAd(w http.ResponseWriter, r *http.Request) {
	createAdRequestParam := CreateAdRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createAdRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateAdRequestRequired(createAdRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateAd(r.Context(), createAdRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateClient - Регистрация клиента системы
func (c *DefaultApiController) CreateClient(w http.ResponseWriter, r *http.Request) {
	createClientRequestParam := CreateClientRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createClientRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateClientRequestRequired(createClientRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateClient(r.Context(), createClientRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateEvent - Добавление информации о событии
func (c *DefaultApiController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	createEventRequestParam := CreateEventRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createEventRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateEventRequestRequired(createEventRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEvent(r.Context(), createEventRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateEventType - Добавление типа события
func (c *DefaultApiController) CreateEventType(w http.ResponseWriter, r *http.Request) {
	userTokenParam := r.Header.Get("User-Token")
	createEventTypeRequestParam := CreateEventTypeRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createEventTypeRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateEventTypeRequestRequired(createEventTypeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEventType(r.Context(), userTokenParam, createEventTypeRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteClient - Удаление клиента из системы
func (c *DefaultApiController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	deleteClientRequestParam := DeleteClientRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deleteClientRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeleteClientRequestRequired(deleteClientRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DeleteClient(r.Context(), deleteClientRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteUser - Удаление пользователя
func (c *DefaultApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	deleteUserRequestParam := DeleteUserRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deleteUserRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeleteUserRequestRequired(deleteUserRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DeleteUser(r.Context(), deleteUserRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Filter - Фильтрация событий
func (c *DefaultApiController) Filter(w http.ResponseWriter, r *http.Request) {
	filterRequestParam := FilterRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&filterRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertFilterRequestRequired(filterRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Filter(r.Context(), filterRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAds - Получение рекламных рассылок
func (c *DefaultApiController) GetAds(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAds(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetClient - Получение информации о клиенте системы
func (c *DefaultApiController) GetClient(w http.ResponseWriter, r *http.Request) {
	getClientRequestParam := GetClientRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&getClientRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGetClientRequestRequired(getClientRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GetClient(r.Context(), getClientRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetClientStats - Получение статистики системы
func (c *DefaultApiController) GetClientStats(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetClientStats(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetClients - Получение информации о всех клиентах системы
func (c *DefaultApiController) GetClients(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetClients(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEventTypes - Получение типов событий
func (c *DefaultApiController) GetEventTypes(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetEventTypes(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUser - Получение информации о пользователе по логину
func (c *DefaultApiController) GetUser(w http.ResponseWriter, r *http.Request) {
	getUserRequestParam := GetUserRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&getUserRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGetUserRequestRequired(getUserRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GetUser(r.Context(), getUserRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUserByID - Получение информации о пользователе по ID
func (c *DefaultApiController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	getUserByIdRequestParam := GetUserByIdRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&getUserByIdRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGetUserByIdRequestRequired(getUserByIdRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GetUserByID(r.Context(), getUserByIdRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsers - Получение информации о всех пользователях системы
func (c *DefaultApiController) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetUsers(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GrantUserAdmin - Выдача пользователю прав администратора
func (c *DefaultApiController) GrantUserAdmin(w http.ResponseWriter, r *http.Request) {
	grantAdminRequestParam := GrantAdminRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&grantAdminRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGrantAdminRequestRequired(grantAdminRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GrantUserAdmin(r.Context(), grantAdminRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Login - Вход пользователя в систему
func (c *DefaultApiController) Login(w http.ResponseWriter, r *http.Request) {
	authRequestParam := AuthRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&authRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAuthRequestRequired(authRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Login(r.Context(), authRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Register - Регистрация пользователя
func (c *DefaultApiController) Register(w http.ResponseWriter, r *http.Request) {
	authRequestParam := AuthRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&authRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAuthRequestRequired(authRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Register(r.Context(), authRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserInfo - Получение информации о текущем пользователе
func (c *DefaultApiController) UserInfo(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserInfo(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
