export function fetcherJson(url: string) {
  return fetch(url)
    .then((res) => res.json())
    .catch((e) => console.log(e));
}

export function authorizeFetcherJson(url: string, token: string) {
  return fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then((res) => res.json())
    .catch((e) => console.log(e));
}
