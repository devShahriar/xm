package config

var AppConfig *Config

func GetAppConfig() *Config {
	return AppConfig
}

type Config struct {
	DbConfig     DbConfig `yaml:"db_config"`
	JWTSecretKey string   `yaml:"jwt_secret_key"`
}

type DbConfig struct {
	Host               string `yaml:"host"`
	Password           string `yaml:"password"`
	User               string `yaml:"user"`
	DbName             string `yaml:"db_name"`
	Port               string `yaml:"port"`
	SlowQueryThreshold int    `yaml:"slow_query_threshold"`
}
