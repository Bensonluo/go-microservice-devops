package handler

import (
	"../domain/model"
	"../domain/service"
	user "../proto/user"
	"context"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	userRegisterInfo := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegisterInfo)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "Added User successfully"
	return nil
}

func (u *User) Login(ctx context.Context, userLoginRequest *user.UserLoginRequest, userLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	userLoginResponse.Message = "Logged in successfully"
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserInfoTrans(userInfo)
	return nil
}

func UserInfoTrans(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
