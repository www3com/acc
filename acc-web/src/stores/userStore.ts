import {makeAutoObservable} from "mobx";
import {signIn, signUp} from "@/services/user";
import {sha256} from "js-sha256";

export class UserStore {
  name = ""

  constructor() {
    makeAutoObservable(this)
  }

  async login(username: string, password: string) {
    return await signIn(username, sha256.hex(password))
  }

  async register(user: any) {
    user.password = sha256.hex(user.password)
    return await signUp(user)
  }
}
