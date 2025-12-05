import { error } from '@sveltejs/kit';
import { getPrediction } from '$lib/predictions';
import { API_BASE_URL } from '$lib/config';

/** @type {import('./$types').PageLoad} */
export async function load({ params, fetch }) {
    const predictionId = params.id;


    // Fetch the prediction data from the backend
    const url = `${API_BASE_URL}/api/prediction/${predictionId}`;
    // console.log(`Fetching prediction from: ${url}`);

    const response = await fetch(url);

    if (!response.ok) {
        console.error(`Fetch failed for ${url} with status ${response.status}`);
        const text = await response.text();
        console.error(`Response body: ${text}`);

        let errorData;
        try {
            errorData = JSON.parse(text);
        } catch (e) {
            // ignore json parse error
        }

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
