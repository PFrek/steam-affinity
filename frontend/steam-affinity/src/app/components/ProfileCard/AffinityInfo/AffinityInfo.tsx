import { PlayerAffinity } from "@/app/definitions/types"
import styles from "./AffinityInfo.module.css"
import { Color, colorToStyle } from "@/app/helpers/colorLerp";

export default function AffinityInfo({ affinity, color }: { affinity: PlayerAffinity, color: Color }) {
  const ratioOwned = affinity.player2Ratio;

  return (
    <div className={styles.container}>
      <p className={styles.ratioOwned} >You own <span className={styles.percent}>{`${(ratioOwned * 100.0).toFixed(2)}%`}</span> of their games</p>
      <p className={styles.affinity} style={{ color: colorToStyle(color) }}>
        <span className={styles.affinityText}>Affinity: </span> {affinity.affinity.toFixed(2)}
      </p>
    </div>
  )
}
