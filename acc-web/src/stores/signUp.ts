import {makeAutoObservable} from "mobx";
import {sha256} from "js-sha256";
import {signUp} from "@/services/signUp";

export class SignUp {
    okVisible = false;

    constructor() {
        makeAutoObservable(this)
    }

    async register(user: any) {
        user.password = sha256.hex(user.password)
        return await signUp(user)
    }

    showOk(visible: boolean) {
        this.okVisible = visible;
    }
}
