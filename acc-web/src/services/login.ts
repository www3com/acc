import request from "@/components/Request";

export async function signIn(username: string, password: string) {
  return await request.post('/api/user/sign-in', {username: username, password: password})
}
