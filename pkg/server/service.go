package server

import (
	"context"
	"server/pkg/pb"
	"server/pkg/repository"
)

type UserServiceMethods interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
	UserDetails(context.Context, *pb.UserDetailsRequest) (*pb.UserDetailsResponse, error)
	UserListDetails(context.Context, *pb.UserListDetailsRequest) (*pb.UserListDetailsResponse, error)
}

type UserService struct {
	pb.UnimplementedUserServer
	Repo repository.RepositoryMethods
}

func NewUserServices(repo repository.RepositoryMethods) UserService {
	return UserService{
		Repo: repo,
	}
}

func (u UserService) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	return &pb.Response{Result: "Server is running"}, nil
}

func (u UserService) UserDetails(ctx context.Context, req *pb.UserDetailsRequest) (*pb.UserDetailsResponse, error) {

	response, err := u.Repo.UserDataByID(req.UserID)
	if err != nil {
		return nil, err
	}

	result := &pb.UserDetails{
		UserID:    response.UserID,
		FirstName: response.FirstName,
		City:      response.City,
		Phone:     response.Phone,
		Height:    response.Height,
		Married:   response.Married,
	}

	return &pb.UserDetailsResponse{Result: result}, nil
}

func (u UserService) UserListDetails(ctx context.Context, req *pb.UserListDetailsRequest) (*pb.UserListDetailsResponse, error) {

	response, notfound, err := u.Repo.UserDataListByID(req.UserIDList)
	if err != nil {
		return nil, err
	}
	userDetailsList := make([]*pb.UserDetails, 0, len(response))
	for _, user := range response {
		userDetailsList = append(userDetailsList, &pb.UserDetails{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			City:      user.City,
			Phone:     user.Phone,
			Height:    user.Height,
			Married:   user.Married,
		})
	}
	return &pb.UserListDetailsResponse{Result: userDetailsList, NotFound: notfound.UsersNotFound}, nil
}
