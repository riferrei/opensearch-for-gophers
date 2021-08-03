package domain

/*************************************************/
/*********** Internal for the Context ************/
/*************************************************/

type contextKey struct {
	Key int
}

var MoviesKey contextKey = contextKey{Key: 1}
var ClientKey contextKey = contextKey{Key: 2}

/*************************************************/
/*********** Types for the Application ***********/
/*************************************************/

type MovieRaw struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Info  *InfoRaw `json:"info"`
}

type InfoRaw struct {
	RunningTime float32  `json:"running_time_secs"`
	ReleaseDate string   `json:"release_date"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	Directors   []string `json:"directors"`
}

type Movie struct {
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	RunningTime float32  `json:"runningTime"`
	ReleaseDate string   `json:"releaseDate"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	Directors   []string `json:"directors"`
}

/*************************************************/
/************ Types for the Livestream ***********/
/*************************************************/
