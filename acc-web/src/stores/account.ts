import {makeAutoObservable} from "mobx";

import {listAccount} from "@/services/account";

export class Account {
  accounts: [] = []

  constructor() {
    makeAutoObservable(this)
  }

  async list(ledgerId: number) {
    let r = await listAccount(ledgerId)
    this.accounts = r.data
  }
}
