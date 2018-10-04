package modules

import (
    "database/sql"
    "fmt"
    "os"
    "log"

    _ "github.com/lib/pq"
)

type Post struct {
    NAME        string
    STUDENT_ID  string
}

var Db *sql.DB

func OpenDB() (Db *sql.DB) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    return
}

func GetLineID(line_id string) (mode string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    err := Db.QueryRow("SELECT modes.name FROM users LEFT JOIN user_mode ON users.id = user_mode.user_id LEFT JOIN modes ON user_mode.mode_id = modes.id WHERE users.line_id = $1", line_id).Scan(&mode)
    if err != nil {
        log.Println(err)
    }
    if mode == nil {
        var id int
        _, err := Db.Exec("INSERT INTO users (line_id, mode) VALUES ($1, $2)", line_id, 1)
        if err != nil {
            fmt.Println(err)
        }
        mode = "init_new"
    }
    return
}

func UpdateMode(mode_int int, line_id string) (mode string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    _, errs := Db.Exec("UPDATE users SET mode=$1 WHERE line_id=$2", mode_int, line_id)
    if errs != nil {
        log.Print(errs)
    }
    err := Db.QueryRow("SELECT modes.name FROM users LEFT JOIN user_mode ON users.id = user_mode.user_id LEFT JOIN modes ON user_mode.mode_id = modes.id WHERE users.line_id = $1", line_id).Scan(&mode)
    if err != nil {
        log.Println(err)
    }
    return
}

func InsertGenre(genre string, line_id string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    var user_id int
    err := Db.QueryRow("SELECT users.id FROM users LEFT JOIN user_mode ON users.id = user_mode.user_id LEFT JOIN modes ON user_mode.mode_id = modes.id WHERE users.line_id = $1", line_id).Scan(&user_id)
    if err != nil {
        log.Println(err)
    }
    var genre_id int
    err := Db.QueryRow("SELECT genres.id FROM users LEFT JOIN user_mode ON users.id = user_mode.user_id LEFT JOIN modes ON user_mode.mode_id = modes.id WHERE genres.name = $1", genre).Scan(&genre_id)
    if err != nil {
        log.Println(err)
    }
    _, errs := Db.Exec("INSERT INTO user_genre (user_id, genre_id) VALUES ($1, $2)", user_id, genre_id)
    if errs != nil {
        log.Print(errs)
    }
}

func InsertName(name string, line_id string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    _, errs := Db.Exec("UPDATE users SET name=$1 WHERE line_id=$2", name, line_id)
    if errs != nil {
        log.Print(errs)
    }
}

func InsertStudentID(student_id string, line_id string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    _, errs := Db.Exec("UPDATE users SET student_id=$1 WHERE line_id=$2", student_id, line_id)
    if errs != nil {
        log.Print(errs)
    }
}

func GetPost(genre string) (result string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
	rows, err := Db.Query("SELECT users.name, users.student_id FROM users LEFT JOIN user_genre ON users.id = user_genre.user_id LEFT JOIN genres ON user_genre.genre_id = genres.id WHERE genres.name = $1", genre)
    if err != nil {
        log.Println(err)
    }

    var complete_es []Post
    for rows.Next() {
        var e Post
	    rows.Scan(&e.NAME, &e.STUDENT_ID)
	    complete_es = append(complete_es, e)
	}
    for _, r := range complete_es {
        result += r.NAME + "\t" + r.STUDENT_ID + "\n"
    }
    return
}

func GetGenres() (genre_names []string) {
    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Print(err)
        Db.Close()
    }
    row, errs := Db.Query("SELECT name FROM genres")
    if errs != nil {
        log.Print(errs)
    }

    for row.Next() {
        var e string
        row.Scan(&e)
        genre_names = append(genre_names, e)
    }
    log.Print(genre_names)
    return
}

func InsertData(name string, line_id string, student_id string, genre []string) {
    var id int
    Db, errs := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errs != nil {
        log.Print(errs)
        Db.Close()
    }
    err := Db.QueryRow("INSERT INTO users (name, line_id, student_id) VALUES ($1, $2, $3) RETURNING id", name, line_id, student_id).Scan(&id)
    if err != nil {
        fmt.Println(err)
    }
    for _, g := range genre {
        var genre_id int
        errs := Db.QueryRow("SELECT id FROM genres WHERE name=$1", g).Scan(&genre_id)
        if errs != nil {
            fmt.Println(errs)
        }
        errors := Db.QueryRow("INSERT INTO user_genre (user_id, genre_id) VALUES ($1, $2) RETURNING user_id", id, genre_id)
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