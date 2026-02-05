package repo

import (
	"context"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	query := `
	INSERT INTO users (email,name,password)
	VALUES ($1,$2,$3)
	RETURNING id, created_at
	`

	err := u.db.QueryRow(ctx, query, user.Email, user.Name, user.Password).Scan(&user.ID, &user.CreatedAt)

	return err
}

func (r *userRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `
	SELECT id, email, name, password, created_at FROM users
	WHERE id = $1
	`

	user := &model.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, pgx.ErrNoRows
	}

	return user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
    SELECT id, email, name, password, created_at FROM users
    WHERE email = $1
    `

	user := &model.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, pgx.ErrNoRows
	}

	return user, nil
}
