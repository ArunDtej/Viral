import { json, error } from '@sveltejs/kit';

const BACKEND_URL = 'http://127.0.0.1:8080';

/** @type {import('./$types').RequestHandler} */
export async function GET({ params }) {
    try {
        const response = await fetch(`${BACKEND_URL}/api/prediction/${params.id}`);

        if (!response.ok) {
            // Forward the error status and message
            const text = await response.text();
            let errorData;
            try {
                errorData = JSON.parse(text);
            } catch (e) {
                errorData = { error: 'Backend error', details: text };
            }
            throw error(response.status, errorData.error || 'Prediction fetch failed');
        }

        const data = await response.json();
        return json(data);
    } catch (err) {
        console.error('Prediction Proxy API error:', err);
        if (typeof err === 'object' && err !== null && 'status' in err) {
            throw err;
        }
        throw error(500, 'Internal server error during proxy');
    }
}
