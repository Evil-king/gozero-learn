package account

import (
	"context"
	"time"
	"user-api/biz"
	"user-api/internal/model"

	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userModel := model.NewUserModel(l.svcCtx.MysqlCon)
	//1、查询用户是否存在
	user, err := userModel.FindNameByUser(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("Register FindByUsername err: ", err)
		return nil, biz.DBError
	}
	if user != nil {
		//代表已经注册
		return nil, biz.AlreadyRegister
	}
	_, err = userModel.Insert(l.ctx, &model.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return
}
