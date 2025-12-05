import type { PageLoad } from './$types';
import { predictionList } from '$lib/predictions';
import { API_BASE_URL } from '$lib/config';

export const load: PageLoad = async ({ fetch }) => {
	let recentPredictions = [];
	try {
		const response = await fetch(`${API_BASE_URL}/api/predictions/recent`);
		if (response.ok) {
			const data = await response.json();
			recentPredictions = data.recent_predictions || [];
		}
	} catch (error) {
		console.error('Failed to fetch recent predictions:', error);
	}

	return {
		predictions: predictionList,
		recentPredictions
	};
};
