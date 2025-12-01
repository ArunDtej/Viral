import { json, error } from '@sveltejs/kit';

const BACKEND_URL = 'http://localhost:8000';

/** @type {import('./$types').RequestHandler} */
export async function POST({ request }) {
    try {
        const body = await request.json();

        // Validate required fields
        if (!body.name || !body.dob || !body.gender) {
            throw error(400, 'Missing required fields: name, dob, and gender are required');
        }

        // Call the backend API
        const response = await fetch(`${BACKEND_URL}/api/predict`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                page_type: 'future_prediction',
                name: body.name,
                dob: body.dob,
                gender: body.gender
            })
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ error: 'Failed to create prediction' }));
            throw error(response.status, errorData.error || 'Failed to create prediction');
        }

        const data = await response.json();
        return json(data);
    } catch (err) {
        console.error('Prediction API error:', err);

        if (err.status) {
            throw err;
        }

        throw error(500, 'Internal server error');
    }
}
