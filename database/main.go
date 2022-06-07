package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
// テーブル作成

type Person struct {
	Name string
	Age int
}

var DB *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	// テーブル作成
	// cmd := `CREATE TABLE IF NOT EXISTS person(
	// 	name STRING,
	// 	age INT)`
	// _, err := Db.Exec(cmd)

	// データの追加
	// Valuesを?にすることで置換できる
	// cmd := `INSERT INTO person(name, age) VALUES(?, ?)`
  // _, err := Db.Exec(cmd, "hanako", 19)

	// データの更新
	// cmd := `UPDATE person SET age = ? WHERE name = ?`
	// _, err := Db.Exec(cmd, 25, "taro")

	// データの取得
	// cmd := `SELECT * FROM person where age = ?`
	// row := Db.QueryRow(cmd, 25)
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
	// cmd := `SELECT * FROM person`
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
	cmd := `DELETE FROM person WHERE name = ?`
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatal(err)
	}

}

/*
go run main.goを実行するとexample.sqlが作成される
sqlite3 example.sqlとコマンド実行すると操作が可能
.exitで終了

*/