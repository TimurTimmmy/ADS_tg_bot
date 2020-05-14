package AdsBot

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Ads struct {
	adsId    int
	adsName  string
	adsPrice int
	adsUser  int
}

type Users struct {
	userId      int
	userName    string
	userAddress string
	userAdmin   int
}

func DBRequest(off int) []Ads {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Ads order by 1 limit 5 offset $1", off)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	ads := []Ads{}

	for rows.Next() {
		a := Ads{}
		err := rows.Scan(&a.adsId, &a.adsName, &a.adsPrice, &a.adsUser)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ads = append(ads, a)
	}
	return ads
}

func GetUsers() []Users {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []Users{}

	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.userId, &u.userName, &u.userAddress, &u.userAdmin)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	return users
}

func UpdateUserName(userid int, username string) string {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update Users set userName = $1 where userId = $2", username, userid)
	if err != nil {
		panic(err)
	}
	log.Print(result.RowsAffected())
	return "Новое имя " + username
}

func UpdateUserAddress(userid int, useraddres string) string {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update Users set userAddress = $1 where userId = $2", useraddres, userid)
	if err != nil {
		panic(err)
	}
	log.Print(result.RowsAffected())
	return "Новый адрес =" + useraddres
}

func InsertUser(userid int, username string) (res string) {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Users (userId, userName, userAddress, userAdmin) values ($1, $2, 0, 0)", userid, username)
	if err != nil {
		res = "Ошибка"
		log.Println(err)
	}
	res = "Добавлен"
	log.Print(result)
	return res
}

func DeleteUser(userid int) {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from Users where userId = $1", userid)
	if err != nil {
		panic(err)
	}
	log.Print(result.RowsAffected())
}

func FindUser(id int) Users {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Users where userId = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := Users{}

	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.userId, &u.userName, &u.userAddress, &u.userAdmin)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = u
	}
	return users
}

func CountAds() string {
	db, err := sql.Open("sqlite3", "./DB/ads.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var count string
	rows := db.QueryRow("select count(*) from Ads")
	rows.Scan(&count)
	return count
}
