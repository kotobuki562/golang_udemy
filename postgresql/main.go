package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=test_db password=password sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	defer Db.Close()

	// データの追加
	// Valuesを$にすることで置換できる
	// cmd := `INSERT INTO persons(name, age) VALUES($1, $2)`
  // _, err := Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Println(err)
	// }

	// データの更新
	// cmd := `UPDATE persons SET age = $1 WHERE name = $2`
	// _, err := Db.Exec(cmd, 25, "hanako")
	// if err != nil {
	// 	log.Println(err)
	// }

	// データの取得
	// cmd := `SELECT * FROM persons where age = $1`
	// row := Db.QueryRow(cmd, 19)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No rows were returned!")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

	// 複数のデータの取得
	// cmd = `SELECT * FROM persons`
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p)
	// }

	// データの削除
		cmd := `DELETE FROM persons WHERE name = $1`
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatal(err)
	}
}

/*
	psqlでデータベースの操作
	\lで作成済みのDBを確認する
	\qで終了する
*/