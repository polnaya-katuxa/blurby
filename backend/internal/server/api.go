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
	"context"
	"net/http"
)

// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface {
	CreateAd(http.ResponseWriter, *http.Request)
	CreateClient(http.ResponseWriter, *http.Request)
	CreateEvent(http.ResponseWriter, *http.Request)
	CreateEventType(http.ResponseWriter, *http.Request)
	DeleteClient(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	Filter(http.ResponseWriter, *http.Request)
	GetAds(http.ResponseWriter, *http.Request)
	GetClient(http.ResponseWriter, *http.Request)
	GetClientStats(http.ResponseWriter, *http.Request)
	GetClients(http.ResponseWriter, *http.Request)
	GetEventTypes(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GrantUserAdmin(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	UserInfo(http.ResponseWriter, *http.Request)
}

// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface {
	CreateAd(context.Context, CreateAdRequest) (ImplResponse, error)
	CreateClient(context.Context, CreateClientRequest) (ImplResponse, error)
	CreateEvent(context.Context, CreateEventRequest) (ImplResponse, error)
	CreateEventType(context.Context, string, CreateEventTypeRequest) (ImplResponse, error)
	DeleteClient(context.Context, DeleteClientRequest) (ImplResponse, error)
	DeleteUser(context.Context, DeleteUserRequest) (ImplResponse, error)
	Filter(context.Context, FilterRequest) (ImplResponse, error)
	GetAds(context.Context) (ImplResponse, error)
	GetClient(context.Context, GetClientRequest) (ImplResponse, error)
	GetClientStats(context.Context) (ImplResponse, error)
	GetClients(context.Context) (ImplResponse, error)
	GetEventTypes(context.Context) (ImplResponse, error)
	GetUser(context.Context, GetUserRequest) (ImplResponse, error)
	GetUserByID(context.Context, GetUserByIdRequest) (ImplResponse, error)
	GetUsers(context.Context) (ImplResponse, error)
	GrantUserAdmin(context.Context, GrantAdminRequest) (ImplResponse, error)
	Login(context.Context, AuthRequest) (ImplResponse, error)
	Register(context.Context, AuthRequest) (ImplResponse, error)
	UserInfo(context.Context) (ImplResponse, error)
}
