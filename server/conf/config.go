package conf

type ServerConfig struct {
	Name        string
	Host        string
	Port        int
	MySQLInfo   MySQLConfig `mapstructure:"mysql"`
	LogsAddress string
}

type MySQLConfig struct {
	Host     string
	Port     int
	Name     string
	Password string
	DBName   string
}
