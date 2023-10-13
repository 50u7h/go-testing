package main

import (
	"net/http"
)

// Credentials is the type used to unmarshal a JSON payload
// during authentication.
type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// authenticate is the handler used to try to authenticate a user, and
// send them a JWT token if successful.
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

}

// refresh is the handler called to request a new token pair, when
// the jwt token has expired. We expect the refresh token to come
// from a POST request. We validate it, look up the user in the db,
// and if everything is good we send back a new token pair
// as JSON. We also set an http only, secure cookie with the refresh
// token stored inside.
func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

}

// allUsers returns a list of all users as JSON
func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {

}

// getUser returns one user as JSON
func (app *application) getUser(w http.ResponseWriter, r *http.Request) {

}

// updateUser updates a user from a JSON payload, and returns just a header
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

}

// deleteUser deletes one user based on ID in URL, and returns a header
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

}

// insertUser inserts a user using a JSON payload, and returns a header
func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {

}
