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

export function getCurrentLedger() {
  let json = localStorage.getItem(STORAGE_LEDGER_KEY);
  return JSON.parse(typeof json === "string" ? json : "1");
}

export function setCurrentLedger(ledger: any) {
  localStorage.setItem(STORAGE_LEDGER_KEY, JSON.stringify(ledger));
}

