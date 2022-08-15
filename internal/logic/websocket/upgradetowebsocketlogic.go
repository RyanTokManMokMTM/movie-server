package websocket

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpgradeToWebSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpgradeToWebSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpgradeToWebSocketLogic {
	return &UpgradeToWebSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpgradeToWebSocketLogic) UpgradeToWebSocket(req *types.UpgradeToWebSocketReq) (resp *types.UpgradeToWebSocketResp, err error) {
	// todo: add your logic here and delete this line

	return
}
