package service

import (
	"context"
	"errors"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/cdlinkin/system-booking/internal/repo"
)

type BookingService interface {
	CreateBooking(ctx context.Context, input *BookingDTO) error
}

type bookingService struct {
	bookingRepo repo.BookingRepo
}

func NewBookingService(bookingRepo repo.BookingRepo) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

type BookingDTO struct {
	UserID     int
	ResourceID int
}

func (b *bookingService) CreateBooking(ctx context.Context, input *BookingDTO) error {
	if input.UserID <= 0 {
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
