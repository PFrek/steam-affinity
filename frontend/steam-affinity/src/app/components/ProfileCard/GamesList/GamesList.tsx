import { Game } from "@/app/definitions/types";
import styles from "./GamesList.module.css";
import GameIcon from "./GameIcon/GameIcon";
import { MouseEventHandler } from "react";


export default function GamesList({ games, listType, setMatchingGamesDisplay, setMissingGamesDisplay }:
  {
    games: Game[] | null,
    listType: "matching" | "missing",
    setMatchingGamesDisplay: MouseEventHandler<HTMLButtonElement>,
    setMissingGamesDisplay: MouseEventHandler<HTMLButtonElement>,
  }) {

  const iconWidth = 50;
  const iconHeight = 50;

  if (games != null) {
    games?.sort((a: Game, b: Game): number => {
      if (a.name < b.name) {
        return -1;
      } else if (a.name > b.name) {
        return 1;
      }

      return 0;
    })
  }

  return (
    <div className={`${styles.container} ${listType === "missing" ? styles.missingList : ""}`}>
      <div className={styles.buttonsContainer}>
        <button
          onClick={setMatchingGamesDisplay}
          className={`${styles.button} ${listType === "matching" ? styles.active : ""}`}
          disabled={listType === "matching"}>
          <h3>Matching Games</h3>
        </button>
        <button
          onClick={setMissingGamesDisplay}
          className={`${styles.button} ${listType === "missing" ? styles.activeMissing : ""}`}
          disabled={listType === "missing"}>
          <h3>Missing Games</h3>
        </button>
      </div>
      <div className={styles.list}>
        {games != null && games.length > 0 && games.map((game) => {
          return (
            <a
              className={`${styles.game} ${listType === "missing" ? styles.missingEntry : ""}`}
              href={`https://store.steampowered.com/app/${game.appid}/`}
              target="_blank"
              rel="noopener"
              key={game.appid}
              onClick={(ev) => { ev.stopPropagation() }}
            >
              <GameIcon game={game} width={iconWidth} height={iconHeight} />
              <p>{game.name}</p>
            </a>
          )
        })}
      </div>
    </div>
  )
}
