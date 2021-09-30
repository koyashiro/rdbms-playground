import type { NextPage } from "next";
import { useRouter } from "next/dist/client/router";
import Head from "next/head";
import Error from "next/error";

import * as Api from "../../lib/api/api";
import Terminal from "../../lib/components/terminal";

const Workspace: NextPage = () => {
  const router = useRouter();

  const id = (() => {
    const { id } = router.query;
    if (id == null) {
      return id;
    }

    if (Array.isArray(id)) {
      return id[0];
    }

    return id;
  })();

  if (id == null) {
    return <Error statusCode={404} />;
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen min-h-screen pt-0 pb-2">
      <Head>
        <title>RDBMS Playground</title>
        <meta name="description" content="RDBMS Playground" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center w-screen h-screen">
        <Terminal
          command={async (cmd: string) => {
            const res = await Api.postWorkspaceQuery(id, {
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

export default Workspace;
