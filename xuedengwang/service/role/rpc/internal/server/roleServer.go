// Code generated by goctl. DO NOT EDIT!
// Source: role.proto

package server

import (
	"context"

	"xuedengwang/service/role/rpc/internal/logic"
	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"
)

type RoleServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRoleServer
}

func NewRoleServer(svcCtx *svc.ServiceContext) *RoleServer {
	return &RoleServer{
		svcCtx: svcCtx,
	}
}

func (s *RoleServer) RoleAll(ctx context.Context, in *pb.RoleInterface) (*pb.RoleAllResp, error) {
	l := logic.NewRoleAllLogic(ctx, s.svcCtx)
	return l.RoleAll(in)
}

func (s *RoleServer) GetRole(ctx context.Context, in *pb.GetRoleReq) (*pb.GetRoleResp, error) {
	l := logic.NewGetRoleLogic(ctx, s.svcCtx)
	return l.GetRole(in)
}

func (s *RoleServer) GetRoleByID(ctx context.Context, in *pb.GetRoleByIDReq) (*pb.GetRoleByIDResp, error) {
	l := logic.NewGetRoleByIDLogic(ctx, s.svcCtx)
	return l.GetRoleByID(in)
}

func (s *RoleServer) UpdateRole(ctx context.Context, in *pb.UpdateRoleReq) (*pb.RoleInterface, error) {
	l := logic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *RoleServer) AddRole(ctx context.Context, in *pb.AddRoleReq) (*pb.RoleInterface, error) {
	l := logic.NewAddRoleLogic(ctx, s.svcCtx)
	return l.AddRole(in)
}

func (s *RoleServer) DeleteRoleByID(ctx context.Context, in *pb.DeleteRoleByIDReq) (*pb.RoleInterface, error) {
	l := logic.NewDeleteRoleByIDLogic(ctx, s.svcCtx)
	return l.DeleteRoleByID(in)
}