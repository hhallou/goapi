package healthz

import "net/http"

//Handler ...
type Handler struct {
}

//Return a new Handler
func New() Handler {
	return Handler{}
}

//Handler ...
func (h Handler) getHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
