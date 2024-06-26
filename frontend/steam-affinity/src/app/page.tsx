'use client'

import { useState } from "react";
import styles from "./page.module.css";

export default function Home() {
  const [userID, setUserID] = useState<string>("");
  const [validFormState, setValidFormState] = useState<boolean>(true);

  function extractID(entry: string): string {
    const match = entry.match(/(\d+)/);
    if (match === null) {
      setValidFormState(false);
      return "";
    }

    return match[1];
  }

  return (
    <>
      <header className={styles.header}>
        <h1>Steam Affinity</h1>
      </header>
      <main className={styles.main}>
        <div className={styles.section}>
          <form action={`${userID}`} method="POST" className={styles.form}>
            <label htmlFor="userID">Steam ID:</label>
            <input name="userID" type="text" required value={userID} onChange={(val) => setUserID(extractID(val.target.value))} />
            <button type="submit">Submit</button>
          </form>
        </div>
        <div className={styles.section}>
          <h2>What is Steam Affinity?</h2>
          <p>Steam Affinity is a website project that allows comparing a user&apos;s Steam game library with those of their Steam friends.</p>
          <p>The user&apos;s friends are ranked according to the similarity of their game libraries. Users with <em>similarly sized libraries</em> and with <em>many games in common</em> are ranked higher.</p>
          <p>
            You can also see the games you both own, or the games that are missing from your library that the other user owns.
          </p>
        </div>
        <div className={styles.section}>
          <h2>How to Use</h2>
          <p>
            Enter your numeric Steam ID or that of a user you&apos;d like to verify in the website URL.
          </p>
          <p>
            You can also use the form above to enter the numeric ID. Doing so and clicking the submit button will redirect you to the appropriate page.
          </p>
          <p>
            You can also paste the user&apos;s profile URL, and the form should extract the numeric part automatically.
          </p>
        </div>
        <div className={styles.section}>
          <h2>How is the Affinity calculated?</h2>
          <p>
            The following formula is used to calculate the Affinity between two users:
          </p>
          <p className={styles.formula}>
            Affinity<sub>A,B</sub> = Similarity<sub>A,B</sub> ⋅ Weight<sub>A,B</sub>
          </p>
          <p>
            The Similarity between two user&apos;s libraries is the number of matching games divided by the total number of games owned by both users, like so:
          </p>
          <p className={styles.formula}>
            N<sub>A</sub> = number of games owned by user A
          </p>
          <p className={styles.formula}>
            N<sub>B</sub> = number of games owned by user B
          </p>
          <p className={styles.formula}>
            Similarity<sub>A,B</sub> = number of matching games / (N<sub>A</sub> + N<sub>B</sub>)
          </p>
          <p>
            This value is then weighted, giving users with libraries of similar sizes a higher weight:
          </p>
          <p className={styles.formula}>
            Weight<sub>A,B</sub> = (2 ⋅ N<sub>A</sub> ⋅ N<sub>B</sub>) / (N<sub>A</sub> + N<sub>B</sub>)
          </p>
        </div>
        <div className={styles.section}>
          <h2>Limitations</h2>
          <p>
            Since this website uses the public Steam API to get the users&apos; game libraries and friends lists, it is limited by the 100,000 daily calls they have set up.
          </p>
          <p>
            Some effort was made to cache that information whenever possible to reduce the number of calls made, but it remains a possible issue.
          </p>
        </div>
      </main >
    </>
  );
}
