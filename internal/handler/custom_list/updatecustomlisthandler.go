package custom_list

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/custom_list"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateCustomListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCustomListReq
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
			httpx.Error(w, errors.Wrap(errx.NewErrCode(errx.REQ_PARAM_ERROR), fmt.Sprintf("Validated err: %v", errs[0].Translate(trans))))
			return
		}

		l := custom_list.NewUpdateCustomListLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCustomList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
