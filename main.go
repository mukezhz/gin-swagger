package main

import (
    "github.com/mukezhz/gin_swag/bootstrap"

    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    _ = bootstrap.RootApp.Execute()
}
