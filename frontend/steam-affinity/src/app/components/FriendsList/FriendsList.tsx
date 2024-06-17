import { getAffinityRanking, getPlayerSummaries } from "@/app/actions/actions"
import styles from "./FriendsList.module.css"
import ProfileCard from "../ProfileCard/ProfileCard";
import { AffinityBoundaries, PlayerSummary } from "@/app/definitions/types";


export default async function FriendsList({ steamid }: { steamid: string }) {
	const affinities = await getAffinityRanking(steamid);


	let boundaries: AffinityBoundaries = {
		min: Number.MAX_VALUE,
		max: Number.MIN_VALUE,
	};

	let steamids: string[] = [];

	affinities.forEach((entry) => {
		steamids.push(entry.player2ID);

		if (entry.affinity > boundaries.max) {
			boundaries.max = entry.affinity;
		}

		if (entry.affinity < boundaries.min) {
			boundaries.min = entry.affinity;
		}
	})

	interface IPlayerHash {
		[id: string]: PlayerSummary;
	}
	const players: IPlayerHash = {};
	const summaries = await getPlayerSummaries(steamids.join(","));

	summaries.forEach((summary) => {
		players[summary.steamid] = summary;
	})

	return (
		<div className={styles.container}>
			<h2 className={styles.title}>Friends</h2>
			{affinities.length == 0 ?
				<p>No friends found</p>
				: <ul>
					{affinities.map((entry) => {
						return <li key={entry.player2ID}>
							<ProfileCard
								affinity={entry}
								summary={players[entry.player2ID]}
								affinityBoundaries={boundaries} />
						</li>
					})}
				</ul>

			}
		</div>
	)
}
