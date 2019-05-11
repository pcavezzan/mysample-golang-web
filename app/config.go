package app

type Config struct {
	Port int
	Private bool
	DataSource
}

type DataSource struct {
	User string
	Password string
	Host string
	Database string
	Port int
}