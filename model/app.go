package model

import (
	"github.com/beratn/sho/client"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type App struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (a *App) Initialize() {
	a.DB = client.InitDb()
	a.DB.AutoMigrate(&Link{})
	a.Redis = client.InitRedis()
}
