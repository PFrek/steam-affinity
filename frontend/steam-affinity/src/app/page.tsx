import Image from "next/image";
import styles from "./page.module.css";
import ProfileCard from "./components/ProfileCard/ProfileCard";
import { getFriendsList, getPlayerSummaries } from "./actions/actions";
import FriendsList from "./components/FriendsList/FriendsList";

export default async function Home() {
  const steamid = "76561198081577408";

  const resp = await getPlayerSummaries(steamid);
  const summary = resp[0];

  return (
    <>
      <header className={styles.header}>
        <h1>Steam Affinity</h1>
      </header>
      <main className={styles.main}>
        <ProfileCard summary={summary} affinity={null} affinityBoundaries={{ min: 0, max: 0 }} />
        <FriendsList steamid={steamid} />
      </main>

    </>
  );
}
