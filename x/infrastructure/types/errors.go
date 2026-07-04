package types

import "cosmossdk.io/errors"

const (
	ErrInvalidInfrastructureID uint32 = 1000
	ErrInfrastructureExists    uint32 = 1001
	ErrInfrastructureNotFound  uint32 = 1002
)

var (
	ErrInvalidID = errors.Register(
		ModuleName,
		ErrInvalidInfrastructureID,
		"invalid infrastructure id",
	)

	ErrAlreadyExists = errors.Register(
		ModuleName,
		ErrInfrastructureExists,
		"infrastructure already exists",
	)

	ErrNotFound = errors.Register(
		ModuleName,
		ErrInfrastructureNotFound,
		"infrastructure not found",
	)
)
