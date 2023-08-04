package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

type MockIViewRepo struct {
	*MockIBaseRepo
}

func NewMockIViewRepo(ctrl *gomock.Controller) *MockIViewRepo {
	mock := &MockIViewRepo{
		MockIBaseRepo: NewMockIBaseRepo(ctrl),
	}
	return mock
}
