package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./historytemp.db"

const createHistDataSttm = `
				create table if not exists HistoryData (
					id integer not null primary key autoincrement,
					avg real,
					min real,
					max real,
					date text
					);
				`
const addSttm = `
				 insert into HistoryData(avg, min, max, date)
				 values(?, ?, ?, ?);
				`
const getDataSttm = "select avg, min, max, date from HistoryData"

var db *sql.DB

func InitDb() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	sqlSttm := createHistDataSttm

	_, err = db.Exec(sqlSttm)
	if err != nil {
		log.Fatal(err)
	}
}

func AddData(avg float32, min float32, max float32, date string) {
	_, err := db.Exec(addSttm, avg, min, max, date)
	if err != nil {
		log.Fatal(err)
	}
}

type Dbdata struct {
	Avg  float32
	Min  float32
	Max  float32
	Date string
}

func GetAllDatas() []Dbdata {
	rows, err := db.Query(getDataSttm)
	if err != nil {
		log.Fatal(err)
	}
	var datas []Dbdata
	for rows.Next() {
		var avg float32
		var min float32
		var max float32
		var date string
		err := rows.Scan(&avg, &min, &max, &date)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, Dbdata{avg, min, max, date})
	}
	return datas
}
