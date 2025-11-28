import type { PageLoad } from './$types';
import { predictionList } from '$lib/predictions';

export const load: PageLoad = () => {
	return {
		predictions: predictionList
	};
};
