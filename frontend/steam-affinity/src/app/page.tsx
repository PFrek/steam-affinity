import Image from "next/image";
import styles from "./page.module.css";
import ProfileCard from "./components/ProfileCard/ProfileCard";
import { getFriendsList } from "./actions/actions";
import FriendsList from "./components/FriendsList/FriendsList";

export default async function Home() {
  const steamid = "76561198081577408";

  return (
    <>
      <header className={styles.header}>
        <h1>Steam Affinity</h1>
      </header>
      <main className={styles.main}>
        <ProfileCard steamid={steamid} affinity={null} />
        <FriendsList steamid={steamid} />
      </main>

    </>
  );
}
