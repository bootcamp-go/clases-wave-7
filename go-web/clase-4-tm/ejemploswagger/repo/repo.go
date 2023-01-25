package repo

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// base de datos del repositorio
var albums = []Album{
	{ID: "1", Title: "cualquiera", Artist: "cualquiera", Year: 2002},
}

func GetAlbums() []Album {
	return albums
}

func Create(album Album) []Album{
	albums = append(albums, album)
	return albums
}