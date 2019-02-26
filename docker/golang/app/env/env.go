package env

import (
  "os"
  "fmt"
  "log"
  "github.com/joho/godotenv"
)

func Env_load() {
  if os.Getenv("GO_ENV") == "" {
    os.Setenv("GO_ENV", "develop")
  }

  err := godotenv.Load(fmt.Sprintf("./envs/.env.%s", os.Getenv("GO_ENV")))
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  err = godotenv.Load("./envs/version.env")
  if err != nil {
    log.Fatal("Error loading version.env file")
  }
}