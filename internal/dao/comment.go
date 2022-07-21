package dao

import (
	"context"
)

func (d *DAO) InsertCommentToPost(ctx context.Context, postID, userID uint, comment string) error {
	return nil
}

func (d *DAO) RemoveCommentFormPost(ctx context.Context, postID, userID uint, comment string) error {
	return nil
}
