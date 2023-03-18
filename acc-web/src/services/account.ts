import request from '@/components/Request';

export async function listAccount() {
    return await request.get('/api/account');
}

export async function saveAccount(account: any) {
    // return await request.post('/api/account', {
    //   ...account,
    //   ledgerId: getCurrentLedger().id,
    // });
}