package internal

import (
	"database/sql"
	"net/http"
)

type Database = *sql.DB

type AuthenticationDetails struct {
	IsAdmin  bool
	IsDev    bool
	ID       string
	ClientId string
}

type ApiStatus struct {
	Version       string
	Status        string
	Introspection bool
	Playground    bool
}

type GraphQLHandler struct {
	Handler             http.Handler
	EnableIntrospection bool
}
