// Code generated by goctl. DO NOT EDIT.
package types

type HealthCheckReq struct {
}

type HealthCheckResp struct {
	Result string `json:"result"`
}

type MetaData struct {
	TotalPages   uint `json:"total_pages"`
	TotalResults uint `json:"total_results"`
	Page         uint `json:"page"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,max=32,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserLoginResp struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}

type UserSignUpReq struct {
	UserName string `json:"name"`
	Email    string `json:"email" validate:"email,max=32"`
	Password string `json:"password" validate:"min=8,max=32"`
}

type UserSignUpResp struct {
	Token       string `json:"token"`
	ExpiredTime uint   `json:"expired_time"`
}

type UserInfoReq struct {
	ID uint `path:"id"`
}

type UserInfoResp struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Cover  string `json:"cover"`
}

type UserProfileReq struct {
}

type UserProfileResp struct {
	ID               uint                  `json:"id"`
	Name             string                `json:"name"`
	Email            string                `json:"email"`
	Avatar           string                `json:"avatar"`
	Cover            string                `json:"cover"`
	NotificationInfo NotificationCountInfo `json:"notification_info"` //only logged in user will have this information
}

type UpdateProfileReq struct {
	Name string `json:"name"`
}

type UpdateProfileResp struct {
}

type UploadImageReq struct {
}

type UploadImageResp struct {
	Path string `json:"path"`
}

type CountFriendReq struct {
	UserId uint `path:"user_id"`
}

type CountFriendResp struct {
	Total uint `json:"total"`
}

type GetFriendListReq struct {
	UserId uint `path:"user_id"`
}

type GetFriendListResp struct {
	Friends []UserInfo `json:"Friends"`
}

type GetUserFriendRoomListReq struct {
}

type GetUserFriendRoomListResp struct {
	Lists []FriendRoomInfo `json:"lists"`
}

type FriendNotificationReq struct {
}

type FriendNotificationResp struct {
}

type CommentNotificationReq struct {
}

type CommentNotificationResp struct {
}

type LikesNotificationReq struct {
}

type LikesNotificationResp struct {
}

type FriendRoomInfo struct {
	RoomID uint     `json:"id"`
	Info   UserInfo `json:"info"`
}

type UserInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type NotificationCountInfo struct {
	FriendNotificationCount  uint `json:"friend_notification_count"`
	LikesNotificationCount   uint `json:"likes_notification_count"`
	CommentNotificationCount uint `json:"comment_notification_count"`
}

type MoviePageListByGenreReq struct {
	Id uint `path:"genre_id" validate:"numeric"`
}

type MoviePageListByGenreResp struct {
	Resp []*MovieInfo `json:"movie_infos"`
}

type MovieGenresInfoReq struct {
	Id uint `path:"movie_id" validate:"numeric"`
}

type MovieGenresInfoResp struct {
	Resp []*GenreInfo `json:"genres"`
}

type MovieDetailReq struct {
	MovieID uint `path:"movie_id"`
}

type MovieDetailResp struct {
	Info MovieDetailInfo `json:"info"`
}

type CountMovieLikesReq struct {
	MovieID uint `path:"movie_id"`
}

type CountMovieLikedResp struct {
	Count uint `json:"total_liked"`
}

type CountMovieCollectedReq struct {
	MovieID uint `path:"movie_id"`
}

type CountMovieCollectedResp struct {
	Count uint `json:"total_collected"`
}

type MovieInfo struct {
	MovieID     uint    `json:"id"`
	PosterPath  string  `json:"poster"`
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
}

type MovieDetailInfo struct {
	Adult            bool        `json:"adult"`
	BackdropPath     string      `json:"backdrop_path"`
	MovieId          uint        `json:"movie_id"`
	OriginalLanguage string      `json:"original_language"`
	OriginalTitle    string      `json:"original_title"`
	Overview         string      `json:"overview"`
	Popularity       float64     `json:"popularity"`
	PosterPath       string      `json:"poster_path"`
	ReleaseDate      string      `json:"release_date"`
	Title            string      `json:"title"`
	RunTime          int64       `json:"run_time"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"Vote_average"`
	VoteCount        int64       `json:"vote_count"`
	Genres           []GenreInfo `json:"genres"`
}

type GenreInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type LikedMovieReq struct {
	MovieID uint `json:"movie_id"`
}

type LikedMovieResp struct {
}

type AllUserLikedMoviesReq struct {
	ID    uint `path:"user_id"`
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type AllUserAllLikedMoviesResp struct {
	LikedMoviesList []*LikedMovieInfo `json:"liked_movies"`
	MetaData        MetaData          `json:"meta_data"`
}

type IsLikedMovieReq struct {
	MovieID uint `path:"movie_id"`
}

type IsLikedMovieResp struct {
	IsLiked bool `json:"is_liked_movie"`
}

type RemoveLikedMovieReq struct {
	MovieID uint `json:"movie_id"`
}

type RemoveLikedMovieResp struct {
}

type LikedMovieInfo struct {
	MovieID      uint        `json:"id"`
	MovieName    string      `json:"movie_name"`
	Genres       []GenreInfo `json:"genres"`
	MoviePoster  string      `json:"movie_poster"`
	MovieVoteAvg float64     `json:"vote_average"`
}

type CreateCustomListReq struct {
	Title string `json:"title"`
	Intro string `json:"intro"`
}

type CreateCustomListResp struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Intro string `json:"intro"`
}

type UpdateCustomListReq struct {
	ID    uint   `json:"list_id"`
	Title string `json:"title"`
	Intro string `json:"intro"`
}

type UpdateCustomListResp struct {
}

type DeleteCustomListReq struct {
	ID uint `json:"list_id"`
}

type DeleteCustomListResp struct {
}

type AllCustomListReq struct {
	ID    uint `path:"user_id"`
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type AllCustomListResp struct {
	Lists    []ListInfo `json:"lists"`
	MetaData MetaData   `json:"meta_data"`
}

type UserListReq struct {
	ID uint `path:"list_id"`
}

type UserListResp struct {
	List ListInfo `json:"list"`
}

type InsertMovieReq struct {
	ListID  uint `path:"list_id"`
	MovieID uint `path:"movie_id"`
}

type InsertMovieResp struct {
}

type RemoveMovieReq struct {
	ListID  uint `path:"list_id"`
	MovieID uint `path:"movie_id"`
}

type RemoveMovieResp struct {
}

type GetOneMovieFromUserListReq struct {
	MovieID uint `path:"movie_id"`
}

type GetOneMovieFromUserListResp struct {
	ListId        uint `json:"list_id"`
	IsMovieInList bool `json:"is_movie_in_list"`
}

type RemoveListMoviesReq struct {
	ListId   uint   `path:"id"`
	MovieIds []uint `json:"movie_ids"`
}

type RemoveListMoviesResp struct {
}

type ListInfo struct {
	ID     uint        `json:"id"`
	Title  string      `json:"title"`
	Intro  string      `json:"intro"`
	Movies []MovieInfo `json:"movie_list"`
}

type CreatePostReq struct {
	PostTitle string `json:"post_title"`
	PostDesc  string `json:"post_desc"`
	MovieID   uint   `json:"movie_id"`
}

type CreatePostResp struct {
	PostID     uint  `json:"id"`
	CreateTime int64 `json:"create_time"`
}

type UpdatePostReq struct {
	PostID    uint   `json:"post_id"`
	PostTitle string `json:"post_title"`
	PostDesc  string `json:"post_desc"`
}

type UpdatePostResp struct {
}

type DeletePostReq struct {
	PostID uint `json:"post_id"`
}

type DeletePostResp struct {
}

type AllPostsInfoReq struct {
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type AllPostsInfoResp struct {
	Infos    []PostInfo `json:"post_info"`
	MetaData MetaData   `json:"meta_data"`
}

type FollowPostsInfoReq struct {
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type FollowPostsInfoResp struct {
	Infos    []PostInfo `json:"post_info"`
	MetaData MetaData   `json:"meta_data"`
}

type PostInfoByIdReq struct {
	PostID uint `path:"post_id"`
}

type PostInfoByIdResp struct {
	Info PostInfo `json:"post_info"`
}

type PostsInfoReq struct {
	UserID uint `path:"user_id"`
	Page   uint `form:"page,default=1"`
	Limit  uint `form:"limit,default=20"`
}

type PostsInfoResp struct {
	Infos    []PostInfo `json:"post_info"`
	MetaData MetaData   `json:"metadata"`
}

type CountPostLikedReq struct {
	PostID uint `json:"post_id"`
}

type CountPostLikedResp struct {
	Count uint `json:"total_liked"`
}

type CountUserPostsReq struct {
	UserId uint `path:"user_id"`
}

type CountUserPostsResp struct {
	Count uint `json:"total_posts"`
}

type PostInfo struct {
	PostID            uint          `json:"id"`
	PostUser          PostUserInfo  `json:"user_info"`
	PostTitle         string        `json:"post_title"`
	PostDesc          string        `json:"post_desc"`
	PostMovie         PostMovieInfo `json:"post_movie_info"`
	PostLikeCount     int64         `json:"post_like_count"`
	PostCommentCount  int64         `json:"post_comment_count"`
	IsPostLikedByUser bool          `json:"is_post_liked"`
	CreateAt          int64         `json:"create_at"`
}

type PostMovieInfo struct {
	MovieID    uint   `json:"id"`
	PosterPath string `json:"poster_path"`
	Title      string `json:"title"`
}

type PostUserInfo struct {
	UserID     uint   `json:"id"`
	UserName   string `json:"name"`
	UserAvatar string `json:"avatar"`
}

type CreateCommentReq struct {
	PostID  uint   `path:"post_id"`
	Comment string `json:"comment"`
}

type CreateCommentResp struct {
	CommentID uint  `json:"id"`
	CreateAt  int64 `json:"create_at"`
}

type CreateReplyCommentReq struct {
	PostID          uint   `path:"post_id"`
	ReplyCommentId  uint   `path:"comment_id"`
	ParentCommentID uint   `json:"parent_id"`
	Comment         string `json:"comment"`
}

type CreateReplyCommentResp struct {
	CommentID uint  `json:"id"`
	CreateAt  int64 `json:"create_at"`
}

type UpdateCommentReq struct {
	CommentID uint   `path:"comment_id"`
	Comment   string `json:"comment"`
}

type UpdateCommentResp struct {
	UpdateAt int64 `json:"update_at"`
}

type DeleteCommentReq struct {
	CommentID uint `path:"comment_id"`
}

type DeleteCommentResp struct {
}

type GetPostCommentsReq struct {
	PostID uint `path:"post_id"`
	Page   uint `form:"page,default=1"`
	Limit  uint `form:"limit,default=20"`
}

type GetPostCommentsResp struct {
	Comments []CommentInfo `json:"comments"`
	MetaData MetaData      `json:"meta_data"`
}

type GetReplyCommentReq struct {
	CommentId       uint `path:"comment_id"`
	Page            uint `form:"page,default=1"`
	Limit           uint `form:"limit,default=5"`
	ParentCommentID uint `path:"comment_id"`
}

type GetReplyCommentResp struct {
	ReplyComments []CommentInfo `json:"reply"`
	MetaData      MetaData      `json:"meta_data"`
}

type CountPostCommentsReq struct {
	PostId uint `path:"post_id"`
}

type CountPostCommentsResp struct {
	TotalComment uint `json:"total_comment"`
}

type CommentInfo struct {
	CommentID       uint        `json:"id"`
	UserInfo        CommentUser `json:"user_info"`
	Comment         string      `json:"comment"`
	UpdateAt        int64       `json:"update_at"`
	ReplyID         uint        `json:"reply_id"`
	ReplyTo         UserInfo    `json:"reply_to"`
	ReplyComment    uint        `json:"reply_comments"`
	LikesCount      uint        `json:"comment_likes_count"`
	ParentCommentID uint        `json:"parent_comment_id"`
	IsLiked         bool        `json:"is_liked"`
}

type CommentUser struct {
	UserID     uint   `json:"id"`
	UserName   string `json:"name"`
	UserAvatar string `json:"avatar"`
}

type AddFriendReq struct {
	UserID uint `json:"user_id"`
}

type AddFriendResp struct {
	SenderID  uint `json:"sender"`
	RequestID uint `json:"request_id"`
}

type RemoveFriendReq struct {
	FriendID uint `json:"user_id"`
}

type RemoveFriendResp struct {
}

type AcceptFriendNotificationReq struct {
	RequestID uint `json:"request_id"`
}

type AcceptFriendNotificationResp struct {
	Message string `json:"message"`
}

type DeclineFriendNotificationReq struct {
	RequestID uint `json:"request_id"`
}

type DeclineFriendNotificationResp struct {
	Message string `json:"message"`
}

type CancelFriendNotificationReq struct {
	RequestID uint `json:"request_id"`
}

type CancelFriendNotificationResp struct {
	Message string `json:"message"`
}

type GetFriendRequestReq struct {
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type GetFriendRequestResp struct {
	Requests []FriendRequest `json:"requests"`
	MetaData MetaData        `json:"meta_data"`
}

type IsFriendReq struct {
	UserID uint `path:"user_id"`
}

type IsFriendResp struct {
	IsFriend      bool             `json:"is_friend"`
	IsSentRequest bool             `json:"is_sent_request"`
	RequestInfo   BasicRequestInfo `json:"request"`
}

type BasicRequestInfo struct {
	RequestID uint `json:"request_id"`
	SenderID  uint `json:"sender_id"`
}

type FriendRequest struct {
	RequestID uint     `json:"request_id"`
	Sender    UserInfo `json:"sender"`
	SentTime  int64    `json:"send_time"`
	State     uint     `json:"state"`
}

type CreateCommentLikesReq struct {
	CommentId uint `json:"comment_id"`
}

type CreateCommentLikesResp struct {
}

type RemoveCommentLikesReq struct {
	CommentId uint `json:"comment_id"`
}

type RemoveCommentLikesResq struct {
}

type CreatePostLikesReq struct {
	PostId uint `json:"post_id"`
}

type CreatePostLikesResp struct {
}

type RemovePostLikesReq struct {
	PostId uint `json:"post_id"`
}

type RemovePostLikesResq struct {
}

type IsPostLikedReq struct {
	PostId uint `path:"post_id"`
}

type IsPostLikedResp struct {
	IsLiked bool `json:"is_liked"`
}

type CountPostLikesReq struct {
	PostId uint `path:"post_id"`
}

type CountPostLikesResp struct {
	TotalLikes uint `json:"total_likes"`
}

type UpdateUserGenreReq struct {
	GenreIds []uint `json:"genre_ids"`
}

type UpdateUserGenreResp struct {
}

type GetUserGenreReq struct {
	UserId uint `path:"user_id"`
}

type GetUserGenreResp struct {
	UserGenres []GenreInfo `json:"user_genres"`
}

type CreateRoomReq struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

type CreateRoomResp struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"room_name"`
	Info   string `json:"room_info"`
}

type DeleteRoomReq struct {
	ID uint `json:"room_id"`
}

type DeleteRoomResp struct {
}

type JoinRoomReq struct {
	RoomID uint `path:"room_id"`
}

type JoinRoomResp struct {
}

type LeaveRoomReq struct {
	RoomID uint `path:"room_id"`
}

type LeaveRoomResp struct {
}

type GetRoomMembersReq struct {
	RoomID uint `path:"room_id"`
}

type GetRoomMembersResp struct {
	Members []UserInfo `json:"members"`
}

type GetUserRoomsReq struct {
}

type GetUserRoomsResp struct {
	Rooms []ChatRoomData `json:"rooms"`
}

type GetRoomInfoReq struct {
	RoomID uint `path:"room_id"`
}

type GetRoomInfoResp struct {
	Info ChatRoomData `json:"info"`
}

type SetIsReadReq struct {
	RoomID uint `path:"room_id"`
}

type SetIsReadResp struct {
}

type ChatRoomData struct {
	ID           uint          `json:"id"`
	Users        []UserInfo    `json:"users"`
	Messages     []MessageInfo `json:"messages"`
	LastSenderID uint          `json:"last_sender_id"`
	IsRead       bool          `json:"is_read"`
}

type MessageInfo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Sender   uint   `json:"sender_id"`
	SentTime int64  `json:"sent_time"`
}

type GetRoomMessageReq struct {
	RoomID uint `path:"room_id"`
	Page   uint `form:"page,default=1"`
	Limit  uint `form:"limit,default=20"`
}

type GetRoomMessageResp struct {
	Messagees []MessageData `json:"messages"`
	MetaData  MetaData      `json:"meta_data"`
}

type MessageData struct {
	MessageID string   `json:"id"`
	UserInfo  UserInfo `json:"users"`
	Content   string   `json:"content"`
	SendTime  int64    `json:"send_time"`
}

type GetLikeNotificationReq struct {
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type GetLikeNotificationResp struct {
	LikedNotificationList []LikedNotification `json:"notifications"`
	MetaData              MetaData            `json:"meta_data"`
}

type LikedNotification struct {
	ID          uint              `json:"id"`
	LikedBy     UserInfo          `json:"liked_by"`
	PostInfo    SimplePostInfo    `json:"post_info"`
	CommentInfo SimpleCommentInfo `json:"comment_info"`
	Type        uint              `json:"type"`
	LikedAt     uint              `json:"liked_at"`
}

type SimplePostInfo struct {
	PostID    uint          `json:"id"`
	PostMovie PostMovieInfo `json:"post_movie_info"`
}

type SimpleCommentInfo struct {
	CommentID uint   `json:"id"`
	Comment   string `json:"comment"`
}

type GetCommentNotificationReq struct {
	Page  uint `form:"page,default=1"`
	Limit uint `form:"limit,default=20"`
}

type GetCommentNotificationResp struct {
	CommentNotificationList []CommentNotification `json:"notifications"`
	MetaData                MetaData              `json:"meta_data"`
}

type CommentNotification struct {
	ID               uint              `json:"id"`
	CommentBy        UserInfo          `json:"comment_by"`
	PostInfo         SimplePostInfo    `json:"post_info"`
	CommentInfo      SimpleCommentInfo `json:"comment_info"`
	ReplyCommentInfo SimpleCommentInfo `json:"reply_comment_info"` //only reply type will have this info
	Type             uint              `json:"type"`
	CommentAt        uint              `json:"comment_at"`
}
