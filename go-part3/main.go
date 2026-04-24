package main

import (
	"database/sql"
	"fmt"
	// "os"
    "log"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
    ID int64
    Title string
    Artist string
    Price float32
}

func main() {
    log.SetPrefix("db-access: ")
	log.SetFlags(0)

	cfg := mysql.NewConfig()
	// cfg.User = os.Getenv("DBUSER")
	// cfg.Passwd = os.Getenv("DBPASS")
	cfg.User = "root"
	cfg.Passwd = "ABab1234!"
	cfg.Net = "tcp"
	cfg.Addr = "vps:3306"
	cfg.DBName = "recordings"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

    // 查询数据
    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

    fmt.Println("============")

    albumItem, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found: %v\n", albumItem)

    fmt.Println("==============")

    albumId, err := addAlbum(Album{
        Title: "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price: 49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %+v\n", albumId)
}

func albumsByArtist(artistName string) ([]Album, error) {
    var albums []Album
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artistName)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
    }

    defer rows.Close()
    for rows.Next() {
        var albumItem Album
        if err := rows.Scan(&albumItem.ID, &albumItem.Title, &albumItem.Artist, &albumItem.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
        }
        albums = append(albums, albumItem)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist, %q: %v", artistName, err)
    }
    
    return albums, nil
}

func albumByID(id int64) (Album, error) {
    var albumItem Album
    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&albumItem.ID, &albumItem.Title, &albumItem.Artist, &albumItem.Price); err != nil {
        if err == sql.ErrNoRows {
            return albumItem, fmt.Errorf("albumsByID %d: no sub album", id)
        }
        return albumItem, fmt.Errorf("albumsByID %d: %v", id, err)
    }
    return albumItem, nil
}

func addAlbum(album Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %+v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum, %+v", err)
    }

    return id, nil
}
