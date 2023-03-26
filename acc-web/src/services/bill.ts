import request from '@/components/Request';

export async function listExpenses() {
    return await request.get('/api/account/expense');
}
