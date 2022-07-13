package mocks

//go:generate rm -rf ./*_mm.go
//go:generate minimock -i github.com/P2PCloud/miner-search-api/internal/app/test/mocks.* -o ./ -s "_mm.go"

import (
	"github.com/P2PCloud/miner-search-api/internal/app/repository"
)

type (
	// Repository ...
	Repository interface {
		repository.Repository
	}
)
