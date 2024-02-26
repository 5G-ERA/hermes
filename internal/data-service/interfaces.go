package data_service

type DataFetcher interface {
	Fetch(dir string) error
}
type DataPoster interface {
	Post(dir string) error
}
