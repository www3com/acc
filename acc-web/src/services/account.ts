import request from '@/components/Request';

const ACCOUNTS_CACHE_KEY = 'ACCOUNTS';

export async function getOverview() {
    return await request.get('/api/accounts/overview');
}

export async function listAccounts() {
    return await request.get('/api/accounts');
}

export async function listAccountsByTypes(types: number[]) {
    let accounts = await listAllAccounts();
    let arr = [];
    for (const account of accounts) {
        if (types.includes(account.type)) {
            arr.push(account)
        }
    }
    return arr;
}

export async function saveAccount(account: any) {
    sessionStorage.removeItem(ACCOUNTS_CACHE_KEY);
    return await request.post('/api/accounts', account);
}

export async function deleteAccount(code: string) {
    sessionStorage.removeItem(ACCOUNTS_CACHE_KEY);
    return await request.delete('/api/accounts', {code: code});
}

export async function updateAccount(account: any) {
    sessionStorage.removeItem(ACCOUNTS_CACHE_KEY);
    return await request.put('/api/accounts', account);
}


async function listAllAccounts() {
    let json = sessionStorage.getItem(ACCOUNTS_CACHE_KEY);
    if (json != null) {
        return JSON.parse(json);
    }
    let r = await request.get('/api/accounts/all');
    sessionStorage.setItem(ACCOUNTS_CACHE_KEY, JSON.stringify(r.data))
    return r.data;
}
