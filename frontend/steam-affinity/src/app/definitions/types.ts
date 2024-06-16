export interface PlayerSummary {
	steamid: string;
	communityvisibilityState: number;
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
	matches: number;
}
