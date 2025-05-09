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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userModel := model.NewUserModel(l.svcCtx.MysqlCon)
	//1、查询用户是否存在
	user, err := userModel.FindNameByUser(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("Register FindByUsername err: ", err)
		return nil, biz.DBError
	}
	if user == nil {
		//代表已经注册
		return nil, biz.NoRegister
	}
	//登录成功 生成token
	secret := l.svcCtx.Config.Auth.Secret
	expire := l.svcCtx.Config.Auth.Expire
	token, err := biz.GetJwtToken(secret, time.Now().Unix(), expire, user.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, biz.TokenError
	}
	resp = &types.LoginResp{
		Token: token,
	}
	return
}
