package server

import (
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/internal/dao"
	"github.com/ryantokmanmokmtm/movie-server/internal/handler"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/serverWs"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func SetUpEngine(c config.Config, dao dao.Store) *rest.Server {

	//defer server.Stop()
	server := rest.MustNewServer(c.RestConf)
	ctx := svc.NewServiceContext(c, dao)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errx.CommonError:
			return http.StatusOK, e.ToJSONResp()
		default:
			return http.StatusInternalServerError, errx.NewCommonMessage(errx.SERVER_COMMON_ERROR, err.Error()).ToJSONResp()
		}
	})

	//Adding Static Route
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/resources/:file",
		Handler: http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))).ServeHTTP,
	})

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ws",
		Handler: serverWs.NewServerWS(ctx),
	}, rest.WithJwt(ctx.Config.Auth.AccessSecret))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	return server
}
