import request from "@/components/Request";

const SESSION_LEDGERS_KEY = "LEDGERS"
const SESSION_SELECTED_LEDGER_KEY = "SELECTED-LEDGER"

export interface Ledger {
  key?: string
  label?: string
  icon?: any
  type?: string
}

export function clearAll() {
  sessionStorage.clear()
}

export function selectLedger(ledgerId: number) {
  sessionStorage.setItem(SESSION_SELECTED_LEDGER_KEY, String(ledgerId))
}

export function getSelectLedger(): Ledger {
  let ledgerId = sessionStorage.getItem(SESSION_SELECTED_LEDGER_KEY)
  let sessionLedgers = getSessionLedgers()
  if (sessionLedgers.length == 0) {
    return null
  }

  for (const ledger of sessionLedgers) {
    if (ledger.key == ledgerId) {
      return ledger
    }
  }

  return sessionLedgers[0]
}

export async function listLedgers(): Promise<Ledger[]> {
  let sessionLedgers = getSessionLedgers()
  if (sessionLedgers.length == 0) {
    let r = await request.get('/api/ledger')
    sessionLedgers = r.data
    sessionStorage.setItem(SESSION_LEDGERS_KEY, JSON.stringify(sessionLedgers))
  }

  return sessionLedgers
}

export async function saveLedger(ledger: object) {
  await request.post("/api/ledger", ledger)
  sessionStorage.removeItem(SESSION_LEDGERS_KEY)
}


function getSessionLedgers(): Ledger[] {
  let sessionLedgers = sessionStorage.getItem(SESSION_LEDGERS_KEY)
  if (sessionLedgers != null) {
    return JSON.parse(sessionLedgers)
  }
  return []
}
