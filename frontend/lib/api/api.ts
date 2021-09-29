import { fetchApi } from "./util";

export type ErrorResponse = {
  error: string;
};

export type Workspace = {
  id: string;
  container: Container;
};

export type Container = {
  id: string;
  name: string;
  image: string;
};

export type Empty = Record<string, never>;

type Query = {
  query: string;
};

export type ExecuteResult = {
  columns: Column[];
  rows: string[][];
};

export type Column = {
  name: string;
  databaseType: string;
  nullable?: boolean;
  length?: number;
  precision?: number;
  scale?: number;
};

export const getAllWorkspaces = async (): Promise<
  Workspace[] | ErrorResponse
> => {
  const res = await fetchApi("GET", "/workspaces", {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Workspace[];
};

export const getWorkspaceById = async (
  id: number
): Promise<Workspace | ErrorResponse> => {
  const res = await fetchApi("GET", `/workspaces/${id}`, {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Workspace;
};

export const postWorkspace = async (
  db: "mysql" | "postgres"
): Promise<Workspace | ErrorResponse> => {
  // TODO: db selection
  const res = await fetchApi("POST", `/workspaces`, {
    body: { db },
  });
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Workspace;
};

export const deleteWorkspace = async (
  id: number
): Promise<Empty | ErrorResponse> => {
  const res = await fetchApi("DELETE", `/workspaces/${id}`, {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return {};
};

export const postWorkspaceQuery = async (
  id: string,
  body: Query
): Promise<ExecuteResult | ErrorResponse> => {
  const res = await fetchApi("POST", `/workspaces/${id}/query`, { body });
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as ExecuteResult;
};
