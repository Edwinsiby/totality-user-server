package server

import (
	"context"
	"server/pkg/common"
	"server/pkg/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockRepo = new(MockRepository)
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) UserDataByID(userID int32) (common.UserData, error) {
	return common.UserData{
		UserID:    userID,
		FirstName: "MockFirstName",
		City:      "MockCity",
		Phone:     "MockPhone",
		Height:    0,
		Married:   false,
	}, nil
}

func (m *MockRepository) UserDataListByID(ids []int32) ([]common.UserData, common.NotFoundList, error) {
	var users []common.UserData
	var notfound common.NotFoundList

	for _, id := range ids {
		if id == 999 {
			notfound.UsersNotFound = append(notfound.UsersNotFound, id)
		} else {

			user := common.UserData{
				UserID:    id,
				FirstName: "MockFirstName",
				City:      "MockCity",
				Phone:     "MockPhone",
				Height:    0,
				Married:   false,
			}
			users = append(users, user)
		}
	}

	return users, notfound, nil
}

func TestUserService_HealthCheck(t *testing.T) {
	userService := NewUserServices(mockRepo)

	ctx := context.Background()
	req := &pb.Request{}

	resp, err := userService.HealthCheck(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Server is running", resp.Result)
}

func TestUserService_UserDetails(t *testing.T) {

	userService := NewUserServices(mockRepo)

	ctx := context.Background()
	req := &pb.UserDetailsRequest{
		UserID: 1,
	}
	mockRepo.On("UserDataByID", req.UserID).Return(&pb.UserDetails{
		UserID:    1,
		FirstName: "John",
		City:      "New York",
		Phone:     "123-456-7890",
		Height:    175,
		Married:   false,
	}, nil)

	resp, err := userService.UserDetails(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUserService_UserListDetails(t *testing.T) {

	userService := NewUserServices(mockRepo)

	ctx := context.Background()
	req := &pb.UserListDetailsRequest{
		UserIDList: []int32{1, 2, 3},
	}

	mockRepo.On("UserDataListByID", req.UserIDList).Return([]pb.UserDetails{
		{
			UserID:    1,
			FirstName: "John",
			City:      "New York",
			Phone:     "123-456-7890",
			Height:    175,
			Married:   false,
		},
	},
		nil)

	resp, err := userService.UserListDetails(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
