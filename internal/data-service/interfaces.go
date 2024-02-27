package data_service

type DataFetcher interface {
	Fetch(netAppKey, targetDir string) error
}
type DataPoster interface {
	Post(netAppKey, sourceDir string) error
}
