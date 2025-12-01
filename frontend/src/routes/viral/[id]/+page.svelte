<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';

	export let data: PageData;

	let mounted = false;
	onMount(() => {
		mounted = true;
	});

	// The generate function is now part of the prediction object from data.
	// We need to call it with appropriate arguments.
	// For simplicity, let's assume we have default values or get them from user input.
	let generatedPrediction: string = '';
	let name: string = 'Mysterious User'; // Placeholder
	let dob: string = '2000-01-01'; // Placeholder

	$: if (data.prediction && mounted) {
		generatedPrediction = data.prediction.generate(name, dob);
	}

	function share() {
		if (navigator.share) {
			navigator
				.share({
					title: 'My Viral Prediction!',
					text: `I just found out my viral prediction: ${generatedPrediction}`,
					url: window.location.href
				})
				.catch(console.error);
		} else {
			// Fallback for browsers that don't support navigator.share
			navigator.clipboard.writeText(window.location.href).then(
				() => {
					alert('Link copied to clipboard!');
				},
				() => {
					alert('Failed to copy link. Please copy the URL manually.');
				}
			);
		}
	}
</script>

<svelte:head>
	<title>{data.prediction.title} - Viral Predictions</title>
	<meta name="description" content={generatedPrediction} />
</svelte:head>

<div class="page-container relative z-10">
	<div class="card" class:animate-card={mounted}>
		<div class="card-header" class:animate-header={mounted}>
			<span class="text-5xl mb-4 block">{data.prediction.icon}</span>
			<h1 class="animate-text">{data.prediction.title}</h1>
			<p class="animate-text" style="--anim-delay: 0.2s">
				Unveiling your unique viral destiny!
			</p>
		</div>

		<div class="card-body" class:animate-body={mounted}>
			<p class="prediction-text animate-text" style="--anim-delay: 0.8s">
				{generatedPrediction}
			</p>
		</div>

		<div class="card-footer" class:animate-footer={mounted}>
			<button class="share-btn" on:click={share}>
				<span>Share your prediction</span>
			</button>
		</div>
	</div>
</div>

<style>
	/* Original card styles */
	.page-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem;
	}

	.card {
		max-width: 550px;
		background: var(--surface-color);
		border-radius: 24px;
		overflow: hidden;
		border: 1px solid var(--border-color);
		box-shadow: 0 20px 50px rgba(0, 0, 0, 0.6);
		opacity: 0;
		transform: translateY(20px);
	}

	.card-header {
		padding: 3rem 3rem 2rem;
		text-align: center;
		background: rgba(255, 255, 255, 0.03);
		border-bottom: 1px solid var(--border-color);
		opacity: 0;
	}

	h1 {
		font-size: 2rem;
		font-weight: 700;
		color: var(--text-color);
		margin: 0 0 0.5rem;
	}

	.card-header p {
		color: var(--text-muted-color);
		font-size: 1rem;
		margin: 0;
	}

	.card-body {
		padding: 4rem 3rem;
		text-align: center;
		opacity: 0;
	}

	.prediction-text {
		font-size: 1.75rem;
		font-weight: 500;
		color: var(--primary-color);
		line-height: 1.6;
		margin: 0;
		text-shadow: 0 0 20px rgba(139, 92, 246, 0.3);
	}

	.card-footer {
		padding: 2rem 3rem;
		background: rgba(255, 255, 255, 0.03);
		border-top: 1px solid var(--border-color);
		display: flex;
		justify-content: center;
		opacity: 0;
	}

	.share-btn {
		display: inline-flex;
		align-items: center;
		gap: 0.75rem;
		padding: 0.875rem 1.75rem;
		border: none;
		border-radius: 50px;
		background-image: linear-gradient(to right, var(--primary-color), var(--primary-hover));
		color: white;
		font-size: 1rem;
		font-weight: 500;
		cursor: pointer;
		transition:
			transform 0.2s ease-out,
			box-shadow 0.2s ease-out;
	}

	.share-btn:hover {
		transform: translateY(-3px) scale(1.03);
		box-shadow: 0 8px 25px rgba(139, 92, 246, 0.3);
	}

	/* Animations */
	.animate-card {
		animation: slide-up 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
	}
	.animate-header {
		animation: fade-in 0.6s 0.4s ease-out forwards;
	}
	.animate-body {
		animation: fade-in 0.6s 0.6s ease-out forwards;
	}
	.animate-footer {
		animation: fade-in 0.6s 0.8s ease-out forwards;
	}
	.animate-text {
		opacity: 0;
		transform: translateY(15px);
		animation: slide-up-text 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
		animation-delay: calc(0.5s + var(--anim-delay, 0s));
	}

	@keyframes slide-up {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
	@keyframes fade-in {
		to {
			opacity: 1;
		}
	}
	@keyframes slide-up-text {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
</style>