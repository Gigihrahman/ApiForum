package main

import (
	"forumapp-restapi/internal/configs"
	"forumapp-restapi/internal/handlers/memberships"
	"forumapp-restapi/internal/handlers/posts"
	membershipRepo "forumapp-restapi/internal/repository/memberships"
	postRepo "forumapp-restapi/internal/repository/posts"
	membershipSVC "forumapp-restapi/internal/service/memberships"
	postSVC "forumapp-restapi/internal/service/posts"
	"forumapp-restapi/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	
	if err!= nil{
		log.Fatal("gagal", err)
	}
	cfg = configs.Get()
	log.Println("config",cfg)
	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)
	if err !=nil{
		log.Fatal("gagal konek",err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSVC.NewService(cfg,membershipRepo)
	postService := postSVC.NewService(cfg,postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	postHandler:= posts.NewHandler(r,postService)
		
	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}