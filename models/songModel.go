package models

import (
	"database/sql"
)

type (
	SongModel struct {
		DBHandle *sql.DB
	}

	song struct {
		Id      string
		Title   string
		Author  string
		Album   string
		Genre   string
		AddedBy string
	}
)

func (sm *SongModel) GetAll() ([]*song, error) {
	songs := []*song{}

	stmt, err := sm.DBHandle.Prepare("SELECT * FROM songs")
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
		songs = append(songs, &song{Id: id, Title: title, Author: author, Genre: genre,
			AddedBy: addedby})
	}

	return songs, nil
}

func (sm *SongModel) Get(id string) {

}
