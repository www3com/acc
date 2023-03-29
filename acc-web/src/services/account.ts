import request from '@/components/Request';

export async function getOverview() {
    return await request.get('/api/account/overview');
}

export async function listAccounts() {
    return await request.get('/api/accounts');
}

export async function listIncomes() {
    return await request.get('/api/account/incomes');
}

export async function listExpenses() {
    return await request.get('/api/account/expenses');
}

export async function listIncomeExpenses() {
    return await request.get('/api/account/income-expenses');
}

export async function saveAccount(account: any) {
    return await request.post('/api/account', account);
}

export async function deleteAccount(code: string) {
    return await request.delete('/api/account', {code: code});
}

export async function updateName(id: number, name: string) {
    return await request.put('/api/account/name', {id: id, name: name});
}

export async function updateRemark(id: number, remark: string) {
    return await request.put('/api/account/remark', {id: id, remark: remark});
}

export async function updateBalance(id: number, amount: number) {
    return await request.put('/api/account/balance', {id: id, amount: amount});
}