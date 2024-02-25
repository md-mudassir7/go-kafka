package config

type Config struct {
	Kafka struct {
		Host  string `json:"host"`
		Port  string `json:"port"`
		Topic string `json:"topic"`
	} `json:"kafka"`
}

func LoadConfig() Config {
	config := Config{
		Kafka: struct {
			Host  string "json:\"host\""
			Port  string "json:\"port\""
			Topic string "json:\"topic\""
		}{
			Host:  "localhost",
			Port:  "9092",
			Topic: "common",
		},
	}
	return config
}
