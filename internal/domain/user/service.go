package user

import (
	"fmt"
	"gotickets/internal/auth"
	"gotickets/internal/domain/user/dto"
	"gotickets/internal/domain/user/types"
)

type service struct {
	repo Repository
	jwt auth.JWTService
}

func NewService(repo Repository,jwt auth.JWTService) *service {
	return &service{repo,jwt}
}

func (s *service) CreateUser(req dto.RegisterRequest)(*dto.RegisterResponse,error){
	user:=User{
		Name: req.Name,
		Email: req.Email,
		Role: types.Driver,
	}

	if err:=user.hashPassword((req.Password));err!=nil{
		return nil,err
	}
	if err:=s.repo.CreateUser(&user);err!=nil{
		return nil,err
	}

	response:=dto.RegisterResponse{
		Success: true,
		Message: "User created successfully",
		Data: dto.RegData{
			Id: user.ID,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
			CreatedAt: user.CreatedAt.String(),
			UptedAt: user.UpdatedAt.String(),
		},
	}
	return &response,nil
}
func (s *service) Login(req dto.LoginRequest)(*dto.LoginResponse,error){
	user,err:=s.repo.GetUserByEmail(req.Email)
	if err!=nil{
		return nil,err
	}

	if err:=user.checkPassword(req.Password,);err!=nil{
		return nil,err
	}

	token,err:=s.jwt.GenerateToken(user.ID,user.Name,user.Email)
    if err != nil {
		return nil,fmt.Errorf("failed to generate token: %w", err)
	}


	response:=dto.LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginData{
			Token: token,
			User: dto.UserData{
				Id: user.ID,
				Name: user.Name,
				Email: user.Email,
				Role: user.Role,
			},
		},
	}
	return &response,nil
}