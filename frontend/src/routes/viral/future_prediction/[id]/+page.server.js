import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params, fetch }) {
	const response = await fetch(`http://localhost:8000/api/prediction/${params.id}`);

	if (response.ok) {
		const prediction = await response.json();
		return { prediction };
	}

	const errorData = await response.json().catch(() => null);
	const message = errorData?.error || 'Prediction not found';

	throw error(response.status, message);
}