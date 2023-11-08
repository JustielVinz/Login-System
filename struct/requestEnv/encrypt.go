package requestenv

type Env struct {
	DatabaseName     string `json:"database_name"`
	PostgresUsename  string `json:"postgres_username"`
	PostgresPassword string `json:"postgres_password"`
	PostgresHost     string `json:"postgres_host"`
	PostgresPort     string `json:"postgres_port"`
}
