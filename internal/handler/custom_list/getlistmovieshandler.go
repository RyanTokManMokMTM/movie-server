package custom_list

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ryantokmanmokmtm/movie-server/common/errx" //common error package
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/ryantokmanmokmtm/movie-server/internal/logic/custom_list"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
)

func GetListMoviesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetListMoviesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		eng := en.New()
		uti := ut.New(eng, eng)
		trans, _ := uti.GetTranslator("en")
		validate := validator.New()
		en_translations.RegisterDefaultTranslations(validate, trans)

		if err := validate.StructCtx(r.Context(), req); err != nil {
			errs := err.(validator.ValidationErrors)
			commonErr := errx.NewCommonMessage(errx.REQ_PARAM_ERROR, errs[0].Translate(trans))
			httpx.WriteJson(w, commonErr.StatusCode(), commonErr.ToJSONResp())
			return
		}

		l := custom_list.NewGetListMoviesLogic(r.Context(), svcCtx)
		resp, err := l.GetListMovies(&req)

		if err != nil {
			if r, ok := err.(*errx.CommonError); ok {
				httpx.WriteJson(w, r.StatusCode(), r.ToJSONResp())
			} else {
				httpx.Error(w, err)
			}
		} else {
			httpx.OkJson(w, resp)
		}

	}
}
