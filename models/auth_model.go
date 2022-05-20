package models

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/djarum76-bot/taskplanner_golang/config"
	"github.com/djarum76-bot/taskplanner_golang/db"
	"github.com/djarum76-bot/taskplanner_golang/helper"

	"github.com/golang-jwt/jwt"
)

func Register(username string, password string) (ResponseToken, bool, error) {
	var user User
	var err error
	var res ResponseToken

	con := db.CreateCon()
	conf := config.GetConfig()

	sqlStatement := `INSERT INTO users (username, password, image, role) VALUES ($1, $2, $3, $4)`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, false, err
	}

	_, err = stmt.Exec(username, password, nil, nil)
	if err != nil {
		return res, false, err
	}

	user.Username = username

	claims := &JwtCustomClaims{
		user.Id,
		user.Username,
		user.Image.String,
		user.Role.String,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		return res, false, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Register Success"
	res.Data = user
	res.Token = t

	return res, true, nil
}

func Login(username string, password string) (ResponseToken, bool, error) {
	var user User
	var pwdHash string
	var res ResponseToken

	con := db.CreateCon()
	conf := config.GetConfig()

	sqlStatement := "SELECT * FROM users WHERE username = ($1)"

	err := con.QueryRow(sqlStatement, username).Scan(&user.Id, &user.Username, &pwdHash, &user.Image, &user.Role)
	if err == sql.ErrNoRows {
		return res, false, err
	}
	if err != nil {
		return res, false, err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, false, err
	}

	claims := &JwtCustomClaims{
		user.Id,
		user.Username,
		user.Image.String,
		user.Role.String,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		return res, false, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Login"
	res.Data = user
	res.Token = t

	return res, true, nil
}
