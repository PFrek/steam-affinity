export interface PlayerSummary {
	steamid: string;
	communityvisibilitystate: number;
	personaname: string;
	avatar: string;
	avatarmedium: string;
	avatarfull: string;
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
}

export interface AffinityBoundaries {
	min: number;
	max: number;
}
