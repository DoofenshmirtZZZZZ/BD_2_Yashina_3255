package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")

var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CheckTable() error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Goods (Id SERIAL PRIMARY KEY, UserId BIGINT, Text TEXT)"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()
	return nil
}

func Set_product(userId int64, text string) error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("INSERT INTO Goods (userid, text) VALUES ('" + strconv.FormatInt(userId, 10) + "', '" + text + "')"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()

	return nil
}

func Delete_product(userId int64, text string) error {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if _, err := db.Exec("DELETE FROM Goods WHERE userid ='" + strconv.FormatInt(userId, 10) + "' AND text ='" + text + "'"); err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer db.Close()

	return nil
}

func Get_list_product(userId int64) ([]string, error) {
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		rows, err := db.Query("SELECT * FROM Goods WHERE userid = '" + strconv.FormatInt(userId, 10) + "'")
		defer rows.Close()

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			goods := []string{}

			for rows.Next() {
				pr := Product{}
				err := rows.Scan(&pr.id, &pr.userId, &pr.text)

				if err != nil {
					fmt.Println(err)
					continue
				}

				goods = append(goods, pr.text)
			}

			return goods, nil
		}
	}
}

type Product struct {
	id     int
	userId int64
	text   string
}
