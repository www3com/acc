import { makeAutoObservable} from "mobx";

import {listAccount} from "@/services/account";

export class Account {
    accounts: [] = []

    constructor() {
        makeAutoObservable(this)
    }

    * list(ledgerId: number): any {
        const r = yield listAccount(ledgerId)
        this.accounts = r.data
    }
}