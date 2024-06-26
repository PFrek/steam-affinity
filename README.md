# Steam Affinity

## What is Steam Affinity?

Steam Affinity is a website project that allows comparing a user's Steam game library with those of their Steam friends.

The user's friends are ranked according to the similarity of their game libraries. Users with *similarly sized libraries* and with *many games in common* are ranked higher.

You can also see the games you both own, or the games that are missing from your library that the other user owns.

![Usage demo](https://github.com/PFrek/steam-affinity/blob/main/gif/demo.gif)

## Project structure

The backend code can be found in the project's root folder. Any frontend code will be found in the frontend/steam-affinity folder.

## Installation

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
