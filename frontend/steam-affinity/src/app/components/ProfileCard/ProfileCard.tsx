import Image from "next/image"
import { getPlayerSummaries } from "../../actions/actions"
import { AffinityBoundaries, PlayerAffinity, PlayerSummary } from "../../definitions/types"
import styles from "./ProfileCard.module.css"
import AffinityInfo from "./AffinityInfo/AffinityInfo"
import { Color, getColorForValue } from "@/app/helpers/colorLerp"

export default async function ProfileCard({ summary, affinity, affinityBoundaries }
  : {
    summary: PlayerSummary,
    affinity: PlayerAffinity | null,
    affinityBoundaries: AffinityBoundaries,
  }) {
  let player: PlayerSummary = summary;

  let lowColor: Color = {
    r: 255,
    g: 0,
    b: 0,
  };

  let highColor: Color = {
    r: 0,
    g: 255,
    b: 0,
  }

  let affinityColor: Color = highColor;
  if (affinity != null) {
    affinityColor = getColorForValue(affinity?.affinity, affinityBoundaries.min, affinityBoundaries.max, lowColor, highColor);
  }

  return (
    <div className={styles.container}>
      <div className={styles.profileSection}>
        <Image className={styles.avatar} src={player.avatarmedium} width={50} height={50} alt="Profile picture" />
        <p className={styles.personaname}>{player.personaname}</p>
      </div>

      {affinity != null ?
        <AffinityInfo affinity={affinity} color={affinityColor} />
        :
        <></>}
    </div >
  )
}
