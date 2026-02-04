package service

import (
	"context"
	"errors"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/cdlinkin/system-booking/internal/repo"
)

type BookingService interface {
	CreateBooking(ctx context.Context, input *BookingDTO) error
	GetByID(ctx context.Context, id int) (*model.Booking, error)
	GetByUserID(ctx context.Context, userID int) ([]*model.Booking, error)
}

type bookingService struct {
	bookingRepo repo.BookingRepo
}

func NewBookingService(bookingRepo repo.BookingRepo) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

type BookingDTO struct {
	UserID     int `json:"user_id"`
	ResourceID int `json:"resource_id"`
}

func (b *bookingService) CreateBooking(ctx context.Context, input *BookingDTO) error {
	if input.UserID == 0 {
		return errors.New("некорректный user_id")
	}

	if input.ResourceID <= 0 {
		return errors.New("некорректный resource_id")
	}

	booking := &model.Booking{
		UserID:     input.UserID,
		ResourceID: input.ResourceID,
		Status:     model.BookingStatusPending,
	}

	if err := booking.Validate(); err != nil {
		return err
	}

	if err := b.bookingRepo.Create(ctx, booking); err != nil {
		return err
	}

	return nil
}

func (b *bookingService) GetByID(ctx context.Context, id int) (*model.Booking, error) {
	if id <= 0 {
		return nil, errors.New("Некорректный ID бронирования")
	}
	return b.bookingRepo.GetByID(ctx, id)
}

func (b *bookingService) GetByUserID(ctx context.Context, userID int) ([]*model.Booking, error) {
	if userID <= 0 {
		return nil, errors.New("Некорректный user_id")
	}
	return b.bookingRepo.GetByUserID(ctx, userID)
}
