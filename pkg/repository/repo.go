package repository

import (
	"errors"
	"server/pkg/common"
)

type RepositoryMethods interface {
	UserDataByID(int32) (common.UserData, error)
	UserDataListByID([]int32) ([]common.UserData, common.NotFoundList, error)
}

type Repository struct {
	UserData map[int32]common.UserData
}

func NewRepository() RepositoryMethods {
	repo := &Repository{
		UserData: make(map[int32]common.UserData),
	}
	sampleUsers := []common.UserData{
		{
			UserID:    1,
			FirstName: "John",
			City:      "New York",
			Phone:     "123-456-7890",
			Height:    175,
			Married:   false,
		},
		{
			UserID:    2,
			FirstName: "Alice",
			City:      "Los Angeles",
			Phone:     "987-654-3210",
			Height:    160,
			Married:   true,
		},
		{
			UserID:    3,
			FirstName: "Bob",
			City:      "Chicago",
			Phone:     "555-123-4567",
			Height:    180,
			Married:   true,
		},
		{
			UserID:    4,
			FirstName: "Emma",
			City:      "San Francisco",
			Phone:     "111-222-3333",
			Height:    165,
			Married:   false,
		},
		{
			UserID:    5,
			FirstName: "David",
			City:      "Seattle",
			Phone:     "777-888-9999",
			Height:    170,
			Married:   false,
		},
	}

	for _, user := range sampleUsers {
		repo.UserData[user.UserID] = user
	}
	return repo
}

func (repo *Repository) UserDataByID(id int32) (common.UserData, error) {
	user, ok := repo.UserData[id]
	if !ok {
		return common.UserData{}, errors.New("User not found")
	}
	return user, nil
}

func (repo *Repository) UserDataListByID(ids []int32) ([]common.UserData, common.NotFoundList, error) {
	var users []common.UserData
	var notfound common.NotFoundList
	for _, id := range ids {
		user, ok := repo.UserData[id]
		if !ok {
			notfound.UsersNotFound = append(notfound.UsersNotFound, id)
		} else {
			users = append(users, user)
		}
	}
	return users, notfound, nil
}
