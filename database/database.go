package database

import (
	"github.com/OGFris/voltagedb/utils"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

var Instance *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"Default:null" sql:"index" json:"-"`
}

func InitDB() {
	if Instance == nil {
		var d *gorm.DB

		sql := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_ADDRESS") + ":3306)/" + os.Getenv("DB_NAME")

		d, err := gorm.Open("mysql", sql+"?charset=utf8&parseTime=True&loc=Local")
		utils.PanicErr(err)
		Instance = d
		MigrateDB()
	}
}

func MigrateDB() {
	Instance.AutoMigrate(
		&Player{},
		&Ban{},
	)

	Instance.Model(&Ban{}).AddForeignKey("player_id", "players(id)", "RESTRICT", "RESTRICT")
}
