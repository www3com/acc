import request from "@/components/Request";

export async function listAccount(ledgerId: number)  {
  return await request.get('/api/account', {ledgerId: ledgerId})
}
