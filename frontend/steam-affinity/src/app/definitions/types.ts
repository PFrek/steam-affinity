export interface PlayerSummary {
	steamid: string;
	communityvisibilitystate: number;
	personaname: string;
	avatar: string;
	avatarmedium: string;
	avatarfull: string;
}

export interface Game {
	appid: number;
	name: string;
	img_icon_url: string;
}

export interface PlayerAffinity {
	affinity: number;
	similarity: number;
	weight: number;
	player1ID: string;
	player1Ratio: number;
	player2ID: string;
	player2Ratio: number;
	player2GamesCount: number;
	matches: number;
	matching_games: Game[];
}

export interface AffinityBoundaries {
	min: number;
	max: number;
}
