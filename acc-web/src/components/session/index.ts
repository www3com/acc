const TOKEN_KEY = 'ACC_TOKEN';

export function getToken() {
  return sessionStorage.getItem(TOKEN_KEY);
}

export function setToken(accessToken: string) {
  sessionStorage.setItem(TOKEN_KEY, accessToken);
}

export function clear() {
  sessionStorage.clear();
}

