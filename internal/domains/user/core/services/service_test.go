package service

import (
	"testing"

	custSvcMocks "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/mocks"
	userCtMocks "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/mocks"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/ports"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

/*
	cmd

	go test github.com/kaolnwza/proj-blueprint/internal/domains/user/services -v -cover

	=== Run Specific Test
	go test ./internal/domains/user/core/services -v -testify.m TestName

*/

type TestSuite struct {
	suite.Suite

	mUserRepo    *mocks.MockRepository
	mCustSvcRepo *custSvcMocks.MockRepository
	mUserCtRepo  *userCtMocks.MockRepository

	s ports.Service
}

func (t *TestSuite) SetupTest() {
	ctrl := gomock.NewController(t.T())

	t.mUserRepo = mocks.NewMockRepository(ctrl)
	t.mCustSvcRepo = custSvcMocks.NewMockRepository(ctrl)
	t.mUserCtRepo = userCtMocks.NewMockRepository(ctrl)

	t.s = New(t.mUserRepo, t.mCustSvcRepo, t.mUserCtRepo)
}

// func (t *TestSuite) TearDownTest() {
// 	t.mCTRL.Finish()
// }

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
