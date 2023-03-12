import request from "@/components/Request";

export async function signIn(username: string, password: string) {
  return await request.get('/api/user/sign-in', {username: username, password: password})
}
