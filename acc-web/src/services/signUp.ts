import request from "@/components/Request";

export async function signUp(user: any) {
  return await request.post('/api/user/sign-up', user)
}
