// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/ryantokmanmokmtm/movie-server/internal/handler/comment"
	custom_list "github.com/ryantokmanmokmtm/movie-server/internal/handler/custom_list"
	health "github.com/ryantokmanmokmtm/movie-server/internal/handler/health"
	likedMovie "github.com/ryantokmanmokmtm/movie-server/internal/handler/likedMovie"
	movie "github.com/ryantokmanmokmtm/movie-server/internal/handler/movie"
	posts "github.com/ryantokmanmokmtm/movie-server/internal/handler/posts"
	user "github.com/ryantokmanmokmtm/movie-server/internal/handler/user"
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
				Handler: likedMovie.CreateLikedMovieHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/liked/movie",
				Handler: likedMovie.DeleteLikedMovieHandler(serverCtx),
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
				Method:  http.MethodGet,
				Path:    "/list/:list_id/movie/:movie_id",
				Handler: custom_list.GetOnlyMovieFromListHandler(serverCtx),
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
				Path:    "/posts",
				Handler: posts.GetAllUserPostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/posts/all",
				Handler: posts.GetAllPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/:post_id",
				Handler: posts.GetPostByPostIDHandler(serverCtx),
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
		},
		rest.WithPrefix("/api/v1"),
	)
}
