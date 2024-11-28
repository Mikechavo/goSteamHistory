package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// SteamGame represents a single game's data
type SteamGame struct {
	Name     string
	Playtime float64
	IconURL  string
}

// FetchSteamData fetches the user's games from the Steam API
func FetchSteamData(apiKey, steamID string) ([]SteamGame, error) {
	url := fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=true", apiKey, steamID)
	client := resty.New()

	var response struct {
		Response struct {
			Games []struct {
				AppID      int    `json:"appid"`
				Name       string `json:"name"`
				Playtime   int    `json:"playtime_forever"`
				ImgIconURL string `json:"img_icon_url"`
			} `json:"games"`
		} `json:"response"`
	}

	// Make the API request
	_, err := client.R().SetResult(&response).Get(url)
	if err != nil {
		log.Printf("Error fetching Steam data: %v", err)
		return nil, err
	}

	// Process the games and return them
	var games []SteamGame
	for _, game := range response.Response.Games {
		games = append(games, SteamGame{
			Name:     game.Name,
			Playtime: float64(game.Playtime) / 60, // Convert minutes to hours
			IconURL:  fmt.Sprintf("https://media.steampowered.com/steamcommunity/public/images/apps/%d/%s.jpg", game.AppID, game.ImgIconURL),
		})
	}

	// Sort games by playtime in descending order
	sort.Slice(games, func(i, j int) bool {
		return games[i].Playtime > games[j].Playtime
	})

	return games, nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API Key and Steam ID from environment variables
	apiKey := os.Getenv("STEAM_API_KEY")
	steamID := os.Getenv("STEAM_ID")

	// Ensure the environment variables are set
	if apiKey == "" || steamID == "" {
		log.Fatal("STEAM_API_KEY and STEAM_ID must be set in environment variables")
	}

	// Initialize Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates from the templates directory
	r.LoadHTMLGlob("templates/*.html")

	// Define the route to fetch and display game data
	r.GET("/", func(c *gin.Context) {
		// Fetch games data from Steam API
		games, err := FetchSteamData(apiKey, steamID)
		if err != nil {
			// Handle error with a JSON response
			log.Printf("Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Calculate total playtime and total number of games
		var totalPlaytime float64
		for _, game := range games {
			totalPlaytime += game.Playtime
		}

		// Prepare data to pass to the template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"games":         games,
			"totalPlaytime": fmt.Sprintf("%.2f", totalPlaytime),
			"totalGames":    len(games),
		})
	})

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
