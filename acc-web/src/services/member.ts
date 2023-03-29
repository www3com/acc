import request from '@/components/Request';

export async function listMembers() {
    return await request.get('/api/members');
}
