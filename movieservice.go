package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	//_ "github.com/joho/godotenv"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/internal/dao"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/server"
)

var configFile = flag.String("f", "etc/movieservice.yaml", "the config file")

func main() {
	//_ = godotenv.Load(".env")

	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//fmt.Println(c.MySQL.DataSource)
	ser := server.SetUpEngine(c, dao.NewDAO(models.NewEngine(c)))
	defer ser.Start()
	ser.Start()
}
