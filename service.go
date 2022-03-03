package main

import (
	"encoding/json"
	"fmt"
	"kredit_bee_project/util"
	"net/http"
	"strconv"
)

type Album struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type Photo struct {
	Id           int    `json:"id"`
	AlbumId      int    `json:"albumId"`
	PhotoId      int    `json:"photoId"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	typeValue := v.Get("type")

	if typeValue == "album" {
		idValue := v.Get("id")
		id, err := strconv.Atoi(idValue)
		if err != nil {
			JSONMarshalBadRequest(err.Error(), 500).WriteToResponse(w)
			return
		}
		album, err := fetchAlbumById(id)
		if err != nil {
			JSONMarshalBadRequest(err.Error(), 500).WriteToResponse(w)
			return
		}

		json.NewEncoder(w).Encode(album)
		return

	} else if typeValue == "photo" {

		idValue := v.Get("id")
		albumIdValue := v.Get("album")

		id, err := strconv.Atoi(idValue)
		if err != nil {
			JSONMarshalBadRequest(err.Error(), 500).WriteToResponse(w)
			return
		}

		albumId, err := strconv.Atoi(albumIdValue)
		if err != nil {
			JSONMarshalBadRequest(err.Error(), 500).WriteToResponse(w)
			return
		}

		photo, err := fetchPhotoByIdAndAlbumId(id, albumId)
		if err != nil {
			JSONMarshalBadRequest(err.Error(), 500).WriteToResponse(w)
			return
		}

		json.NewEncoder(w).Encode(photo)
		return

	}

	JSONMarshalBadRequest("Invalid Type, Please provide correct type", 400).WriteToResponse(w)
	return
}

func fetchPhotoByIdAndAlbumId(id int, albumId int) (Photo, error) {

	var photo Photo

	database, err := util.Connection()
	if err != nil {
		fmt.Println("service:fetchPhotoByIdAndAlbumId, unable to get connection ", err.Error())
		return photo, err
	}

	err = database.QueryRow("select id, albumId, photoId, title, url, thumbnailUrl from `avinash.verma1983@gmail.com`.photo where id = ? and albumId = ?", id, albumId).Scan(&photo.Id, &photo.AlbumId, &photo.PhotoId, &photo.Url, &photo.ThumbnailUrl)
	if err != nil {
		fmt.Println("service:fetchPhotoByIdAndAlbumId, unable to fetch record ", err.Error())
		return photo, err
	}
	defer database.Close()
	return photo, nil

}

func fetchAlbumById(id int) (Album, error) {

	var album Album

	database, err := util.Connection()
	if err != nil {
		fmt.Println("service:fetchAlbumById, unable to get connection ", err.Error())
		return album, err
	}

	err = database.QueryRow("SELECT id, userId, title FROM `avinash.verma1983@gmail.com`.album where id = ?", id).Scan(&album.Id, &album.UserId, &album.Title)
	if err != nil {
		fmt.Println("service:fetchAlbumById, unable to fetch record ", err.Error())
		return album, err
	}
	defer database.Close()
	return album, nil
}
