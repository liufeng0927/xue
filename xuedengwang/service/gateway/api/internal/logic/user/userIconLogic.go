package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIconLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var filePath = "E:\\xue\\student-mangement-system-web-master\\src\\assets\\icon\\"

func NewUserIconLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIconLogic {
	return &UserIconLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserIconLogic) UserIcon(file multipart.File, header *multipart.FileHeader) (*types.UserIconResp, error) {
	// todo: add your logic here and delete this line

	f, err := os.OpenFile(filePath+header.Filename, os.O_WRONLY|os.O_CREATE, 0666) //根据文件名创建文件路径
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrapf(err, "open file:%v error", header.Filename)
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return nil, errors.Wrapf(err, "write file:%v error", header.Filename)
	} //写入文件

	return &types.UserIconResp{
		UserIcon: header.Filename,
	}, nil
}
