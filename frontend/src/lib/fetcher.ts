export function fetcherJson(url: string) {
  return fetch(url)
    .then((res) => res.json())
    .then((d) => {
      console.log(d);
      return d;
    })
    .catch((e) => console.log(e));
}

export function authorizeFetcherJson(url: string, token: string) {
  return fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then((res) => res.json())
    .then((d) => {
      console.log(d);
      return d;
    })
    .catch((e) => console.log(e));
}
