import {makeAutoObservable} from "mobx";
import {sha256} from "js-sha256";
import {signIn} from "@/services/login";

export class Login {
    constructor() {
        makeAutoObservable(this)
    }

    async login(username: string, password: string) {
        return await signIn(username, sha256.hex(password))
    }
}
