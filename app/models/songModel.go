package models

import (
	"database/sql"
)

type (
	SongModel struct {
		DbHandle *sql.DB
	}

	songData struct {
		Id      string
		Title   string
		Author  string
		Album   string
		Genre   string
		AddedBy string
	}
)

func (sm *SongModel) GetAll() ([]*songData, error) {
	songs := []*songData{}

	stmt, err := sm.DbHandle.Prepare("SELECT * FROM songs")
	if err != nil {
		return songs, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return songs, err
	}

	var id, title, author, album, genre, addedby string
	for rows.Next() {
		rows.Scan(&id, &title, &author, &album, &genre, &addedby)
		songs = append(songs, &songData{Id: id, Title: title, Author: author, Genre: genre,
			AddedBy: addedby})
	}

	return songs, nil
}

func (sm *SongModel) Get(id string) {

}
