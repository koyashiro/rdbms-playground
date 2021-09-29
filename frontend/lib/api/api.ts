import { fetchApi } from "./util";

export type ErrorResponse = {
  error: string;
};

export type Playground = {
  id: string;
  db: string;
  version: string;
  container: Container;
};

export type Container = {
  id: string;
  image: string;
  status: string;
};

export type Empty = Record<string, never>;

type Query = {
  query: string;
};

export type ExecuteResult = {
  columns: ExportColumn[];
  rows: string[][];
};

export type ExportColumn = {
  name: string;
  databaseType: string;
  nullable?: boolean;
  length?: number;
  precision?: number;
  scale?: number;
};

export const getAllPlaygrounds = async (): Promise<
  Playground[] | ErrorResponse
> => {
  const res = await fetchApi("GET", "/playgrounds", {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Playground[];
};

export const getPlaygroundById = async (
  id: number
): Promise<Playground | ErrorResponse> => {
  const res = await fetchApi("GET", `/playgrounds/${id}`, {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Playground;
};

export const postPlayground = async (): Promise<Playground | ErrorResponse> => {
  // TODO: db selection
  const res = await fetchApi("POST", `/playgrounds`, {
    body: { db: "postgres" },
  });
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as Playground;
};

export const deletePlayground = async (
  id: number
): Promise<Empty | ErrorResponse> => {
  const res = await fetchApi("DELETE", `/playgrounds/${id}`, {});
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return {};
};

export const postPlaygroundQuery = async (
  id: number,
  body: Query
): Promise<ExecuteResult | ErrorResponse> => {
  const res = await fetchApi("POST", `/playgrounds/${id}/query`, { body });
  const json = await res.json();

  if (!res.ok) return json as ErrorResponse;
  return json as ExecuteResult;
};
