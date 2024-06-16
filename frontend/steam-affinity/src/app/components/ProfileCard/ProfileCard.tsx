import Image from "next/image"
import { getPlayerSummaries } from "../../actions/actions"
import { PlayerAffinity, PlayerSummary } from "../../definitions/types"
import styles from "./ProfileCard.module.css"
import AffinityInfo from "./AffinityInfo/AffinityInfo"

export default async function ProfileCard({ steamid, affinity }: { steamid: string, affinity: PlayerAffinity | null }) {
  let player: PlayerSummary = {
    steamid: "unknownID",
    communityvisibilityState: 0,
    personaname: "unknownUsername",
    avatar: "unknown_user.png",
    avatarmedium: "unknown_user.png",
    avatarfull: "unknown_user.png"
  }

  let playerSummaries = await getPlayerSummaries(steamid);


  if (playerSummaries.length > 0) {
    player = playerSummaries[0]
  }

  return (
    <div className={styles.container}>
      <div className={styles.profileSection}>
        <Image className={styles.avatar} src={player.avatarmedium} width={50} height={50} alt="Profile picture" />
        <p className={styles.personaname}>{player.personaname}</p>
      </div>

      {affinity != null ?
        <AffinityInfo affinity={affinity} />
        :
        <></>}
    </div >
  )
}
