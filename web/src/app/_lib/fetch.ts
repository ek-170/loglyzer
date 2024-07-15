export async function fetcher(
  input: URL | RequestInfo,
  init?: RequestInit | undefined,
) {
  const res = await fetch(input, init);
  return res.json();
}
