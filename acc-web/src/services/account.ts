import request from '@/components/Request';

export async function listAccounts() {
    return await request.get('/api/accounts');
}

export async function saveAccount(account: any) {
    // return await request.post('/api/account', {
    //   ...account,
    //   ledgerId: getCurrentLedger().id,
    // });
}