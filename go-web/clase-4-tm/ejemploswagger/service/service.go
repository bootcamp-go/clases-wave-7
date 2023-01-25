package service

import (
	"ejemploswagger/repo"
)

func GetAlbums() []repo.Album {
	return repo.GetAlbums()
}

func Create(album repo.Album) []repo.Album{
	return repo.Create(album)
}