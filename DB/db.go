package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() { //
	URLsShort := map[string]string{"/vk": "https://vk.com/im?peers=194134042", "/stock": "https://ru.tradingview.com/chart/CPNxT7rI/?symbol=NYSE%3ANEE", "/timetable": "https://webservices.mirea.ru/upload/iblock/e38/2cu5la01h6zeg1la9upv0o4fphvt3izu/ИИИ_2 курс_21-22_весна.xlsx"}

	_ = URLsShort
	connStr := "user=alex password=222"
	fmt.Println("Hello SQL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DataBase is not responding, err:", err)
	}

	createTableQuery := `
	DROP TABLE Shortener;
	CREATE TABLE Shortener (
		Path varchar(255),
		Url varchar(255)
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	q := `
	SELECT * FROM Shortener
	`
	rows, err := db.Query(q)
	defer rows.Close()
	col, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(col)

	// Inserting data into the database:
	stmt, err := db.Prepare(`INSERT INTO Shortener(Path, Url) VALUES( $1, $2 )`)
	if err != nil { //
		panic(err)
	}
	for key, value := range URLsShort { //
		if _, err := stmt.Exec(key, value); err != nil {
			fmt.Println(err)
		}
	}

	q = `
	SELECT Path, Url FROM Shortener
	`
	rows, err = db.Query(q)
	defer rows.Close()
	for rows.Next() {
		var Path string
		var URL string
		err = rows.Scan(&Path, &URL)
		if err != nil {
			break
		}
		fmt.Println(Path, URL)
	}

}
