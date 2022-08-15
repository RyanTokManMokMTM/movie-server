package custom_list

import (
	"net/http"

	"github.com/ryantokmanmokmtm/movie-server/internal/logic/custom_list"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InsertMovieToListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertMovieReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := custom_list.NewInsertMovieToListLogic(r.Context(), svcCtx)
		resp, err := l.InsertMovieToList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
