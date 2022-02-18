import type { NextPage, GetServerSideProps } from "next";
import Error from "next/error";
import Head from "next/head";

import * as Api from "../../lib/api/api";
import Terminal from "../../components/terminal";

type Props = { workspace: Api.Workspace };

const Workspace: NextPage<Props> = ({ workspace }) => {
  if (!workspace) {
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
            const res = await Api.postWorkspaceQuery(workspace.id, {
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

export const getServerSideProps: GetServerSideProps<Props> = async (
  context
) => {
  const id = (() => {
    const { id } = context.query;
    if (id == null) {
      return id;
    }

    if (Array.isArray(id)) {
      return id[0];
    }

    return id;
  })();

  if (!id) {
    return { notFound: true };
  }

  const result = await Api.getWorkspaceById(id).catch(() => ({
    error: "API request failure",
  }));

  if ("error" in result) {
    return { notFound: true };
  }

  return { props: { workspace: result } };
};

export default Workspace;
