import request from "@/components/Request";

export async function signUp(user: any) {
    return await request.post('/api/sign-up', user)
}