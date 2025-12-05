import { error } from '@sveltejs/kit';
import { getPrediction } from '$lib/predictions';
import { API_BASE_URL } from '$lib/config';

/** @type {import('./$types').PageLoad} */
export async function load({ params, fetch }) {
    const predictionId = params.id;

    // Fetch the prediction data from the backend
    const response = await fetch(`${API_BASE_URL}/api/prediction/${predictionId}`);

    if (!response.ok) {
        const errorData = await response.json().catch(() => null);
        const message = errorData?.error || 'Prediction not found';
        throw error(response.status, message);
    }

    const backendPrediction = await response.json();
    // console.log('Backend Prediction:', backendPrediction); // Log the backend prediction

    // Now, get the frontend prediction details (title, icon, etc.) using the slug
    // This assumes backendPrediction.page_type will contain the slug
    const frontendPredictionDetails = getPrediction(backendPrediction.page_type);

    if (!frontendPredictionDetails) {
        throw error(404, 'Prediction type not found');
    }

    return {
        prediction: {
            ...frontendPredictionDetails,
            prediction: backendPrediction.prediction, // The actual generated prediction string
            user_data: backendPrediction.user_data,
            title: backendPrediction.title, // Use the title from the backend
            icon: frontendPredictionDetails.icon, // Use the icon from the frontend definition
            previous_predictions: backendPrediction.previous_predictions,
            next_predictions: backendPrediction.next_predictions
        }
    };
}
