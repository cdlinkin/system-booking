package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/cdlinkin/system-booking/internal/service"
)

type UserHandler struct {
	userService    service.UserService
	bookingService service.BookingService
}

func NewUserHandler(userService service.UserService, bookingService service.BookingService) *UserHandler {
	return &UserHandler{
		userService:    userService,
		bookingService: bookingService,
	}
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var userDTO service.RegisterUserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.userService.Register(r.Context(), &userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) GetBookingsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 5 || pathParts[2] != "users" || pathParts[4] != "bookings" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userIDString := pathParts[3]
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookings, err := u.bookingService.GetByUserID(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bookings)
	w.WriteHeader(http.StatusOK)
}
