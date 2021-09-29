import type { NextPage } from "next";
import { useRouter } from "next/dist/client/router";
import Head from "next/head";

import * as Api from "../../lib/api/api";
import Terminal from "../../lib/components/terminal";

const Playground: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;
  const idNumber = Number(id);

  return (
    <div className="min-h-screen pt-0 pb-2 flex flex-col justify-center items-center h-screen">
      <Head>
        <title>Postgres Playground</title>
        <meta name="description" content="PostgreSQL Playground" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col justify-center items-center h-screen w-screen">
        <Terminal
          command={async (cmd: string) => {
            const res = await Api.postPlaygroundQuery(idNumber, {
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
