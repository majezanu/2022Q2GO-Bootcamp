package model

type Pokemon struct {
	Id   int    `json:"id"       example:"1"`
	Name string `json:"name"       example:"Pikachu"`
}

type MultipleFilter struct {
	IdType         string `query:"type"`
	Items          int    `query:"items"`
	ItemsPerWorker int    `query:"items_per_worker"`
}
