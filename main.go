package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/isksss/twitchNotify/commands"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		fmt.Println("DISCORD_BOT_TOKEN is not set")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dg.AddHandler(commands.HandleInteraction)

	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()

	// コマンド登録（ギルド指定で即反映、グローバルにするなら ""）
	appID := dg.State.User.ID
	guildID := "" // or "YOUR_GUILD_ID"
	for _, cmd := range commands.All {
		_, err := dg.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
