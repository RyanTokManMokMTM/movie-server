// Code generated by goctl. DO NOT EDIT.
package types

type HealthCheckReq struct {
}

type HealthCheckResp struct {
	Result string `json:"result"`
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
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInfoReq struct {
	ID uint `path:"id"`
}

type UserInfoResp struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `jsons:"email"`
}

type UserProfileReq struct {
}

type UserProfileResp struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Cover  string `json:"cover"`
}

type MoviePageListByGenreRequest struct {
	Id uint `path:"genre_id" validate:"numeric"`
}

type MoviePageListByGenreResponse struct {
	Resp []*MovieInfo `json:"movieInfos"`
}

type MovieGenresInfoRequest struct {
	Id uint `path:"movie_id" validate:"numeric"`
}

type MovieGenresInfoResponse struct {
	Resp []*GenreInfo `json:"genres"`
}

type MovieDetailReq struct {
	MovieID uint `path:"movie_id"`
}

type MovieDetailResp struct {
	MovieDetailInfo
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

type CreateLikedMovieReq struct {
	MovieID uint `json:"movie_id"`
}

type CreateLikedMovieResp struct {
}

type DeleteLikedMoviedReq struct {
	MovieID uint `json:"movie_id"`
}

type DeleteLikedMovieResp struct {
}

type AllUserLikedMoviesReq struct {
	ID uint `path:"user_id"`
}

type AllUserAllLikedMoviesResp struct {
	LikedMoviesList []*LikedMovieInfo `json:"liked_movies"`
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
}

type CreateCustomListResp struct {
	ID    uint   `json:"list_id"`
	Title string `json:"title"`
}

type UpdateCustomListReq struct {
	ID    uint   `json:"list_id"`
	Title string `json:"title"`
}

type UpdateCustomListResp struct {
}

type DeleteCustomListReq struct {
	ID uint `json:"list_id"`
}

type DeleteCustomListResp struct {
}

type AllCustomListReq struct {
	ID uint `path:"user_id"`
}

type AllCustomListResp struct {
	Lists []ListInfo `json:"lists"`
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

type GetOneMovieFromListReq struct {
	ListID  uint `path:"list_id"`
	MovieID uint `path:"movie_id"`
}

type GetOneMovieFromListResp struct {
	Movie MovieInfo `json:"movie_info"`
}

type ListInfo struct {
	ID     uint        `json:"list_id"`
	Title  string      `json:"title"`
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

type PostsInfoReq struct {
}

type PostsInfoResp struct {
	Infos []PostInfo `json:"post_info"`
}

type PostInfosByIDReq struct {
	PostID uint `path:"post_id"`
}

type PostInfosByIDResp struct {
	Info []PostInfo `json:"post_info"`
}

type PostInfoReq struct {
	PostID uint `path:"post_id"`
}

type PostInfoResp struct {
	Info PostInfo `json:"post_info"`
}

type PostInfo struct {
	PostID           uint          `json:"id"`
	PostUser         PostUserInfo  `json:"user_info"`
	PostTitle        string        `json:"post_title"`
	PostDesc         string        `json:"post_desc"`
	PostMovie        PostMovieInfo `json:"post_movie_info"`
	PostLikeCount    int64         `json:"post_like_count"`
	PostCommentCount int64         `json:"post_comment_count"`
	CreateAt         int64         `json:"create_at"`
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
	CommentID uint  `json:"comment"`
	CreateAt  int64 `json:"create_At"`
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
}

type GetPostCommentsResp struct {
	Comments []CommentInfo `json:"comments"`
}

type CommentInfo struct {
	CommentID uint        `json:"comment_id"`
	PostID    uint        `json:"post_id"`
	UserInfo  CommentUser `json:"user_info"`
	Comment   string      `json:"comment"`
	UpdateAt  int64       `json:"update_at"`
}

type CommentUser struct {
	UserID     uint   `json:"id"`
	UserName   string `json:"name"`
	UserAvatar string `json:"avatar"`
}
