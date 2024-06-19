"use client";

import Image from "next/image"
import { AffinityBoundaries, PlayerAffinity, PlayerSummary } from "../../definitions/types"
import styles from "./ProfileCard.module.css"
import AffinityInfo from "./AffinityInfo/AffinityInfo"
import { Color, getColorForValue } from "@/app/helpers/colorLerp"
import { useEffect, useState } from "react";
import GamesList from "./GamesList/GamesList";

export default function ProfileCard({ summary, affinity, affinityBoundaries }
  : {
    summary: PlayerSummary,
    affinity: PlayerAffinity | null,
    affinityBoundaries: AffinityBoundaries,
  }) {

  const [expanded, setExpanded] = useState<boolean>(false);

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

  const toggleExpand = () => {
    if (affinity && affinity?.player2GamesCount > 0) {
      setExpanded(!expanded)
    }
  }

  return (
    <div className={styles.container} onClick={toggleExpand}>
      <div className={styles.header}>
        <div className={styles.profileSection}>
          <Image className={styles.avatar} src={summary.avatarmedium} width={50} height={50} alt="Profile picture" />
          <p className={styles.personaname}>{summary.personaname}</p>
        </div>

        {affinity == null || affinity.player2GamesCount == 0 ? <></>
          :
          <AffinityInfo affinity={affinity} color={affinityColor} />}

      </div>
      {expanded && affinity && <GamesList games={affinity?.matching_games} />}
    </div >
  )
}
