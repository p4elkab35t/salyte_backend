package config

// "github.com/spf13/viper"

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

// func LoadConfig() (*Config, error) {
//     viper.SetConfigName("config")
//     viper.SetConfigType("yaml")
//     viper.AddConfigPath(".")
//     if err := viper.ReadInConfig(); err != nil {
//         return nil, err
//     }

//     config := &Config{
//         DBHost:     viper.GetString("database.host"),
//         DBPort:     viper.GetInt("database.port"),
//         DBUser:     viper.GetString("database.user"),
//         DBPassword: viper.GetString("database.password"),
//         DBName:     viper.GetString("database.name"),
//     }
//     return config, nil
// }
