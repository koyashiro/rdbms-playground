import type { NextPage } from "next";
import { useRouter } from "next/dist/client/router";
import Head from "next/head";

import * as Api from "../../lib/api/api";
import Terminal from "../../lib/components/terminal";

const Playground: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;
  if (typeof id !== "string") {
    throw new Error("id is not string");
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen min-h-screen pt-0 pb-2">
      <Head>
        <title>Postgres Playground</title>
        <meta name="description" content="PostgreSQL Playground" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center w-screen h-screen">
        <Terminal
          command={async (cmd: string) => {
            const res = await Api.postPlaygroundQuery(id, {
              query: cmd,
            }).catch(() => ({
              error: "API request failure",
            }));
            return JSON.stringify(res, null, "  ");
          }}
        />
      </main>
    </div>
  );
};

export default Playground;
