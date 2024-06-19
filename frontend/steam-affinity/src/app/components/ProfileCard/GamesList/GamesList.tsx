import { Game } from "@/app/definitions/types";
import styles from "./GamesList.module.css";
import GameIcon from "./GameIcon/GameIcon";


export default function GamesList({ games }: { games: Game[] }) {
  const iconWidth = 50;
  const iconHeight = 50;

  games.sort((a: Game, b: Game): number => {
    if (a.name < b.name) {
      return -1;
    } else if (a.name > b.name) {
      return 1;
    }

    return 0;
  })

  return (
    <div className={styles.container}>
      <h3>Matching Games</h3>
      <div className={styles.list}>
        {games.map((game) => {
          return (
            <a className={styles.game} href={`https://store.steampowered.com/app/${game.appid}/`} target="_blank" rel="noopener">
              <GameIcon game={game} width={iconWidth} height={iconHeight} />
              <p>{game.name}</p>
            </a>
          )
        })}
      </div>
    </div>
  )
}
