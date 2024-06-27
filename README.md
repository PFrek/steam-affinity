# Steam Affinity

[Live Demo](https://steam-affinity-frontend-edfp5mfira-uc.a.run.app/)

## What is Steam Affinity?

Steam Affinity is a website project that allows comparing a user's Steam game library with those of their Steam friends.

The user's friends are ranked according to the similarity of their game libraries. Users with *similarly sized libraries* and with *many games in common* are ranked higher.

You can also see the games you both own, or the games that are missing from your library that the other user owns.

![Usage demo](https://github.com/PFrek/steam-affinity/blob/main/gif/demo.gif)

## Project structure

The backend code can be found in the project's root folder. Any frontend code will be found in the frontend/steam-affinity folder.

## Installing / Contributing

After cloning the repository you can run the backend by building and running the Go application, or by using the provided run.sh script.

```bash
go build -o server && ./server
```

If you are running this locally you will need to set up a .env file that includes the following information:

```env
# Optional Port to run the backend server on. If not provided defaults to 8080.
PORT=8080 

# Steam API Key obtainable from Steam for free. This is mandatory when running the app locally.
STEAM_APIKEY={...} 

# Optional number of *minutes* before an entry in the Friends Cache should be renewed. Defaults to 5 minutes.
FRIENDS_CACHE_RENEW=5 

# Optional number of *minutes* before an entry in the Summaries Cache should be renewed. Defaults to 1440 minutes (24 hours).
SUMMARIES_CACHE_RENEW=1440 

# Optional number of *minutes* before an entry in the Games Cache should be renewed. Defaults to 30 minutes.
GAMES_CACHE_RENEW=30 
```

---

The frontend can be run like so:

```bash
# Move to the frontend folder
cd frontend/steam-affinity

# Run it
npm run dev
```

## API Endpoints

The backend API exposes the following endpoints:

### GetFriends

Returns a user's friends list.

**Endpoint:** GET /api/v1/friends?steamid={id}

**Response:**

```json
{
  "players": [
    {
			"steamid": "1234",
			"communityvisibilitystate": 3,
			"personaname": "ExampleUser",
			"avatar": "https://avatars.steamstatic.com/example.jpg",
			"avatarmedium": "https://avatars.steamstatic.com/example_medium.jpg",
			"avatarfull": "https://avatars.steamstatic.com/example_full.jpg"
		}
  ]
}
```

### GetOwnedGames

Returns a list of games owned by the user.

**Endpoint:** GET /api/v1/ownedGames?steamid={id}

**Response:**

```json
{
  "SteamID": "1234",
	"game_count": 1,
	"games": [
		{
			"appid": 70,
			"name": "Half-Life",
			"img_icon_url": "95be6d131fc61f145797317ca437c9765f24b41c"
		}
  ]
}
```

### GetComparison

Returns the affinity comparison between two users' libraries.

**Endpoint:** GET /api/v1/ownedGames/compare?player1={id1}&player2={id2}&listGames=true

**Response:**

```json
{
	"affinity": 41.38402685278917,
	"similarity": 0.12408759124087591,
	"weight": 333.5065693430657,
	"player1ID": "1234",
	"player1Ratio": 0.2961672473867596,
	"player2ID": "5678",
	"player2Ratio": 0.2135678391959799,
	"matches": 85,
	"matching_games": [
		{
			"appid": 4000,
			"name": "Garry's Mod",
			"img_icon_url": "4a6f25cfa2426445d0d9d6e233408de4d371ce8b"
		}
  ]
}
```

### GetAffinityRanking

Runs GetComparison for each friend in the user's friends list, and returns said list ordered by affinity.

**Endpoint:** GET /api/v1/friends/ranking?steamid={id}&listGames=true

**Response:**

```json
{
  "ranking": [
    {
    	"affinity": 41.38402685278917,
    	"similarity": 0.12408759124087591,
    	"weight": 333.5065693430657,
    	"player1ID": "1234",
    	"player1Ratio": 0.2961672473867596,
    	"player2ID": "5678",
    	"player2Ratio": 0.2135678391959799,
    	"matches": 85,
    	"matching_games": [
    		{
    			"appid": 4000,
    			"name": "Garry's Mod",
    			"img_icon_url": "4a6f25cfa2426445d0d9d6e233408de4d371ce8b"
    		}
      ]
    }
  ]
}
```
