import request from '@/components/Request';

export async function listTransactions(params: any) {
    return await request.get('/api/transactions', params);
}

export async function listTotalTransaction(params: any) {
    return await request.get('/api/transaction/total', params);
}

export async function saveTransaction(params: any) {
    return await request.post('/api/transaction', params);
}