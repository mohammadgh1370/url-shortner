// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/base_repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

type MockIUserRepo struct {
	*MockIBaseRepo
}

func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{
		MockIBaseRepo: NewMockIBaseRepo(ctrl),
	}
	return mock
}
