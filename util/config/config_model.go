package config

type Config struct {
	API_PORT string `mapstructure:"API_PORT"`

	MYSQL_DRIVER_NAME string `mapstructure:"MYSQL_DRIVER_NAME"`
	MYSQL_USERNAME    string `mapstructure:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_HOSTNAME    string `mapstructure:"MYSQL_HOSTNAME"`
	MYSQL_DB_NAME     string `mapstructure:"MYSQL_DB_NAME"`

	REDIS_ADDRESS  string `mapstructure:"REDIS_ADDRESS"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_DB_NUM   string `mapstructure:"REDIS_DB_NUM"`
	SKIP_UNIT_TEST string `mapstructure:"SKIP_UNIT_TEST"`
}
