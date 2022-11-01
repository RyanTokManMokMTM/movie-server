// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/ryantokmanmokmtm/movie-server/internal/handler/comment"
	comment_likes "github.com/ryantokmanmokmtm/movie-server/internal/handler/comment_likes"
	custom_list "github.com/ryantokmanmokmtm/movie-server/internal/handler/custom_list"
	friend "github.com/ryantokmanmokmtm/movie-server/internal/handler/friend"
	health "github.com/ryantokmanmokmtm/movie-server/internal/handler/health"
	likedMovie "github.com/ryantokmanmokmtm/movie-server/internal/handler/likedMovie"
	message "github.com/ryantokmanmokmtm/movie-server/internal/handler/message"
	movie "github.com/ryantokmanmokmtm/movie-server/internal/handler/movie"
	notification "github.com/ryantokmanmokmtm/movie-server/internal/handler/notification"
	post_likes "github.com/ryantokmanmokmtm/movie-server/internal/handler/post_likes"
	posts "github.com/ryantokmanmokmtm/movie-server/internal/handler/posts"
	room "github.com/ryantokmanmokmtm/movie-server/internal/handler/room"
	user "github.com/ryantokmanmokmtm/movie-server/internal/handler/user"
	user_genre "github.com/ryantokmanmokmtm/movie-server/internal/handler/user_genre"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: health.HealthCheckHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/signup",
				Handler: user.UserSignUpHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info/:id",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/friends/count/:user_id",
				Handler: user.CountFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/friends/list/:user_id",
				Handler: user.GetFriendListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: user.UserProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/profile",
				Handler: user.UpdateUserProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/avatar",
				Handler: user.UploadUserAvatarHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/cover",
				Handler: user.UploadUserCoverHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/friends/room",
				Handler: user.GetFriendRoomListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/movies/list/:genre_id",
				Handler: movie.MoviePageListByGenreHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/movies/genres/:movie_id",
				Handler: movie.MovieGenreByMovieIDHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/movies/:movie_id",
				Handler: movie.GetMovieDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/movie/count/liked/:movie_id",
				Handler: movie.GetMovieLikedCountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/movie/count/collected/:movie_id",
				Handler: movie.GetMovieCollectedCountHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/liked/movies/:user_id",
				Handler: likedMovie.GetUserLikedMovieListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/liked/movie",
				Handler: likedMovie.LikedMovieHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/liked/movie",
				Handler: likedMovie.RemoveLikedMovieHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/liked/movie/:movie_id",
				Handler: likedMovie.IsLikedMovieHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/lists/:user_id",
				Handler: custom_list.GetAllUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list/:list_id",
				Handler: custom_list.GetListByIDHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/lists",
				Handler: custom_list.CreateCustomListHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/lists",
				Handler: custom_list.UpdateCustomListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/lists",
				Handler: custom_list.DeleteCustomListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list/:list_id/movie/:movie_id",
				Handler: custom_list.InsertMovieToListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/list/:list_id/movie/:movie_id",
				Handler: custom_list.RemoveMovieFromListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/list/movies/:id",
				Handler: custom_list.RemoveListMoviesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list/movie/:movie_id",
				Handler: custom_list.GetOneMovieFromUserListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/posts",
				Handler: posts.CreatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/posts",
				Handler: posts.UpdatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/posts",
				Handler: posts.DeletePostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/all",
				Handler: posts.GetAllPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/follow",
				Handler: posts.GetFollowingPostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/post/:post_id",
				Handler: posts.GetPostByPostIDHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/:user_id",
				Handler: posts.GetUserPostsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/posts/count/:user_id",
				Handler: posts.CountAllUserPostHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comments/:post_id",
				Handler: comment.CreateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comments/:post_id/reply/:comment_id",
				Handler: comment.CreateReplyCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/comments/:comment_id",
				Handler: comment.UpdateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/comments/:comment_id",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/comments/:post_id",
				Handler: comment.GetPostCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comments/reply/:comment_id",
				Handler: comment.GetReplyCommentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/friend",
				Handler: friend.AddFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friend/requests",
				Handler: friend.GetFriendRequestHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/friend",
				Handler: friend.RemoveFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/friend/request/accept",
				Handler: friend.AcceptFriendRequestHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/friend/request/cancel",
				Handler: friend.CancelFriendRequestHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/friend/request/decline",
				Handler: friend.DeclineFriendRequestHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friend/:user_id",
				Handler: friend.IsFriendHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/liked/comment",
				Handler: comment_likes.CreateCommentLikesHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/liked/comment",
				Handler: comment_likes.RemoveCommentLikesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/liked/comment/:comment_id",
				Handler: comment_likes.IsCommentLikedHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/liked/comment/count/:comment_id",
				Handler: comment_likes.CountCommentLikesHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/liked/post",
				Handler: post_likes.CreatePostLikesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/liked/post",
				Handler: post_likes.RemovePostLikesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/liked/post/:post_id",
				Handler: post_likes.IsPostLikedHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/liked/post/count/:post_id",
				Handler: post_likes.CountPostLikesHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPatch,
				Path:    "/user/genres",
				Handler: user_genre.UpdateUserGenreHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/genres/:user_id",
				Handler: user_genre.GetUserGenreHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/room",
				Handler: room.CreateRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/room",
				Handler: room.DeleteRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/room/join/:room_id",
				Handler: room.JoinRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/room/leave/:room_id",
				Handler: room.LeaveRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/room/members/:room_id",
				Handler: room.RoomMembersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/room/rooms",
				Handler: room.GetUserRoomsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/room/:room_id",
				Handler: room.GetRoomINfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/room/:room_id/read",
				Handler: room.UpdateIsReadHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/message/:room_id",
				Handler: message.GetRoomMessageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notification/likes",
				Handler: notification.GetlikenotificationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notification/comments",
				Handler: notification.GetcommentnotificationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
