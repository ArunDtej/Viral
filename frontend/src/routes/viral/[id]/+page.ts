import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { getPrediction } from '$lib/predictions';

export const load: PageLoad = ({ params }) => {
	const prediction = getPrediction(params.id);

	if (!prediction) {
		error(404, 'Prediction not found');
	}

	return {
		prediction
	};
};
