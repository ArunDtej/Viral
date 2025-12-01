export interface Prediction {
	slug: string;
	title: string;
	icon: string; // Added icon property
	stars: number; // Rating stars
	generate: (name: string, dob: string) => string;
}

export const predictions: { [key: string]: Omit<Prediction, 'slug'> } = {
	'future-prediction': {
		title: 'Future Prediction',
		icon: '✨',
		stars: 5,
		generate: (name, dob) => {
			const days = Math.floor(Math.random() * 365) + 1;
			return `In the next ${days} days, something amazing will happen to ${name}.`;
		}
	},
	'when-will-you-get-married': {
		title: 'When will you get married?',
		icon: '❤️', // Placeholder icon
		stars: 5,
		generate: () => {
			const years = Math.floor(Math.random() * 10) + 1;
			return `You will get married in ${years} years.`;
		}
	},
	'when-will-you-become-a-millionaire': {
		title: 'When will you become a millionaire?',
		icon: '💰', // Placeholder icon
		stars: 5,
		generate: () => {
			const years = Math.floor(Math.random() * 20) + 5;
			return `You will become a millionaire in ${years} years.`;
		}
	},
	'your-future-childs-face': {
		title: "Your future child's face.",
		icon: '👶', // Placeholder icon
		stars: 4,
		generate: () => `Here is what your future child could look like.`
	},
	'which-country-will-you-live-in': {
		title: 'Which country will you live in?',
		icon: '✈️', // Placeholder icon
		stars: 5,
		generate: () => {
			const countries = ['Japan', 'Italy', 'Australia', 'Canada', 'Brazil', 'New Zealand'];
			const country = countries[Math.floor(Math.random() * countries.length)];
			return `You will live in ${country}.`;
		}
	},
	'your-2025-fortune-reading': {
		title: 'Your 2025 fortune reading.',
		icon: '🔮', // Placeholder icon
		stars: 5,
		generate: () => {
			const fortunes = [
				'A great opportunity will arise.',
				'You will find unexpected joy.',
				'A long-held wish will come true.'
			];
			return fortunes[Math.floor(Math.random() * fortunes.length)];
		}
	}
};

export const predictionList = Object.entries(predictions).map(([slug, details]) => ({
	slug,
	...details
}));

export function getPrediction(slug: string | undefined) {
	if (!slug || !predictions[slug]) {
		return null;
	}
	return {
		slug,
		...predictions[slug]
	};
}
