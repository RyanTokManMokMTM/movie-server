package post_likes

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/serverWs"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type CreatePostLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLikesLogic {
	return &CreatePostLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLikesLogic) CreatePostLikes(req *types.CreatePostLikesReq) (resp *types.CreatePostLikesResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	u, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//find post
	post, err := l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOnePostLiked(l.ctx, userID, req.PostId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Create a new record
			if err := l.svcCtx.DAO.CreatePostLiked(l.ctx, userID, post); err != nil {
				return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
			}
			//TODO:add a notification record
			//TODO:liked by itself - ignored
			if userID != post.UserId {
				//TODO: is notification exist?
				err = l.svcCtx.DAO.FindOneLikePostNotification(l.ctx, post.UserId, userID, post.PostId)
				if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
					logx.Info("Notification not found...")

					//TODO: add a new one
					if err := l.svcCtx.DAO.InsertOnePostLikeNotification(l.ctx, req.PostId, userID, post.UserId, time.Now()); err != nil {
						return nil, err
					}
					//TODO: send the notification
					go func() {
						logx.Info("TODO: Send a liked post notification")
						_ = serverWs.SendNotificationToUserWithUserInfo(post.UserId, u, fmt.Sprintf("%s給您的文章點讚", u.Name), serverWs.LIKES_NOTIFICATION)
					}()
				}

			}

			return &types.CreatePostLikesResp{}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//if postLiked.State == 1 {
	//	postLiked.State = 1
	//} else {
	//postLiked.State = 1 //always be true
	//}

	//if err := l.svcCtx.DAO.UpdatePostLiked(l.ctx, postLiked); err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}

	return &types.CreatePostLikesResp{}, nil //already liked - do nothing?
}
