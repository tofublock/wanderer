import type { ListShare } from '$lib/models/list_share';
import type { User } from '$lib/models/user';
import { pb } from '$lib/pocketbase';
import { error, json, type RequestEvent } from '@sveltejs/kit';

export async function GET(event: RequestEvent) {
    const sort = event.url.searchParams.get('sort') ?? ""
    const filter = event.url.searchParams.get("filter") ?? "";

    try {
        const r: ListShare[] = await pb.collection('list_share').getFullList<ListShare>({
            sort: sort,
            filter: filter
        })
        for (const share of r) {
            const anonymous_user = await pb.collection('users_anonymous').getOne<User>(share.user)
            share.expand = {
                user: anonymous_user
            }
        }
        return json(r)
    } catch (e: any) {
        throw error(e.status, e);
    }
}

export async function PUT(event: RequestEvent) {
    const data = await event.request.json();

    try {
        const r = await pb.collection('list_share').create<ListShare>(data)
        return json(r);
    } catch (e: any) {
        throw error(e.status, e)
    }
}