import type { NextPage } from "next";
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

  const onClick = async (): Promise<void> => {
    const res = await Api.postPlayground().catch(() => ({
      error: "API request failure",
    }));
    if ("error" in res) {
      setState({ state: "ERROR", error: res.error });
      return;
    }

    router.push(`/playground/${res.id}`);
  };

  return (
    <div className="min-h-screen pt-0 pb-2 flex flex-col justify-center items-center h-screen">
      <Head>
        <title>Postgres Playground</title>
        <meta name="description" content="PostgreSQL Playground" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col justify-center items-center">
        <h1 className="m-2 text-6xl text-center">Postgres Playground</h1>

        <p className="m-2 text-2xl text-center">
          Get started by Creating Workspace
        </p>

        {state.state === "ERROR" && (
          <div role="alert" className="m-2">
            <div className="bg-red-500 text-white font-bold rounded-t px-4 py-2">
              Error
            </div>
            <div className="border border-t-0 border-red-400 rounded-b bg-red-100 px-4 py-3 text-red-700">
              <p>{state.error}</p>
            </div>
          </div>
        )}

        <button className="m-2 btn btn-blue" onClick={onClick}>
          Create Workspace
        </button>
      </main>
    </div>
  );
};

export default Home;
