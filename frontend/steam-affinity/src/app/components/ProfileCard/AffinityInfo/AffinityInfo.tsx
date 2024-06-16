import { PlayerAffinity } from "@/app/definitions/types"
import styles from "./AffinityInfo.module.css"

export default function AffinityInfo({ affinity }: { affinity: PlayerAffinity }) {
  const ratioOwned = affinity.player2Ratio;

  return (
    <div className={styles.container}>
      <p className={styles.ratioOwned} >You own <span className={styles.percent}>{`${(ratioOwned * 100.0).toFixed(2)}%`}</span> of their games</p>
      <p className={styles.affinity}>{affinity.affinity.toFixed(2)}</p>
    </div>
  )
}
