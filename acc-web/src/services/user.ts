import request from "@/components/Request";

export async function signUp(user: any) {
  return await request.post('/api/user/sign-up', user)
}


export async function signIn(username: string, password: string) {
  return await request.post('/api/user/sign-in', {username: username, password: password})
}
