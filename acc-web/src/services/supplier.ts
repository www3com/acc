import request from '@/components/Request';

export async function listSuppliers() {
    return await request.get('/api/suppliers');
}
