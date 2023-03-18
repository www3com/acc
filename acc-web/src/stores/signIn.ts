import {makeAutoObservable} from "mobx";
import {sha256} from "js-sha256";
import {signIn} from "@/services/signIn";

export class SignInStore {
    constructor() {
        makeAutoObservable(this)
    }

    async signIn(username: string, password: string) {
        return await signIn(username, sha256.hex(password))
    }
}
