package posts

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostReq) (resp *types.UpdatePostResp, err error) {
	// todo: add your logic here and delete this line
	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	//if err != nil && err != sqlx.ErrNotFound {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//
	//if user == nil {
	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	//}
	//
	//post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.PostID)
	//if err != nil && err != sqlx.ErrNotFound {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//if post == nil {
	//	return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
	//}
	//
	//if len(req.PostTitle) != 0 {
	//	post.PostTitle = req.PostTitle
	//}
	//
	//if len(req.PostDesc) != 0 {
	//	post.PostDesc = req.PostDesc
	//}
	//
	//err = l.svcCtx.PostModel.Update(l.ctx, post)
	//if err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	return
}
