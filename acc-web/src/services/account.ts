import request from '@/components/Request';
import { getCurrentLedger } from '@/components/session';

export async function listAccount(ledgerId: number) {
  return await request.get('/api/account', { ledgerId: ledgerId });
}

export async function saveAccount(account: any) {
  return await request.post('/api/account', {
    ...account,
    ledgerId: getCurrentLedger().id,
  });
}