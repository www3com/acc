import { makeAutoObservable} from "mobx";

export class BillStore {

    constructor() {
        makeAutoObservable(this)
    }

    * list(ledgerId: number): any {
        // const r = yield listAccount(ledgerId)
        // this.accounts = r.data
    }
}