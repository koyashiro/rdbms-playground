type HttpMethod = "GET" | "POST" | "DELETE";
const Protcol = "https://";

const API_ENDPOINT_HOST = "";

export const fetchApi = async (
  method: HttpMethod,
  endpoint: string,
  { parameter, body }: { parameter?: any; body?: any }
): Promise<Response> => {
  if (!API_ENDPOINT_HOST) throw new Error();

  const parameterString = parameter
    ? "?" + new URLSearchParams(parameter).toString()
    : "";
  const url = Protcol + API_ENDPOINT_HOST + endpoint + parameterString;
  return await fetch(url, {
    method: method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
};
