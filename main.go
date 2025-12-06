package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    token := os.Getenv("DISCORD_TOKEN")
    fmt.Println("Token length:", len(token)) // debug

    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("error creating Discord session:", err)
        return
    }

    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection:", err)
        return
    }

    fmt.Println("Bot is running")
    select {}
}
