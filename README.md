Steam Games App - Fullstack with Golang
This is a simple fullstack application built with Golang that integrates with the Steam API to fetch and display information about your Steam games. It uses the Gin framework for the backend and serves dynamic data via API requests. The frontend is built with HTML and CSS for a clean, responsive layout.

Features
Displays a list of Steam games with the following details:
Game Name
Playtime (in hours)
Game Icon
Total Hours Played and Number of Games are shown at the top in a separate card.
The interface is mobile-friendly, with a grid layout that adapts to different screen sizes.
Built with Golang, Gin Framework, and Steam API.
Getting Started
To get this app up and running on your local machine, follow these steps:

Prerequisites
Golang (version 1.18 or higher)
Git (for cloning the repository)
A Steam API Key and Steam ID (instructions below on how to get them)
Installation
Clone the repository:
bash
Copy code
git clone https://github.com/your-username/steam-games-app.git
cd steam-games-app
Install the dependencies:
Gin Framework: Used for routing and handling HTTP requests.
bash
Copy code
go get github.com/gin-gonic/gin
Resty: Used to make HTTP requests to the Steam API.
bash
Copy code
go get github.com/go-resty/resty/v2
Set up environment variables:
You will need your Steam API key and Steam ID. You can generate the API key by visiting Steam Community.
Set your STEAM_API_KEY and STEAM_ID in the environment variables.
For example, on Linux/Mac:

bash
Copy code
export STEAM_API_KEY="your_steam_api_key"
export STEAM_ID="your_steam_id"
On Windows, you can set environment variables in the system settings or use the following in the terminal:

bash
Copy code
set STEAM_API_KEY=your_steam_api_key
set STEAM_ID=your_steam_id
Run the application:
bash
Copy code
go run main.go
This will start the server on http://localhost:8080.

Viewing the App
Once the app is running, you can view it by opening your browser and going to http://localhost:8080. You will see a list of your Steam games, along with their playtime and icons. At the top of the page, you’ll also see the total playtime and the number of games you have.

How It Works
The backend is built using Golang and the Gin framework. It handles the HTTP requests and interacts with the Steam API to fetch game data.
The Steam API is queried using the Resty HTTP client. The API returns the list of games, including their names, playtimes, and icons.
The frontend is rendered using HTML templates (with dynamic data passed from the backend). The game information is displayed using cards in a grid layout that is responsive on mobile devices.
The total playtime and number of games are calculated on the backend and displayed at the top of the page.
Contributing
If you want to contribute to this project, feel free to open a pull request! Here are some ideas for contributions:

Improve the user interface with more advanced CSS or a frontend framework like React.
Add error handling for invalid API keys or Steam IDs.
Implement additional features such as sorting games by playtime or adding filters.
License
This project is open-source and available under the MIT License.

Acknowledgments
Steam API for providing the game data.
Gin Framework for building the backend.
Resty for making HTTP requests easier.
