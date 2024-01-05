package config

import "github.com/caarlos0/env/v10"

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func InitConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
