<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { Share } from 'lucide-svelte'; // Assuming lucide-svelte is installed
	import PredictionText from '$lib/components/PredictionText.svelte';
	import PredictionButtons from '$lib/components/PredictionButtons.svelte';

	export let data: PageData;

	let mounted = false;
	let generatedPrediction: string = data.prediction.prediction;

	onMount(() => {
		mounted = true;
	});

	function share() {
		if (navigator.share) {
			navigator
				.share({
					title: data.prediction.title,
					text: `Hey ${data.prediction.user_data.name}, I just found out my viral prediction: ${generatedPrediction}`,
					url: window.location.href
				})
				.catch(console.error);
		} else {
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
	<div class="form-container" class:animate-slide-in-up={mounted}>
		<div class="title-container" class:animate-fade-in={mounted}>
			<span class="text-5xl mb-4 block">{data.prediction.icon}</span>
			<h1>{data.prediction.title}</h1>
			<p class="animate-text" style="--anim-delay: 0.2s">
				This secret is just for {data.prediction.user_data.name}.
			</p>
		</div>

		<div class="form-group" class:animate-fade-in={mounted} style="--anim-delay: 0.8s">
			<PredictionText text={generatedPrediction} />
		</div>

		<div class="form-group" style="--anim-delay: 1.0s">
			<button class="submit-btn" on:click={share}>
				<Share size={20} />
				<span>Share your prediction</span>
			</button>
		</div>

		<PredictionButtons slug="when-will-you-become-a-millionaire" />
	</div>
</div>

<!-- SEO: Invisible links to other predictions -->
<div
	class="inactive-links"
	aria-hidden="true"
	style="display: none; visibility: hidden; height: 0; width: 0; overflow: hidden;"
>
	{#if data.prediction.previous_predictions}
		{#each data.prediction.previous_predictions as p}
			<a href="/viral/{p.page_type}/{p.uid}">{p.page_type} prediction {p.uid}</a>
		{/each}
	{/if}
	{#if data.prediction.next_predictions}
		{#each data.prediction.next_predictions as p}
			<a href="/viral/{p.page_type}/{p.uid}">{p.page_type} prediction {p.uid}</a>
		{/each}
	{/if}
</div>

<style>
	.page-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem;
	}

	.form-container {
		position: relative;
		width: 100%;
		max-width: 550px; /* Adjusted max-width for result display */
		background: rgba(255, 255, 255, 0.03);
		backdrop-filter: blur(20px);
		-webkit-backdrop-filter: blur(20px);
		padding: 3rem;
		border-radius: 24px;
		border: 1px solid rgba(255, 255, 255, 0.1);
		box-shadow: 0 20px 60px rgba(0, 0, 0, 0.4);
		opacity: 0;
	}

	.form-container::before {
		content: '';
		position: absolute;
		inset: 0;
		border-radius: 24px;
		padding: 2px;
		background: linear-gradient(
			135deg,
			rgba(168, 85, 247, 0.5) 0%,
			rgba(236, 72, 153, 0.5) 50%,
			rgba(245, 158, 11, 0.5) 100%
		);
		-webkit-mask:
			linear-gradient(#fff 0 0) content-box,
			linear-gradient(#fff 0 0);
		-webkit-mask-composite: xor;
		mask-composite: exclude;
		pointer-events: none;
	}
	.title-container {
		text-align: center;
		margin-bottom: 2.5rem;
		opacity: 0;
	}

	h1 {
		font-size: 2.25rem;
		font-weight: 700;
		color: #f9fafb;
		margin: 0;
	}

	.title-container p {
		margin-top: 0.5rem;
		font-size: 1rem;
		color: var(--text-muted-color);
	}

	.form-group {
		margin-bottom: 1.5rem;
	}

	.submit-btn {
		position: relative;
		display: flex; /* Changed to flex for icon */
		align-items: center; /* Center items vertically */
		justify-content: center; /* Center items horizontally */
		gap: 0.75rem; /* Space between icon and text */
		width: 100%;
		padding: 1rem;
		border: none;
		border-radius: 12px;
		background: linear-gradient(135deg, #a855f7 0%, #ec4899 50%, #f59e0b 100%);
		background-size: 200% 200%;
		color: white;
		font-size: 1.125rem;
		font-weight: 600;
		cursor: pointer;
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease,
			background-position 0.3s ease;
		margin-top: 2rem;
		box-shadow: 0 8px 24px rgba(168, 85, 247, 0.4);
	}

	.submit-btn:hover {
		transform: translateY(-2px);
		background-position: 100% 50%;
		box-shadow:
			0 12px 32px rgba(168, 85, 247, 0.5),
			0 0 40px rgba(168, 85, 247, 0.3);
	}

	.submit-btn:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}

	.animate-slide-in-up {
		animation: slide-in-up 0.7s cubic-bezier(0.25, 0.46, 0.45, 0.94) forwards;
	}

	.animate-fade-in {
		animation: fade-in 0.6s ease-out forwards;
		animation-delay: var(--anim-delay, 0s);
	}

	@keyframes slide-in-up {
		from {
			transform: translateY(30px);
			opacity: 0;
		}
		to {
			transform: translateY(0);
			opacity: 1;
		}
	}

	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
</style>
