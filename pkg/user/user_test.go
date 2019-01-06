package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/charlesakalugwu/codecov/pkg/mocks/mock_doer"
	"github.com/charlesakalugwu/codecov/pkg/user"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mock_doer.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}

func TestSay(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mock_doer.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Say to be called once with "Hello GoMock" as parameter, and return nil from the mocked call.
	mockDoer.EXPECT().SaySomething("Hello GoMock").Return(nil).Times(1)

	testUser.Say()
}
