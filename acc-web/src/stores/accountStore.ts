import {makeAutoObservable} from "mobx";
import {signIn, signUp} from "@/services/user";
import {sha256} from "js-sha256";
import {listAccount} from "@/services/account";

export class AccountStore {
  accounts: [] = []

  constructor() {
    makeAutoObservable(this)
  }

  async list(ledgerId: number) {
    let r = await listAccount(ledgerId)
    this.accounts = r.data
  }
}
