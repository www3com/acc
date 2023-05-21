import request from '@/components/Request';

export async function listTransactions(params: any) {
    return await request.post('/api/transactions/list', getTransaction(params));
}

export async function listTotalTransaction(params: any) {
    return await request.post('/api/transactions/total', getTransaction(params));
}

export async function saveTransaction(params: any) {
    return await request.post('/api/transactions', params);
}

function getTransaction(params: any) {
    return {
        ...params,
        startTime: params.startTime?.startOf('day').valueOf(),
        endTime: params.endTime?.endOf('day').valueOf(),
        accounts: params.accounts.map((value: any) => value.id),
        cpAccounts: params.cpAccounts.map((value: any) => value.id),
        projects: params.projects.map((value: any) => value.id),
        members: params.members.map((value: any) => value.id),
        suppliers: params.suppliers.map((value: any) => value.id),
    }
}