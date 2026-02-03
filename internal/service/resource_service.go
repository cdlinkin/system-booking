package service

import (
	"context"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/cdlinkin/system-booking/internal/repo"
)

type ResourceService interface {
	GetResources(ctx context.Context, onlyAvailable bool) ([]*model.Resource, error)
}

type resourceService struct {
	resourceRepo repo.ResourceRepo
}

func NewResourceService(resourceRepo repo.ResourceRepo) ResourceService {
	return &resourceService{resourceRepo: resourceRepo}
}

func (r *resourceService) GetResources(ctx context.Context, onlyAvailable bool) ([]*model.Resource, error) {
	if onlyAvailable {
		res, err := r.resourceRepo.GetAvailable(ctx, true)
		if err != nil {
			return nil, err
		}

		return res, nil
	} else {
		res, err := r.resourceRepo.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
