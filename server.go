// Displays a list of values from a database on a webpage, and gives the
// ability to add values to the database with an input field and submit button
// on the same page. 


package main


import (
    "fmt"
    "net/http"
    "database/sql"
    "log"
    _ "github.com/lib/pq"
    "strconv"
)


func handler(w http.ResponseWriter, r *http.Request) {
    var (
        count int
        rows *sql.Rows
    )

    db, err := sql.Open("postgres",
        "user=username password=password dbname=databasename host=localhost")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if r.Method == "POST" {
        fmt.Println("It's a POST.")
        r.ParseForm()
        for k, v := range r.PostForm {
            fmt.Println("Key:", k, "Value:", v)
            if k == "new_value" {
                fmt.Println("inserting..")
                fmt.Printf("Length of input is %d\n", len(v))
                asInt, err := strconv.Atoi(v[0])
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Println(fmt.Sprintf("insert into sphere values(%d)", asInt))
                db.Exec(fmt.Sprintf("insert into sphere values(%d)", asInt))
            }
        }
    }

    rows, err = db.Query("select count from sphere")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Fprintln(w, "<html><body><h1>The values</h1>")
    for rows.Next() {
        err := rows.Scan(&count)
        if err != nil {
            log.Fatal(err)
        }
        asInt := strconv.Itoa(count)
        fmt.Fprintln(w, asInt)
    }
    fmt.Fprintln(w, "<br><h1>The input</h1><form action='.' method='post'><input type='text' name='new_value'/><input type='submit' value='Submit'></form></body></html>")
}


func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
