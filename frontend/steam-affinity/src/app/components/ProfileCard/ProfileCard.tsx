import Image from "next/image"
import { getPlayerSummaries } from "../../actions/actions"
import { PlayerSummary } from "../../definitions/types"
import styles from "./ProfileCard.module.css"

export default async function ProfileCard({ steamID }: { steamID: string }) {
  const playerSummaries = await getPlayerSummaries(steamID)

  let player: PlayerSummary = {
    steamid: "unknownID",
    communityvisibilityState: 0,
    personaname: "unknownUsername",
    avatar: "unknown_user.png",
    avatarmedium: "unknown_user.png",
    avatarfull: "unknown_user.png"
  }

  if (playerSummaries.length > 0) {
    player = playerSummaries[0]
  }

  return (
    <div className={styles.container}>
      <Image className={styles.avatar} src={player.avatarmedium} width={50} height={50} alt="Profile picture" />
      <h2 className={styles.personaname}>{player.personaname}</h2>
    </div>
  )
}
