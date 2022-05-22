import {signIn, signUp} from '../services/user';

export default function userModel() {

  const login = async (username: string, password: string) => {
    return await signIn(username, password)
  }

  const register = async (user: any) => {
    return await signUp(user)
  }

  return {login, register}
}
