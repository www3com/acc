import {makeAutoObservable} from "mobx";
import {listLedgers} from "@/services/ledger";

const STORAGE_LEDGER_KEY = 'ACC_CURRENT_LEDGER'

export class LayoutStore {
    visible = false
    currentLedger = {id: 0, name: '未选择'}
    ledgers: [] = []

    constructor() {
        makeAutoObservable(this)
        this.list()
    }

    setVisible(visible: boolean) {
        this.visible = visible
    }

    setCurrentLedger(ledger: any) {
        this.setVisible(false)
        this.currentLedger = ledger
        localStorage.setItem(STORAGE_LEDGER_KEY, JSON.stringify(ledger))
    }

    * list(): any {
        let ledgers = yield listLedgers()
        let ledger = localStorage.getItem(STORAGE_LEDGER_KEY)
        this.setCurrentLedger(ledger == null ? ledgers[0] : JSON.parse(ledger))
        this.ledgers = ledgers
    }
}