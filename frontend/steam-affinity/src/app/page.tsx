import Image from "next/image";
import styles from "./page.module.css";
import ProfileCard from "./components/ProfileCard/ProfileCard";

export default function Home() {
  return (
    <>
      <header className={styles.header}>
        <h1>Steam Affinity</h1>
      </header>
      <main className={styles.main}>
        <ProfileCard steamID="76561198081577408" />
      </main>
    </>
  );
}
