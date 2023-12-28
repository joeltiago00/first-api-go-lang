package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func Env(key string) string {
	err := godotenv.Load(fmt.Sprintf("%s/.env", rootDir()))

	if err != nil && os.Getenv("APP_ENV") != "testing" {
		log.Fatalf("Error loading .env file...")
	}

	return os.Getenv(key)
}

func rootDir() string {
	//TODO:: change vars name to make sense on reading...
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
