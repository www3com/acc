import request from "@/components/Request";

export async function listLedgers() {
    let r = await request.get('/api/ledger')
    return r.data
}