package shortener

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// URLShortenerDB that provides
// asdasd (just a text plug)
type URLShortenerDB struct {
	DataBase *sql.DB
}

//NewURLShortenerDB do something
func NewURLShortenerDB(connStr string) *URLShortenerDB {
	dataBase, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	shortener := &URLShortenerDB{DataBase: dataBase}
	err = shortener.CreateTable()
	if err != nil {
		panic(err)
	}
	return shortener
}

//CreateTable creates a table with Path, Url values set to nil values
func (db *URLShortenerDB) CreateTable() error { //

	//Create TABLE Query WITH `Path` AND `Url` VALUES
	createTableQuery := `
	DROP TABLE Shortener;
	CREATE TABLE Shortener (
		Path varchar(255),
		Url varchar(255)
	);`
	// Create table
	_, err := db.DataBase.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

//InsertLongShortData insert given map[string]string data into the database
func (db *URLShortenerDB) InsertLongShortData(data map[string]string) error {
	// Inserting data into the database:
	stmt, err := db.DataBase.Prepare(`INSERT INTO Shortener(Path, Url) VALUES( $1, $2 )`)
	if err != nil { //
		return err
	}
	for key, value := range data { //
		if _, err := stmt.Exec(key, value); err != nil {
			return err
		}
	}
	return nil
}

func (db *URLShortenerDB) String() {
	q := `
	SELECT Path, Url FROM Shortener
	`
	//Getting these values
	rows, err := db.DataBase.Query(q)

	//Printing that values
	defer rows.Close()

	for rows.Next() {
		var Path string
		var URL string
		for rows.Next() {
			err = rows.Scan(&Path, &URL)
			if err != nil {
				break
			}
			fmt.Print(Path, "\t", URL, "\n")
		}
	}

}

//GetLongURL gives long URL after receiving a short one
func (db *URLShortenerDB) GetLongURL(url string) string {
	//Getting Query
	rows, err := db.DataBase.Query("SELECT Url FROM Shortener WHERE Path = $1", url)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Printing that values

	var result string
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return result
}

// func main() { //
// URLsShort := map[string]string{"/vk": "https://vk.com/im?peers=194134042", "/stock": "https://ru.tradingview.com/chart/CPNxT7rI/?symbol=NYSE%3ANEE", "/timetable": "https://webservices.mirea.ru/upload/iblock/e38/2cu5la01h6zeg1la9upv0o4fphvt3izu/ИИИ_2 курс_21-22_весна.xlsx"}

// connStr := "user=alex password=222"

////////////New not parralel version
// shortener := newURLShortenerDB(connStr)
// err := shortener.CreateTable()
// if err != nil {
// 	panic(err)
// }
// err = shortener.insertLongShortData(URLsShort)
// if err != nil {
// 	panic(err)
// }
// // shortener.String()
// url := shortener.getLongURL("/vk")
// fmt.Println(url)

///////////////Last version////////////////////////////////

// ///Connect to database with `connStr` parameters
// db, err := sql.Open("postgres", connStr)
// if err != nil {
// 	panic(err)
// }

// //Ping the database if it is available, else "DataBase is not responding, err", err
// err = db.Ping()
// if err != nil {
// 	fmt.Println("DataBase is not responding, err:", err)
// }

// // Select Query to get all Path, Url values
// q := `
// SELECT Path, Url FROM Shortener
// `
// //Getting these values
// rows, err := db.Query(q)
// defer rows.Close()
// //Printing that values
// for rows.Next() {
// 	var Path string
// 	var URL string
// 	err = rows.Scan(&Path, &URL)
// 	if err != nil {
// 		break
// 	}
// 	fmt.Println(Path, URL)
// }

// }
