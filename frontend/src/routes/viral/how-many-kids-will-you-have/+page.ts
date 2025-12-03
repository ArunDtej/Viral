import { error } from '@sveltejs/kit';
import { getPrediction } from '$lib/predictions';

/** @type {import('./$types').PageLoad} */
export function load() {
    const prediction = getPrediction('how-many-kids-will-you-have');

    if (prediction) {
        return {
            prediction
        };
    }

    throw error(404, 'Prediction not found');
}
