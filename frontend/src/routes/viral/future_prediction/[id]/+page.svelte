<script lang="ts">
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { Share } from 'lucide-svelte';

	export let data: PageData;

	let mounted = false;
	onMount(() => {
		mounted = true;
	});

	function share() {
		if (navigator.share) {
			navigator
				.share({
					title: 'My Future Prediction!',
					text: `I just found out my future! Check it out:`,
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
	<title>Your Future Prediction</title>
	<meta name="description" content={data.prediction.prediction} />
</svelte:head>

<div class="page-container relative z-10">
	<div class="card" class:animate-card={mounted}>
		<div class="card-header" class:animate-header={mounted}>
			<h1 class="animate-text">Here is your future prediction</h1>
			<p class="animate-text" style="--anim-delay: 0.2s">
				This secret is just for {data.prediction.user_data.name}.
			</p>
		</div>

		<div class="card-body" class:animate-body={mounted}>
			<p class="prediction-text animate-text" style="--anim-delay: 0.8s">
				{data.prediction.prediction}
			</p>
		</div>

		<div class="card-footer" class:animate-footer={mounted}>
			<button class="share-btn" on:click={share}>
				<Share size={20} />
				<span>Share with friends</span>
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
		position: relative;
		max-width: 600px;
		background: rgba(255, 255, 255, 0.03);
		backdrop-filter: blur(20px);
		-webkit-backdrop-filter: blur(20px);
		border-radius: 24px;
		overflow: hidden;
		border: 1px solid rgba(255, 255, 255, 0.1);
		box-shadow: 0 20px 60px rgba(0, 0, 0, 0.4);
		opacity: 0;
		transform: translateY(20px);
	}

	.card::before {
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

	.card-header {
		position: relative;
		z-index: 1;
		padding: 3rem 3rem 2rem;
		text-align: center;
		background: rgba(255, 255, 255, 0.02);
		border-bottom: 1px solid rgba(255, 255, 255, 0.1);
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
		position: relative;
		z-index: 1;
		padding: 4rem 3rem;
		text-align: center;
		opacity: 0;
	}

	.prediction-text {
		font-size: 1.875rem;
		font-weight: 600;
		background: linear-gradient(135deg, #a855f7 0%, #ec4899 50%, #f59e0b 100%);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		line-height: 1.6;
		margin: 0;
		text-shadow: 0 0 40px rgba(168, 85, 247, 0.3);
		filter: drop-shadow(0 0 20px rgba(168, 85, 247, 0.2));
	}

	.card-footer {
		position: relative;
		z-index: 1;
		padding: 2rem 3rem;
		background: rgba(255, 255, 255, 0.02);
		border-top: 1px solid rgba(255, 255, 255, 0.1);
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
		background: linear-gradient(135deg, #a855f7 0%, #ec4899 50%, #f59e0b 100%);
		background-size: 200% 200%;
		color: white;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease,
			background-position 0.3s ease;
		box-shadow: 0 8px 24px rgba(168, 85, 247, 0.4);
	}

	.share-btn:hover {
		transform: translateY(-2px);
		background-position: 100% 50%;
		box-shadow:
			0 12px 32px rgba(168, 85, 247, 0.5),
			0 0 40px rgba(168, 85, 247, 0.3);
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
