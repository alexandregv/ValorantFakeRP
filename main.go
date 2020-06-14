package main

import (
	"sync"
	"time"

	discord "github.com/ananagame/rich-go/client"
)

var (
	APP_NAME  string = "ValorantRP"
	VERSION   string = "0.1.0"
	CLIENT_ID string = "697136055719165963"
)

func main() {
	err := discord.Login(CLIENT_ID)
	if err != nil {
		panic(err)
	}

	start, _ := time.Parse(time.RFC822, "07 Apr 20 19:00 UTC")
	//start := time.Now()

	err = discord.SetActivity(discord.Activity{
		Details:    "Closed Beta",
		State:      "Playing Solo | Jett",
		LargeImage: "v_logo",
		LargeText:  "Valorant (Closed Beta)",
		//SmallImage: "v_logo",
		//SmallText:  "Valorant (Closed Beta)",
		Party: &discord.Party{
			ID:         "abc123",
			Players:    1,
			MaxPlayers: 5,
		},
		Secrets: &discord.Secrets{
			//Match:    "abc123",
			Join:     "def456",
			Spectate: "ghi789",
		},
		Timestamps: &discord.Timestamps{
			Start: &start,
		},
	})

	if err != nil {
		panic(err)
	}

	//fmt.Printf("[%s]:[%s] %s est lancé en arrière plan.\n", APP_NAME, VERSION, APP_NAME)

	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}
