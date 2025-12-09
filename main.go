package main

import (
	"log"
	"os"

	"flag"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var session *discordgo.Session

func init() { flag.Parse() }

func init() {
	_ = godotenv.Load() // best-effort .env load
	var err error
	flag.Set("token", os.Getenv("DISCORD_TOKEN")) 
	session, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	log.Println(*BotToken, "hihi")
}

func main() {
	session.Open()

	session.AddHandler(onMessage)

	defer session.Close()

	log.Println("Bot is running...")
	select {} // keep alive
}
