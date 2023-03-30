package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"xuedengwang/core/errx"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
	// todo: add your logic here and delete this line

	userDB := l.svcCtx.Query.User
	where := userDB.WithContext(l.ctx).Preload(userDB.Role)

	if len(in.Status) > 0 {
		where.Where(userDB.Status.Eq(cast.ToInt32(in.Status)))
	}

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where.Where(userDB.Realname.Like(searchValue)).Or(userDB.Username.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	users, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find user req: %+v", in)

	}
	var content []*pb.UserDao
	for _, user := range users {
		role := &pb.Role{
			CreateTime: user.Role.CreateTime.String(),
			CreateBy:   user.Role.CreateBy,
			UpdateTime: user.Role.CreateTime.String(),
			UpdateBy:   user.Role.UpdateBy,
			Remarks:    user.Role.Remarks,
			ID:         user.Role.ID,
			Name:       user.Role.Name,
			Code:       user.Role.Code,
		}
		userDao := &pb.UserDao{
			ID:         user.ID,
			CreateTime: user.CreateTime.String(),
			UpdateTime: user.UpdateTime.String(),
			SysRole:    role,
		}
		_ = copier.Copy(userDao, user)

		content = append(content, userDao)
	}

	return &pb.GetUserResp{
		Content:       content,
		TotalElements: count,
	}, nil
}
