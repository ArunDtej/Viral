import { json, error } from '@sveltejs/kit';

import { API_BASE_URL } from '$lib/config';

import { env } from '$env/dynamic/private';

const BACKEND_URL = env.PRIVATE_API_URL || 'http://127.0.0.1:8080';

/** @type {import('./$types').RequestHandler} */
export async function POST({ request }) {
	try {
		const body = await request.json();
		console.log('Received prediction request body:', body);

		// Check if data is nested in user_data (new frontend structure) or top-level (legacy?)
		const userData = body.user_data || body;
		const slug = body.page_type || body.slug || 'future-prediction';

		// Validate required fields
		if (!userData.name || !userData.dob) {
			console.log('Validation failed:', { name: userData.name, dob: userData.dob });
			throw error(400, 'Missing required fields: name and dob are required');
		}

		const payload = {
			slug: slug,
			page_type: slug, // Add page_type to trigger newReq path in backend
			name: userData.name,
			dob: userData.dob,
			gender: userData.gender,
			partner: userData.partner,
			user_data: userData
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
			const errorMessage = errorData.error || errorData.message || 'Failed to create prediction';
			console.error('Backend error:', errorMessage, errorData);
			throw error(response.status, errorMessage);
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
