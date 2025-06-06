// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package handler

import (
	"net/http"

	account "user-api/internal/handler/account"
	"user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: account.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: account.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/userInfo",
				Handler: account.GetUserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
		rest.WithJwt(serverCtx.Config.Auth.Secret),
	)
}
