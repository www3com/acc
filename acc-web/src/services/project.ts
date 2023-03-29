import request from '@/components/Request';

export async function listProjects() {
    return await request.get('/api/projects');
}
