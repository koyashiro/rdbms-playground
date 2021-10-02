import type { GetServerSideProps, NextPage } from "next";
import { useRouter } from "next/dist/client/router";
import Head from "next/head";

import * as Api from "../lib/api/api";
import { useState } from "react";

type State =
  | {
      state: "NORMAL";
    }
  | {
      state: "ERROR";
      error: string;
    };

const Home: NextPage = () => {
  const router = useRouter();
  const [state, setState] = useState<State>({ state: "NORMAL" });

  const onClick = async (
    db: "mysql" | "postgres" | "mariadb"
  ): Promise<void> => {
    const res = await Api.postWorkspace(db).catch(() => ({
      error: "API request failure",
    }));
    if ("error" in res) {
      setState({ state: "ERROR", error: res.error });
      return;
    }

    router.push(`/workspaces/${res.id}`);
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen min-h-screen pt-0 pb-2">
      <Head>
        <title>RDBMS Playground</title>
        <meta name="description" content="RDBMS Playground" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center">
        <h1 className="m-2 text-6xl text-center">RDBMS Playground</h1>

        <p className="m-2 text-2xl text-center">
          Get started by Creating Workspace
        </p>

        {state.state === "ERROR" && (
          <div role="alert" className="m-2">
            <div className="px-4 py-2 font-bold text-white bg-red-500 rounded-t">
              Error
            </div>
            <div className="px-4 py-3 text-red-700 bg-red-100 border border-t-0 border-red-400 rounded-b">
              <p>{state.error}</p>
            </div>
          </div>
        )}

        <div>
          <button
            className="px-4 py-2 m-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
            onClick={() => onClick("mysql")}
          >
            Create MySQL Workspace
          </button>

          <button
            className="px-4 py-2 m-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
            onClick={() => onClick("mariadb")}
          >
            Create MariaDB Workspace
          </button>

          <button
            className="px-4 py-2 m-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
            onClick={() => onClick("postgres")}
          >
            Create PostgreSQL Workspace
          </button>
        </div>
      </main>
    </div>
  );
};

// For automatic static optimization suppression
export const getServerSideProps: GetServerSideProps = async (_) => {
  return {
    props: {},
  }
}

export default Home;
