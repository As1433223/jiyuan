package logic

import (
	"context"
	"errors"
	"github.com/As1433223/jiyuan/models"
	"github.com/As1433223/jiyuan/proto"
)

func (s *ServerRpc) UserLogin(ctx context.Context, in *proto.UserLoginReq) (*proto.UserLoginRes, error) {
	var rr proto.UserLoginRes
	rr.Userid = 0
	user, _ := models.GetUserInfo(models.User{
		Username: in.Username,
	})
	if user.ID == 0 {
		return &rr, errors.New("账号不存在")
	}
	if user.Password != in.Password {
		return &rr, errors.New("密码错误")
	}
	rr.Userid = int64(user.ID)
	return &rr, nil
}
