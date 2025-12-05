import { json, error } from '@sveltejs/kit';

import { API_BASE_URL } from '$lib/config';

const BACKEND_URL = API_BASE_URL;

/** @type {import('./$types').RequestHandler} */
export async function POST({ request }) {
	try {
		const body = await request.json();

		// Validate required fields
		if (!body.name || !body.dob) {
			throw error(400, 'Missing required fields: name and dob are required');
		}

		const payload = {
			slug: body.slug || 'future-prediction',
			name: body.name,
			dob: body.dob,
			gender: body.gender
		};

		// Call the backend API
		const response = await fetch(`${BACKEND_URL}/api/predict`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(payload)
		});

		if (!response.ok) {
			const errorData = await response.json().catch(() => ({ error: 'Failed to create prediction' }));
			throw error(response.status, errorData.error || 'Failed to create prediction');
		}

		const data = await response.json();
		return json(data);
	} catch (err) {
		console.error('Prediction API error:', err);

		if (typeof err === 'object' && err !== null && 'status' in err) {
			throw err;
		}

		throw error(500, 'Internal server error');
	}
}
