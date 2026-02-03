package repo

import (
	"context"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ResourceRepo interface {
	GetAll(ctx context.Context) ([]*model.Resource, error)
	GetByID(ctx context.Context, id int) (*model.Resource, error)
	GetAvailable(ctx context.Context, isAvailable bool) ([]*model.Resource, error)
}

type resourceRepo struct {
	db *pgxpool.Pool
}

func NewResourceRepo(db *pgxpool.Pool) ResourceRepo {
	return &resourceRepo{db: db}
}

func (r *resourceRepo) GetAll(ctx context.Context) ([]*model.Resource, error) {
	query := `
	SELECT id, name, type, is_available, created_at FROM resources
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []*model.Resource

	for rows.Next() {
		res := &model.Resource{}

		err := rows.Scan(&res.ID, &res.Name, &res.Type, &res.IsAvailable, &res.CreatedAt)
		if err != nil {
			return nil, err
		}

		resources = append(resources, res)
	}

	return resources, nil
}

func (r *resourceRepo) GetByID(ctx context.Context, id int) (*model.Resource, error) {
	query := `
    SELECT id, name, type, is_available, created_at FROM resources
    WHERE id = $1
    `

	res := &model.Resource{}

	err := r.db.QueryRow(ctx, query, id).Scan(&res.ID, &res.Name, &res.Type, &res.IsAvailable, &res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *resourceRepo) GetAvailable(ctx context.Context, isAvailable bool) ([]*model.Resource, error) {
	query := `
	SELECT id, name, type, is_available, created_at FROM resources
	WHERE is_available = $1
	`

	rows, err := r.db.Query(ctx, query, isAvailable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []*model.Resource

	for rows.Next() {
		res := &model.Resource{}

		err := rows.Scan(&res.ID, &res.Name, &res.Type, &res.IsAvailable, &res.CreatedAt)
		if err != nil {
			return nil, err
		}

		resources = append(resources, res)
	}

	return resources, nil
}
