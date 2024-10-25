package server

import "net/http"

func GetAction(r *http.Request) string {
	action := r.URL.Query().Get("action")
	return action
}
