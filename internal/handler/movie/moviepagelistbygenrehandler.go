package movie

import (
	"github.com/go-playground/validator/v10"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/movie"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func MoviePageListByGenreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MoviePageListByGenreRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//simple validate
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := movie.NewMoviePageListByGenreLogic(r.Context(), svcCtx)
		resp, err := l.MoviePageListByGenre(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
