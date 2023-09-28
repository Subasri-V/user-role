package controller

import (
	"context"
	"fmt"

	"github.com/Subasri-V/user-role.git/interfaces"
	"github.com/Subasri-V/user-role.git/models"

	cus "github.com/Subasri-V/user-role.git/proto"
)

type RPCServer struct {
	cus.UnimplementedUserRoleServiceServer
}

var (
	UserRoleDetails interfaces.IUserRole
)

func (s *RPCServer) AddUser(ctx context.Context, req *cus.UserRequest) (*cus.UserResponse, error) {
	newUser := &models.User{Name: req.Name, Email: req.Email, Password: req.Password, DOB: req.DOB, Role: req.Role, Status: req.Status}
	_, err := UserRoleDetails.AddUser(newUser)

	if err != nil {
		return nil, err
	} else {
		UserResponse := &cus.UserResponse{
			Message: "Success",
		}
		return UserResponse, nil
	}
}

func (s *RPCServer) EnableUser(ctx context.Context, req *cus.Name) (*cus.EnableResponse, error) {
	_, err := UserRoleDetails.EnableUser(req.Name)

	if err != nil {
		return nil, err
	} else {
		Enable := cus.EnableResponse{
			Message: "Enabled Successfully",
		}
		return &Enable, nil
	}
}

func (s *RPCServer) DisableUser(ctx context.Context, req *cus.Name) (*cus.DisableResponse, error) {
	fmt.Println("controller", req.Name)

	_, err := UserRoleDetails.DisableUser(req.Name)

	if err != nil {
		return nil, err
	} else {
		Disable := &cus.DisableResponse{
			Message: "disabled successfully",
		}
		return Disable, nil
	}
}

func (s *RPCServer) UpdateRole(ctx context.Context, req *cus.UpdateRoleRequest) (*cus.UpdateRoleResponse, error) {
	_, err := UserRoleDetails.UpdateRole(req.Name, req.Role)
	if err != nil {
		return nil, err
	} else {
		Update := cus.UpdateRoleResponse{
			Message: "Updated Role successfully",
		}
		return &Update, nil
	}
}

func (s *RPCServer) AssociateRole(ctx context.Context, req *cus.AssociateRoleRequest) (*cus.AssociateRoleResponse, error) {
	_, err := UserRoleDetails.AssociateRole(req.Name, req.Role)
	if err != nil {
		return nil, err
	} else {
		Update := cus.AssociateRoleResponse{
			Message: "Updated Associate Role successfully",
		}
		return &Update, nil
	}
}

func (s * RPCServer) Remove(ctx context.Context, req *cus.AssociateRoleRequest)(*cus.AssociateRoleResponse,error){
	_, err := UserRoleDetails.Remove(req.Name, req.Role)
	if err != nil {
		return nil, err
	} else {
		Update := cus.AssociateRoleResponse{
			Message: "Removed role successfully",
		}
		return &Update, nil
	}
}

func (s *RPCServer) AppendArray(ctx context.Context,req *cus.UpdateRoleRequest) (*cus.AssociateRoleResponse,error){
	_, err := UserRoleDetails.AppendArray(req.Name,req.Role)
	if err != nil {
		return nil, err
	} else {
		Update := cus.AssociateRoleResponse{
			Message: "appended array role successfully",
		}
		return &Update, nil
	}
}

func (s *RPCServer) ListFeatures(ctx context.Context, req *cus.Role)(*cus.ListFeaturesResponse,error){
	err:=UserRoleDetails.ListFeatures(req.Role)
	if err!=nil{
		return nil,err 
	}else{
		Update:=cus.ListFeaturesResponse{
			Message: "listed feature successfully",
		}
		return &Update,nil
	}
}

func (s *RPCServer) ListFeaturesInPostman(ctx context.Context,req *cus.Role) (*cus.Postresponse,error){
	fmt.Println("controller: 1")
	
	list:=&models.Role{
		Role: req.Role,
	}
	fmt.Println("controller: 2")

	result,err:=UserRoleDetails.ListFeaturesInPostman(list)
	fmt.Println("controller: 3")

	if err!=nil{
		return nil,err 
	}else{
		Update:=cus.Postresponse{
			Role:           result.Role,
			Access:         result.Access,
			Responsibility: result.Responsibility,
		}
		return &Update,nil
	}
}
