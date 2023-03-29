import request from '@/components/Request';

export async function listTransactions(params:any) {
    return await request.get('/api/transactions', );
}
