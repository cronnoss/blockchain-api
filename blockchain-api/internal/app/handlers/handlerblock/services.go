//go:generate mockgen -destination=./mocks.go -source=./services.go -package=handlerblock

package handlerblock

import (
	"context"
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// BlockService is a block service.
type BlockService interface {
	GetBlockByNumber(ctx context.Context, id *big.Int) (*models.Block, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error)

	GetBlockHeaderByNumber(ctx context.Context, id *big.Int) (*models.Block, error)
	GetBlockHeaderByHash(ctx context.Context, hash common.Hash) (*models.Block, error)
}
