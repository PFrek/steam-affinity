import { Game } from "@/app/definitions/types";
import Image from "next/image";
import { useState } from "react";
import styles from "./GameIcon.module.css";

const imageLoader = ({ src }: { src: string }) => {
  return `https://media.steampowered.com/steamcommunity/public/images/apps/${src}`;
}

export default function GameIcon({ game, width, height }: { game: Game, width: number, height: number }) {
  const [failedLoad, setFailedLoad] = useState<boolean>(false);

  const blurDataURL = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mOsmA0AAZEBFajlllsAAAAASUVORK5CYII=";

  return (
    failedLoad ?
      <Image
        className={styles.icon}
        alt="Icon-round-Question mark"
        src="/unknown_game.png"
        width={width} height={height}
        blurDataURL={blurDataURL}
        placeholder="blur" />
      :
      <Image
        className={styles.icon}
        loader={imageLoader}
        src={`${game.appid}/${game.img_icon_url}.jpg`}
        alt={"GameIcon"}
        width={width} height={height}
        blurDataURL={blurDataURL}
        placeholder="blur"
        onError={() => { setFailedLoad(true) }} />

  )
}
