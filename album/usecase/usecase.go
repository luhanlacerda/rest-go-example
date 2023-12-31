package usecase

import model "example/rest-go-example/album/model"

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []model.Album {
	return albums
}

func SaveNewAlbum(album model.Album) {
	albums = append(albums, album)
}
