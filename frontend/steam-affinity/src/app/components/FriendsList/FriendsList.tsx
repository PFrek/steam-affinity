import { getAffinityRanking } from "@/app/actions/actions"
import styles from "./FriendsList.module.css"
import ProfileCard from "../ProfileCard/ProfileCard";


export default async function FriendsList({ steamid }: { steamid: string }) {
	const affinities = await getAffinityRanking(steamid);

	return (
		<div className={styles.container}>
			<h2 className={styles.title}>Friends</h2>
			{affinities.length == 0 ?
				<p>No friends found</p>
				: <ul>
					{affinities.map((entry) => {
						return <li key={entry.player2ID}>
							<ProfileCard affinity={entry} steamid={entry.player2ID} />
						</li>
					})}
				</ul>

			}
		</div>
	)
}
