//go:generate mockgen -destination=./mocks.go -source=./services.go -package=handlerindex

package handlerindex

import (
	"context"

	"github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
)

// IndexService is an index service.
type IndexService interface {
	GetIndex(ctx context.Context, id int64) (*models.Index, error)
}
