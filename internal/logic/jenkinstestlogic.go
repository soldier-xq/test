package logic

import (
	"context"

	"jenkins_test/internal/svc"
	"jenkins_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Jenkins_testLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJenkins_testLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Jenkins_testLogic {
	return &Jenkins_testLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Jenkins_testLogic) Jenkins_test(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
