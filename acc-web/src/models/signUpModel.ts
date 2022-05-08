import {signUp} from '../services/signUp';

export default function signUpModel() {

  const register = async (user: any) => {
    return await signUp(user)
  }

  return {register}
}
