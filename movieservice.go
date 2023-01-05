package main

import (
	"flag"
	"fmt"
	//_ "github.com/joho/godotenv"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/internal/dao"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/server"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/movieservice.yaml", "the config file")

func main() {
	//_ = godotenv.Load(".env")

	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	fmt.Println(c.MySQL.DataSource)
	ser := server.SetUpEngine(c, dao.NewDAO(models.NewEngine(c)))
	defer ser.Start()
	ser.Start()
}
