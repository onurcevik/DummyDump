package model

type Config struct {
	Source     SOURCE_TYPE
	Import     bool
	Export     bool
	User       string
	Password   string
	Path       string
	DB         string
	BinaryPath string
}
