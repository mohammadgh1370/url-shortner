package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

type MockILinkRepo struct {
	*MockIBaseRepo
}

func NewMockILinkRepo(ctrl *gomock.Controller) *MockILinkRepo {
	mock := &MockILinkRepo{
		MockIBaseRepo: NewMockIBaseRepo(ctrl),
	}
	return mock
}
