package dto

type ErrorBadRequestResponse struct { //400
	Error string `json:"error" example:"Bad Request"`
}

type ErrorInternalServerResponse struct { //500
	Error string `json:"error" example:"Internal Server Error"`
}

type ErrorNotFoundResponse struct { //404
	Error string `json:"error" example:"Song not found"`
}

type ResponseOK struct { //200
	Success string `json:"success" example:"true"`
}

type PlaylistResponseDto struct {
	Total int64             `json:"total" example:"1"`
	Page  int               `json:"page" example:"3"`
	Limit int               `json:"limit" example:"15"`
	Songs []ResponseSongDTO `json:"data"`
}

type ResponseSongDTO struct {
	Group       string `json:"group"  example:"Young Rus"`
	Song        string `json:"song"  example:"NeonBones"`
	Text        string `json:"text"  example:"I'm a sad guy, a humid day\n the humidity of my line"`
	Link        string `json:"link"  example:"example.com"`
	ReleaseDate string `json:"releaseDate"  example:"16.07.2006"`
}
