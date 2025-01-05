package ctxData

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	CtxUserIdKey   = "userId"
	CtxUsernameKey = "username"
	CtxUserRoleKey = "userRole"
	CtxIpKey       = "ip"
)

func GetUseridFromCtx(ctx context.Context) int64 {
	var Uid int64
	if jsonUid, ok := ctx.Value(CtxUserIdKey).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			Uid = int64Uid

		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	} else {
		logx.WithContext(ctx).Errorf("GetUidFromCtx:err")
	}

	return Uid
}

func GetUserRoleFromCtx(ctx context.Context) int64 {

	var userRole int64
	if role, ok := ctx.Value(CtxUserRoleKey).(int64); ok {
		userRole = role
	}
	return userRole
}
func GetUsernameFormCtx(ctx context.Context) string {
	if jsonUserRole, ok := ctx.Value(CtxUsernameKey).(string); ok {
		return jsonUserRole
	} else {
		logx.WithContext(ctx).Errorf("GetUserRoleFromCtx")
	}
	return ""
}

func GetIpFromCtx(ctx context.Context) string {
	if jsonIp, ok := ctx.Value(CtxIpKey).(string); ok {
		return jsonIp
	} else {
		logx.WithContext(ctx).Errorf("GetIpFromCtx")
	}
	return ""
}
