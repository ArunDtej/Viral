import { error } from '@sveltejs/kit';
import { getPrediction } from '$lib/predictions';

/** @type {import('./$types').PageLoad} */
export function load() {
    const prediction = getPrediction('what-cat-are-you');

    if (prediction) {
        return {
            prediction
        };
    }

    throw error(404, 'Prediction not found');
}
