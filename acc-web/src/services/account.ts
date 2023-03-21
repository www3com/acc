import request from '@/components/Request';

export async function listAccounts() {
    return await request.get('/api/accounts');
}

export async function saveAccount(account: any) {
    return await request.post('/api/account', account);
}

export async function deleteAccount(code: string) {
    return await request.delete('/api/account', {code: code});
}