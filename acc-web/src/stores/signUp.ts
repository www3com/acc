import {makeAutoObservable} from "mobx";
import {sha256} from "js-sha256";
import {signUp} from "@/services/signUp";

export class SignUpStore {
    visible = false;

    constructor() {
        makeAutoObservable(this)
    }

    async signUp(user: any) {
        user.password = sha256.hex(user.password)
        return await signUp(user)
    }

    show(visible: boolean) {
        this.visible = visible;
    }
}
