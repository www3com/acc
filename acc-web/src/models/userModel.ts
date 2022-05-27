import {signIn, signUp} from '../services/user';
import {sha256} from "js-sha256";

export default function userModel() {

  const login = async (username: string, password: string) => {
    return await signIn(username, sha256.hex(password))
  }

  const register = async (user: any) => {
    user.password = sha256.hex(user.password)
    return await signUp(user)
  }

  return {login, register}
}
