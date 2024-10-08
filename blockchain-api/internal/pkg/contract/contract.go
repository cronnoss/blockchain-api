package contract

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/cronnoss/blockchain-api/blockchain-api/internal/abi"
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Contract is a smart contract.
type Contract struct {
	*abi.Contract
}

// Bind binds to already deployed contract.
func Bind(contractAddress string, backend bind.ContractBackend) (*Contract, error) {
	if contractAddress == "" {
		return nil, errors.New("no contract address provided")
	}
	if backend == nil {
		return nil, errors.New("no contract backend provided")
	}

	// Get the contract address from hex string
	contractAddr := common.HexToAddress(contractAddress)

	ctr, err := abi.NewContract(contractAddr, backend)
	if err != nil {
		return nil, fmt.Errorf("contract.NewContract failed: %w", err)
	}

	return &Contract{ctr}, nil
}

// GetGroupIDs returns all contract group IDs.
func (ec *Contract) GetGroupIDs(ctx context.Context) ([]int64, error) {
	groupIDs, err := ec.Contract.GetGroupIds(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("ctr.GetGroupIds failed: %w", err)
	}

	result := make([]int64, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		result = append(result, groupID.Int64())
	}

	return result, nil
}

// GetIndex fetches contract group by ID.
func (ec *Contract) GetGroup(ctx context.Context, id int64) (*models.Group, error) {
	group, err := ec.Contract.GetGroup(&bind.CallOpts{Context: ctx}, big.NewInt(id))
	if err != nil {
		return nil, fmt.Errorf("ctr.GetGroup failed: %w", err)
	}

	groupJSON, err := json.Marshal(group)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal failed: %w", err)
	}

	var result models.Group
	err = json.Unmarshal(groupJSON, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal failed: %w", err)
	}

	return &result, nil
}

// GetIndex fetches contract index by ID.
func (ec *Contract) GetIndex(ctx context.Context, id int64) (*models.Index, error) {
	index, err := ec.Contract.GetIndex(&bind.CallOpts{Context: ctx}, big.NewInt(id))
	if err != nil {
		return nil, fmt.Errorf("ctr.GetIndex failed: %w", err)
	}

	indexJSON, err := json.Marshal(index)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal failed: %w", err)
	}

	var result models.Index
	err = json.Unmarshal(indexJSON, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal failed: %w", err)
	}

	return &result, nil
}
