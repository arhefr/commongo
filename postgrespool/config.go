package postgrespool

type Config struct {
	User         string
	Password     string
	Host         string
	Port         string
	Database     string
	MaxAttemps   int64
	DelayAttemps int64
}
