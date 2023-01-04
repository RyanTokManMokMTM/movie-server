package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"gorm.io/gorm"
	"time"
)

type Store interface {
	CountFriends(ctx context.Context, friendID uint) int64
	GetFriendsList(ctx context.Context, UserID uint) ([]*models.User, error)
	GetFriendRoomList(ctx context.Context, UserID uint) (*models.User, error)
	InsertOneFriendNotification(ctx context.Context, sender uint, receiver *models.User) (uint, error)
	FindOneFriendNotification(ctx context.Context, sender uint, receiver uint) (*models.FriendNotification, error)
	FindOneFriendNotificationByID(ctx context.Context, requestID uint) (*models.FriendNotification, error)
	AcceptFriendNotification(ctx context.Context, fr *models.FriendNotification) error
	CancelFriendNotification(ctx context.Context, requestID uint) error
	DeclineFriendNotification(ctx context.Context, requestID uint) error
	FindOneFriend(ctx context.Context, userID uint, friendID uint) (*models.User, error)
	RemoveFriend(ctx context.Context, userID uint, friendID uint) error
	GetFriendRequest(ctx context.Context, userID uint, limit int, pageOffset int) ([]*models.FriendNotification, int64, error)
	IsFriend(ctx context.Context, userID uint, friendID uint) (bool, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUserByID(ctx context.Context, userID uint) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, userID uint, user *models.User) error
	GetUserLikedMovies(ctx context.Context, userID uint, limit int, pageOffset int) (*models.User, error)
	CreateLikedMovie(ctx context.Context, movieID uint, userID uint) error
	CountLikedMovie(ctx context.Context, userID uint) int64
	CreatePostLiked(ctx context.Context, userId uint, postId *models.Post) error
	GetUserRooms(ctx context.Context, userID uint) ([]*models.Room, error)
	GetUserActiveRooms(ctx context.Context, userID uint) ([]models.Room, error)
	GetUserRoomsWithMembers(ctx context.Context, userID uint) (*models.User, error)
	InsertOneCommentLike(ctx context.Context, userID uint, commentID uint, count uint) error
	RemoveOneCommentLike(ctx context.Context, userID uint, commentID uint, count uint) error
	UpdateFriendNotification(ctx context.Context, u *models.User, count uint) error
	ResetFriendNotification(ctx context.Context, u *models.User) error
	UpdateLikesNotification(ctx context.Context, u *models.User, count uint) error
	ResetLikesNotification(ctx context.Context, u *models.User) error
	UpdateCommentNotification(ctx context.Context, u *models.User, count uint) error
	ResetCommentNotification(ctx context.Context, u *models.User) error
	FindOneMovieDetail(ctx context.Context, movieID uint) (*models.MovieInfo, error)
	FindOneMovie(ctx context.Context, movieID uint) (*models.MovieInfo, error)
	CreatePostComment(ctx context.Context, userID uint, PostID uint, comment string) (*models.Comment, error)
	CreatePostReplyComment(ctx context.Context, userID uint, PostID uint, replyCommentId uint, parentID uint, replyUserID uint, comment string) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) error
	DeleteComment(ctx context.Context, commentID uint) error
	FindPostComments(ctx context.Context, postID uint, likedBy uint, limit int, pageOffset int) ([]*models.Comment, int64, error)
	FindReplyComments(ctx context.Context, parentID uint, likedBy uint, limit int, pageOffset int) ([]*models.Comment, int64, error)
	FindOneComment(ctx context.Context, commentID uint) (*models.Comment, error)
	UpdateCommentCount(ctx context.Context, comment *models.Comment, updateCount uint) error
	FindOneUserLikedMovie(ctx context.Context, movieID uint, userID uint) (*models.User, error)
	CountLikesOfMovie(ctx context.Context, movieID uint) (int64, error)
	FindMovieListByGenreID(ctx context.Context, genreID uint) (*models.GenreInfo, error)
	DeletePostLikes(ctx context.Context, postLiked *models.PostLiked) error
	FindOnePostLiked(ctx context.Context, userId uint, postId uint) (*models.PostLiked, error)
	CountPostLikes(ctx context.Context, postId uint) (int64, error)
	CreateNewPost(ctx context.Context, post *models.Post) error
	UpdatePostInfo(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, postID uint, userID uint) error
	FindOnePostInfoWithUserLiked(ctx context.Context, postID uint, userID uint) (*models.Post, error)
	FindOnePostInfo(ctx context.Context, postID uint) (*models.Post, error)
	FindAllPosts(ctx context.Context, userID uint, limit int, offset int) ([]*models.Post, int64, error)
	FindFollowingPosts(ctx context.Context, userID uint, limit int, pageOffset int) ([]*models.Post, int64, error)
	FindUserPosts(ctx context.Context, userID uint, likedBy uint, limit int, pageOffset int) ([]*models.Post, int64, error)
	CountUserPosts(ctx context.Context, userID uint) (int64, error)
	CountCommentLikes(ctx context.Context, commentId uint) (int64, error)
	CountMovieCollected(ctx context.Context, movieId uint) (int64, error)
	FindOneMovieFormAnyList(ctx context.Context, movieID uint, userID uint) (*models.ListMovie, error)
	CreateNewList(ctx context.Context, ListTitle string, ListIntro string, userID uint) (*models.List, error)
	UpdateList(ctx context.Context, list *models.List) error
	DeleteList(ctx context.Context, listID uint, userID uint) error
	FindOneList(ctx context.Context, listID uint) (*models.List, error)
	FindOneUserList(ctx context.Context, listID uint, userID uint) (*models.List, error)
	FindUserLists(ctx context.Context, userID uint, limit int, pageOffset int) ([]*models.List, int64, error)
	CountListMovies(ctx context.Context, listID uint) (int64, error)
	CountCollectedMovie(ctx context.Context, userID uint) (int64, error)
	FindOneMovieFromList(ctx context.Context, movieID uint, listID uint, userID uint) (*models.MovieInfo, error)
	InsertMovieToList(ctx context.Context, movieID uint, listID uint, userID uint) error
	RemoveMovieFromList(ctx context.Context, movieID uint, listID uint, userID uint) error
	RemoveMoviesFromList(ctx context.Context, movieIds []uint, listID uint, userID uint) error
	FindListMovies(ctx context.Context, listID uint, lastCreateTime uint, limit int) ([]models.ListMovieInfoWithCreateTime, int64, error)
	UpdateUserGenres(ctx context.Context, ids []uint, userId uint) error
	FindUserGenres(ctx context.Context, userId uint) (*[]models.GenreInfo, error)
	InsertOneCommentNotification(ctx context.Context, receiverID uint, commentBy uint, postID uint, commentID uint, commentTime time.Time) error
	InsertOneReplyCommentNotification(ctx context.Context, receiverID uint, commentBy uint, postID uint, commentID uint, replyCommentID uint, commentTime time.Time) error
	FindOneCommentNotification(ctx context.Context, receiverID uint, limit int, pageOffset int) ([]*models.CommentNotification, int64, error)
	InsertOnePostLikeNotification(ctx context.Context, postID uint, likedBy uint, Receiver uint, likedTime time.Time) error
	InsertOneCommentLikeNotification(ctx context.Context, postID uint, likedBy uint, commentID uint, Receiver uint, likedTime time.Time) error
	FindLikesNotificationByReceiver(ctx context.Context, receiverID uint, limit int, pageOffset int) ([]*models.LikeNotification, int64, error)
	FindOneLikePostNotification(ctx context.Context, receiverID uint, likedBy uint, postID uint) error
	FindOneLikeCommentNotification(ctx context.Context, receiverID uint, likedBy uint, commentID uint) error
	InsertOneRoom(ctx context.Context, name string, info string, userID uint) (*models.Room, error)
	RemoveOneRoom(ctx context.Context, roomID uint) error
	FindOneOwnerRoom(ctx context.Context, roomID uint, userID uint) (*models.Room, error)
	FindOneByRoomID(ctx context.Context, roomID uint) (*models.Room, error)
	FindOneByRoomIDWithInfo(ctx context.Context, roomID uint) (*models.Room, error)
	JoinOneRoom(ctx context.Context, roomID uint, u *models.User) error
	LeaveOneRoom(ctx context.Context, roomID uint, u *models.User) error
	FindRoomMembers(ctx context.Context, roomID uint) ([]*models.User, error)
	FindOneRoomMember(ctx context.Context, roomID uint, userID uint) (*models.User, error)
	UpdateLastSender(ctx context.Context, roomID uint, sender uint) error
	UpdateIsRead(ctx context.Context, roomID uint, isRead bool) error
	CountMessage(ctx context.Context, roomID uint) (int64, error)
	ExistInTheRoom(ctx context.Context, userID uint, roomID uint) error
	GetRoomUsers(ctx context.Context, roomID uint) ([]uint, error)
	UpdateRoomActiveState(ctx context.Context, roomID uint, userID uint, state bool) error
	InsertOneMessage(ctx context.Context, roomID uint, userId uint, message string, messageID string, sentTime int64) error
	GetRoomMessage(ctx context.Context, roomID uint, lastID uint, limit int, pageOffset int) ([]*models.Message, int64, error)
	GetRoomLatestMessage(ctx context.Context, roomID uint) ([]models.Message, error)
	RemoveUserLikedMovie(ctx context.Context, movieID, userID uint) error
}

var _ Store = (*DAO)(nil)

type DAO struct {
	engine *gorm.DB
}

func NewDAO(engine *gorm.DB) Store {
	return &DAO{engine: engine}
}
