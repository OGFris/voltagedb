package database

import "time"

type Ban struct {
	Model
	Until    time.Time `gorm:"Type:datetime;Column:until;NOT NULL" json:"until"`
	PlayerId uint      `gorm:"Type:int(10) unsigned;Column:player_id;NOT NULL" json:"player_id"`
}
