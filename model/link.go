package model

import (
	"github.com/beratn/sho/client"
	"gorm.io/gorm"
	"log"
)

type Link struct {
	gorm.Model
	Target  string `json:"target"`
	Address string `json:"address"`
}

func (l *Link) getLink(id int) {
	db := client.GetDb()
	db.First(&l, id)
}

func (l *Link) GetTargetById(id string) {
	db := client.GetDb()
	db.Where("address = ?", id).First(&l)
}

func (l *Link) CreateLink() {
	db := client.GetDb()
	result := db.Create(&l)

	if result == nil {
		log.Fatal("An error occurred while creating link")
	}
}

func (l *Link) SetCache() {
	redisClient := client.GetRedisClient()
	redisClient.Set(l.Address, l.Target, 0)
}

func CheckAddressIsExists(id string) bool {
	db := client.GetDb()
	l := Link{}
	result := db.Where("address = ?", id).First(&l)

	return result.RowsAffected != 0
}
