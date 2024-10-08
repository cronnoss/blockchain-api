package servicegroup

import (
	"context"
	"fmt"
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
)

// GroupService is a group service.
type GroupService struct {
	repo Repository
}

// New creates a new group service.
func New(repo Repository) *GroupService {
	return &GroupService{
		repo: repo,
	}
}

// GetGroups returns all group IDs from repository.
func (svc *GroupService) GetGroupIDs(ctx context.Context) ([]int64, error) {
	groupIDs, err := svc.repo.GetGroupIDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting groups from repository: %w", err)
	}

	return groupIDs, nil
}

// GetGroup returns group from repository by ID.
func (svc *GroupService) GetGroup(ctx context.Context, id int64) (*models.Group, error) {
	group, err := svc.repo.GetGroup(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting group from repository: %w", err)
	}

	return group, nil
}
