import { error } from '@sveltejs/kit';
import { getPrediction } from '$lib/predictions';

/** @type {import('./$types').PageLoad} */
export function load({ params }) {
	const prediction = getPrediction(params.slug);

	if (prediction) {
		return {
			prediction
		};
	}

	throw error(404, 'Prediction not found');
}
