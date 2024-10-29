package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 구조체 정의
type Config struct {
	ServerPort   string
	DatabasePath string
}

// 전역 Config 변수
var AppConfig *Config

// InitConfig 함수: 환경 변수를 불러와 설정 초기화
func InitConfig() *Config {
	// .env 파일 불러오기
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 파일을 불러오지 못했습니다. 기본 설정을 사용합니다.")
	}

	// 환경 변수에서 설정 값 가져오기
	AppConfig = &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DatabasePath: getEnv("DATABASE_PATH", "test.db"),
	}

	return AppConfig
}

// 환경 변수를 불러오고 기본값 설정
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
