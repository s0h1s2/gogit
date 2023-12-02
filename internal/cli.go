package internal

type InitCmd struct{}

type HashObjectCmd struct {
	FilePath string `arg:"positional" placeholder:"FILE"`
}
type HashCatCmd struct {
	Hash string `arg:"positional" placeholder:"HASH"`
}
type WriteTreeCmd struct{}
