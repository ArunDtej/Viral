import { error } from '@sveltejs/kit';
import { API_BASE_URL } from '$lib/config';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params, fetch }) {
	const response = await fetch(`${API_BASE_URL}/api/prediction/${params.id}`);

	if (response.ok) {
		const prediction = await response.json();
		return { prediction };
	}

	const errorData = await response.json().catch(() => null);
	const message = errorData?.error || 'Prediction not found';

	throw error(response.status, message);
}