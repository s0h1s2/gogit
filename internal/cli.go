package internal

type InitCmd struct{}

type HashObjectCmd struct {
	FilePath string `arg:"positional" placeholder:"FILE"`
}
type HashCatCmd struct {
	Type string `arg:"positional"`
	Hash string `arg:"positional"`
}
type WriteTreeCmd struct{}
type ReadTreeCmd struct{}
