package configs

type AppConfig struct {
	Name          string `mapstructure:"APPLICATION_NAME"`
	ServerPort    string `mapstructure:"APPLICATION_SERVER_PORT"`
	ServerTimeout string `mapstructure:"APPLICATION_SERVER_TIMEOUT"`
	MySQLHostDsn  string `mapstructure:"MYSQL_HOST"`
}

func LoadConfig() (*AppConfig, error) {
	var appConfig AppConfig

	viper := NewViperConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}
