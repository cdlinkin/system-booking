package repo

import (
	"context"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepo interface {
	Create(ctx context.Context, booking *model.Booking) error
	GetByID(ctx context.Context, id int) (*model.Booking, error)
	GetByUserID(ctx context.Context, userID int) ([]*model.Booking, error)
}

type bookingRepo struct {
	db *pgxpool.Pool
}

func NewBookingRepo(db *pgxpool.Pool) BookingRepo {
	return &bookingRepo{db: db}
}

func (b *bookingRepo) Create(ctx context.Context, booking *model.Booking) error {
	query := `
	INSERT INTO bookings (user_id, resource_id, status)
	VALUES ($1,$2,$3)
	RETURNING id, booked_at
	`

	err := b.db.QueryRow(ctx, query, booking.UserID, booking.ResourceID, booking.Status).Scan(&booking.ID, &booking.BookedAt)

	return err
}

func (b *bookingRepo) GetByID(ctx context.Context, id int) (*model.Booking, error) {
	query := `
	SELECT id, user_id, resource_id, status, booked_at FROM bookings
	WHERE id = $1
	`

	booking := &model.Booking{}
	err := b.db.QueryRow(ctx, query, id).Scan(&booking.ID, &booking.UserID, &booking.ResourceID, &booking.Status, &booking.BookedAt)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *bookingRepo) GetByUserID(ctx context.Context, userID int) ([]*model.Booking, error) {
	query := `
	SELECT id, user_id, resource_id, status, booked_at FROM bookings
	WHERE user_id = $1
	`

	rows, err := b.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*model.Booking

	for rows.Next() {
		booking := &model.Booking{}

		err := rows.Scan(&booking.ID, &booking.UserID, &booking.ResourceID, &booking.Status, &booking.BookedAt)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}
