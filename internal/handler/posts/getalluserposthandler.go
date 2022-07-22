package posts

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ryantokmanmokmtm/movie-server/common/errx" //common error package
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/ryantokmanmokmtm/movie-server/internal/logic/posts"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
)

func GetAllUserPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostsInfoReq
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
			httpx.Error(w, errx.NewCommonMessage(errx.REQ_PARAM_ERROR, errs[0].Translate(trans)))
			return
		}

		l := posts.NewGetAllUserPostLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUserPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
