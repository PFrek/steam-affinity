"use server";

import { PlayerSummary } from "../definitions/types";
import { baseURL } from "../definitions/urls";

export async function getPlayerSummaries(steamIDs: string): Promise<PlayerSummary[]> {
	const url = new URL("summaries", baseURL);
	url.searchParams.set("steamids", steamIDs);

	let summaries: PlayerSummary[] = [];

	try {
		let resp = await fetch(url, {
			method: "GET",
		});

		if (resp.status >= 400) {
			throw new Error(`Failed http request with status ${resp.status}`)
		}

		let json = await resp.json();

		console.log(json);
		summaries = json.players
		return summaries;
	} catch (e) {
		console.log(`Failed to get player summaries: ${e}`);

		return summaries;
	}
}

