// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import (
	builder "github.com/monitoror/monitoror/pkg/monitoror/builder"

	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/models"

	utils "github.com/monitoror/monitoror/pkg/monitoror/utils"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// DisableTile provides a mock function with given fields: tileType, variant
func (_m *Manager) DisableTile(tileType models.TileType, variant string) {
	_m.Called(tileType, variant)
}

// RegisterMetaTile provides a mock function with given fields: tileType, variant, clientConfigValidator, _a3
func (_m *Manager) RegisterMetaTile(tileType models.TileType, variant string, clientConfigValidator utils.Validator, _a3 builder.MetaTileBuilder) {
	_m.Called(tileType, variant, clientConfigValidator, _a3)
}

// RegisterTile provides a mock function with given fields: tileType, variant, clientConfigValidator, path, initialMaxDelay
func (_m *Manager) RegisterTile(tileType models.TileType, variant string, clientConfigValidator utils.Validator, path string, initialMaxDelay int) {
	_m.Called(tileType, variant, clientConfigValidator, path, initialMaxDelay)
}