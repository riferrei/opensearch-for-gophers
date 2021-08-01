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
	Title string  `json:"title"`
	Year  int     `json:"year"`
	Info  InfoRaw `json:"info"`
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
	Title       string
	Year        int
	RunningTime float32
	ReleaseDate string
	Rating      float32
	Genres      []string
	Actors      []string
	Directors   []string
}

/*************************************************/
/************* Queries and Searches **************/
/*************************************************/

type GetResponse struct {
	Index   string `json:"_index"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Source  Movie  `json:"_source"`
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*SearchHit
	}
}

type SearchHit struct {
	Source MovieRaw `json:"_source"`
}

/*************************************************/
/************* Aggregation Example ***************/
/*************************************************/

type AggregationRequest struct {
	Size int  `json:"size"`
	Aggs Aggs `json:"aggs"`
}

type Aggs struct {
	MovieCountPerGenre MovieCountPerGenreRequest `json:"movieCountPerGenre"`
}

type MovieCountPerGenreRequest struct {
	Terms Terms `json:"terms"`
}

type Terms struct {
	Field string `json:"field"`
	Size  int    `json:"size"`
}

type AggregationResponse struct {
	Aggregations Aggregations `json:"aggregations"`
}

type Aggregations struct {
	MovieCountPerGenreResponse MovieCountPerGenreResponse `json:"movieCountPerGenre"`
}

type MovieCountPerGenreResponse struct {
	Buckets []Bucket `json:"buckets"`
}

type Bucket struct {
	Key           string `json:"key"`
	DocumentCount int    `json:"doc_count"`
}
