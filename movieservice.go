package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/internal/handler"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/movieservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		if ok := errors.As(err, errx.CommonError{}); ok {
			return http.StatusOK, err.(*errx.CommonError).ToJSONResp()
		} else {
			return http.StatusInternalServerError, nil
		}

	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//
//func validatorTrans() *ut.Translator {
//	eng := en.New()
//	uni := ut.New(eng, eng)
//
//	trans, _ := uni.GetTranslator("en")
//	validate := validator.New()
//
//	en_translations.RegisterDefaultTranslations(validate, trans)
//	return &trans
//}
