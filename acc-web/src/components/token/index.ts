const TOKEN_KEY = "ACC-TOKEN"

export function getToken() {
  return sessionStorage.getItem(TOKEN_KEY)
}

export function setToken(accessToken: string) {
  sessionStorage.setItem(TOKEN_KEY, accessToken)
}
