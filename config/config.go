package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("falha ao carregar o arquivo .env: %v", err)
	}else{
		fmt.Println("Arquivo .env carregado com sucesso")
	}

	return nil
}
