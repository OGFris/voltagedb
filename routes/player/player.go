package player

import (
	db "github.com/OGFris/voltagedb/database"
	"github.com/OGFris/voltagedb/utils"
	"net/http"
	"strconv"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
	xuid := r.PostFormValue("xuid")
	username := r.PostFormValue("username")

	db.Instance.Create(&db.Player{
		XUID:     xuid,
		Username: username,
	})
}

func Get(w http.ResponseWriter, r *http.Request) {
	xuid := r.PostFormValue("xuid")
	var player db.Player

	err := db.Instance.Preload("Bans").Where("xuid = ?", xuid).Find(&player).Error
	if err != nil {

		utils.WriteErr(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJson(w, player)
}

func Ban(w http.ResponseWriter, r *http.Request) {
	xuid := r.PostFormValue("xuid")
	hours, err := strconv.Atoi(r.PostFormValue("hours"))
	utils.PanicErr(err)

	var player db.Player

	err = db.Instance.First(&player, &db.Player{XUID: xuid}).Error
	if err != nil {

		utils.WriteErr(w, err.Error(), http.StatusNotFound)
		return
	}

	err = db.Instance.Create(&db.Ban{
		Until:    time.Now().Add(time.Hour * time.Duration(hours)),
		PlayerId: player.ID,
	}).Error
	if err != nil {

		utils.WriteErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
