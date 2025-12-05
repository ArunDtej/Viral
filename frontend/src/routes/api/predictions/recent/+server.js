import { json } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';

const BACKEND_URL = env.PRIVATE_API_URL || 'http://127.0.0.1:8080';

/** @type {import('./$types').RequestHandler} */
export async function GET() {
    try {
        const response = await fetch(`${BACKEND_URL}/api/predictions/recent`);

        if (!response.ok) {
            console.error('Failed to fetch recent predictions:', response.status);
            return json([]);
        }

        const data = await response.json();
        return json(data);
    } catch (err) {
        console.error('Recent Predictions Proxy API error:', err);
        return json([]);
    }
}
