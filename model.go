package main

import (
	"log"

	"gorm.io/gorm"
)

type link struct {
	gorm.Model
	Target  string `json:"target"`
	Address string `json:"address"`
}

func (l *link) getLink(id int) {
	db := GetDb()
	db.First(&l, id)
}

func (l *link) getTargetById(id string) {
	db := GetDb()
	db.Where("address = ?", id).First(&l)
}

func (l *link) createLink() {
	db := GetDb()
	result := db.Select("Target").Create(&l)

	if result == nil {
		log.Fatal("An error occurred while creating link")
	}
}

func CheckAddressIsExists(id string) bool {
	db := GetDb()
	l := link{}
	result := db.Where("address = ?", id).First(&l)

	return result.RowsAffected != 0
}
