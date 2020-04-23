package database

type Player struct {
	Model
	Username string `gorm:"Type:varchar(255);Column:username;NOT NULL" json:"username"`
	Rank     uint   `gorm:"Type:int(10) unsigned;Column:rank;NOT NULL;DEFAULT:0" json:"rank"`
	Coins    int    `gorm:"Type:int(10);Column:coins;NOT NULL;DEFAULT:0" json:"coins"`
	XUID     string `gorm:"Type:varchar(255);Column:xuid;NOT NULL" json:"xuid"`
	Bans     []Ban  `json:"bans"`
}
