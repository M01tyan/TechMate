package modules

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

type Post struct {
    NAME        string
    STUDENT_ID  string
}

var Db *sql.DB

func init() {
    var err error
    Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        panic(err)
        Db.Close()
    }
}

func GetPost(genre []string) (complete_es []Post, err error) {
	for _, g := range genre {
	    rows, err := Db.Query("SELECT users.name, users.student_id FROM users LEFT JOIN users_genres ON users.id = users_genres.user_id LEFT JOIN genres ON users_genres.genre_id = genres.id WHERE genres.name = $1", g)
	    if err != nil {
	        fmt.Println(err)
	    }

	    for rows.Next() {
	        var e Post
	        rows.Scan(&e.NAME, &e.STUDENT_ID)
	        for _, es := range complete_es {
	        	if es.STUDENT_ID != e.STUDENT_ID {
	        		complete_es = append(complete_es, e)
	        	}
	        }
	    }
	}
    return
}

func InsertData(name string, line_id string, student_id string, genre []string) {
    var id int
    err := Db.QueryRow("INSERT INTO users (name, line_id, student_id) VALUES ($1, $2, $3) RETURNING id", name, line_id, student_id).Scan(&id)
    if err != nil {
        fmt.Println(err)
    }/*
    var user_id int
    row.Scan(&user_id)*/
    for _, g := range genre {
        var genre_id int
        errs := Db.QueryRow("SELECT id FROM genres WHERE name=$1", g).Scan(&genre_id)
        if errs != nil {
            fmt.Println(errs)
        }
        errors := Db.QueryRow("INSERT INTO users_genres (user_id, genre_id) VALUES ($1, $2) RETURNING user_id", id, genre_id)
        if errors != nil {
            fmt.Println("error", errors)
        }
    }
}

    /*
func main() {
    db, _ := sql.Open("postgres", "user=m01tyan password=No.1runner dbname=techmate sslmode=disable")
    defer db.Close()

    rows, err := db.Query("SELECT user_table.id, user_table.name, user_table.line_id, user_table.student_id, genre_table.name FROM user_table LEFT JOIN user_genre_table ON user_table.id = user_genre_table.user_id LEFT JOIN genre_table ON user_genre_table.genre_id = genre_table.id")
    rows2, err2 := db.Query("SELECT user_table.name, user_table.student_id FROM user_table LEFT JOIN user_genre_table ON user_table.id = user_genre_table.user_id LEFT JOIN genre_table ON user_genre_table.genre_id = genre_table.id WHERE genre_table.name = 'Python'")

    if err != nil {
        fmt.Println(err)
    }

    var es []GENRE
    for rows.Next() {
        var e GENRE
        rows.Scan(&e.ID, &e.NAME, &e.LINE_ID, &e.STUDENT_ID, &e.GENRE_NAME)
        es = append(es, e)
    }
    fmt.Println("%v", es)

    for rows2.Next() {
        var s_id string
        var s_name string
        rows2.Scan(&s_name, &s_id)
        fmt.Println(s_name, s_id)
    }
    genre := []string{"Python", "Swift", "Deep Learning"}
    InsertData("滝原　航大", "df9s8daf78ddaf7sd8f9sa8df7s9d8afsdasdfa", "s1240230", genre)
    readPost, _ := GetPost("Python")
    fmt.Println(readPost)
}
    */

// 関数名に仮でjsonの名前を使っていますが、dbを使っても大丈夫です。
// やりやすい方でお願いします。関数名の適宜変更、追加などお願いします。
//func json_manager(){}
