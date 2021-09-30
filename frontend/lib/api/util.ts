import getConfig from "next/config";

type HttpMethod = "GET" | "POST" | "DELETE";

export const fetchApi = async (
  method: HttpMethod,
  endpoint: string,
  {
    parameter,
    body,
  }: { parameter?: Record<string, string>; body?: Record<string, string> }
): Promise<Response> => {
  const { publicRuntimeConfig } = getConfig();
  // TODO: use environment variable
  const apiHostUri =
    typeof window === "undefined"
      ? "http://backend:1323"
      : publicRuntimeConfig.apiHostUri;
  const parameterString = parameter
    ? "?" + new URLSearchParams(parameter).toString()
    : "";
  const url = apiHostUri + endpoint + parameterString;
  return await fetch(url, {
    method: method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
};
