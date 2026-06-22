package api

import (
	"encoding/json"
	"go-onboarding/storage"
	"net/http"
	"strconv"
)

type UserHandler struct {
	accessor storage.UserAccessor
}

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func NewUserHandler(ua storage.UserAccessor) *UserHandler {
	return &UserHandler{accessor: ua}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}
	token, err := GenerateToken(10, "Dee")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to issue authentication token"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		idParam := r.URL.Query().Get("id")
		if idParam == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "id is required"}`))
			return
		}

		targerID, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid id"}`))
			return
		}

		viewType := r.URL.Query().Get("type")
		if viewType == "full" {
			fullProfile, found := h.accessor.GetFullProfile(targerID)
			if !found {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"error": "user not found"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(fullProfile)
			return

		}

		user, ok := h.accessor.Get(targerID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "user not found"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	case http.MethodPost:
		var incomingUser storage.User
		//decode the raw json body into the struct
		err := json.NewDecoder(r.Body).Decode(&incomingUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid request body"}`))
			return
		}

		//save the user to the db
		err = h.accessor.Save(incomingUser.ID, incomingUser.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "failed to save user"}`))
			return
		}

		//respone with the user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(incomingUser)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "method not allowed"}`))
	}
}
