const TOKEN_KEY = 'ACC_TOKEN';
const STORAGE_LEDGER_KEY = 'ACC_CURRENT_LEDGER';

export function getToken() {
  return sessionStorage.getItem(TOKEN_KEY);
}

export function setToken(accessToken: string) {
  sessionStorage.setItem(TOKEN_KEY, accessToken);
}

export function clear() {
  sessionStorage.clear();
}

