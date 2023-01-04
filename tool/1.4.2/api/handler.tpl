package {{.PkgName}}

import (
	"net/http"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/ryantokmanmokmtm/movie-server/common/errx" //common error package

	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
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


		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})


		if err != nil {
            if r, ok := err.(*errx.CommonError); ok {
                httpx.WriteJson(w, r.StatusCode(), r.ToJSONResp())
            } else {
                httpx.Error(w, err)
            }
        } else {
           {{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
        }

	}
}
