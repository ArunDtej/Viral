<script lang="ts">
	import { onMount } from 'svelte';
	import { fly, fade } from 'svelte/transition';
	import { backOut } from 'svelte/easing';

	export let text: string;

	let show = false;
	let generating = true;

	$: words = text.split(' ');

	onMount(() => {
		// Simulate a magical generation process
		setTimeout(() => {
			generating = false;
			// Small delay to ensure clean transition
			setTimeout(() => {
				show = true;
			}, 100);
		}, 2000);
	});
</script>

<div class="container">
	{#if generating}
		<div class="generating-state" out:fade={{ duration: 400 }}>
			<div class="magic-orb">
				<div class="inner-orb"></div>
				<div class="ring"></div>
				<div class="sparkles">✨</div>
			</div>
			<p class="generating-text">Divining your destiny...</p>
		</div>
	{:else}
		<div class="prediction-wrapper" aria-label={text}>
			{#if show}
				{#each words as word, i}
					<span class="word" in:fly={{ y: 40, duration: 800, delay: i * 100, easing: backOut }}>
						{word}
					</span>
				{/each}
			{/if}
		</div>
	{/if}
</div>

<style>
	.container {
		min-height: 200px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.generating-state {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
	}

	.magic-orb {
		position: relative;
		width: 60px;
		height: 60px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.inner-orb {
		position: absolute;
		width: 30px;
		height: 30px;
		background: radial-gradient(circle, #f472b6 0%, #c084fc 100%);
		border-radius: 50%;
		box-shadow: 0 0 20px #c084fc;
		animation: pulse 1.5s ease-in-out infinite;
	}

	.ring {
		position: absolute;
		width: 100%;
		height: 100%;
		border: 2px solid rgba(192, 132, 252, 0.3);
		border-top-color: #fbbf24;
		border-radius: 50%;
		animation: spin 1s linear infinite;
	}

	.sparkles {
		position: absolute;
		top: -20px;
		right: -20px;
		font-size: 1.5rem;
		animation: float 2s ease-in-out infinite;
	}

	.generating-text {
		font-size: 1.25rem;
		font-weight: 500;
		background: linear-gradient(90deg, #c084fc, #f472b6, #fbbf24);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		animation: shimmer 2s linear infinite;
		background-size: 200% auto;
	}

	.prediction-wrapper {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		column-gap: 0.35em;
		row-gap: 0.1em;
		font-size: clamp(2rem, 5vw, 3.5rem);
		font-weight: 800;
		line-height: 1.2;
		padding: 0.5rem;
		perspective: 1000px;
	}

	.word {
		display: inline-block;
		background: linear-gradient(135deg, #c084fc 0%, #f472b6 50%, #fbbf24 100%);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		color: transparent; /* Fallback */
		filter: drop-shadow(0 0 20px rgba(192, 132, 252, 0.4));
		transform-style: preserve-3d;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	@keyframes pulse {
		0%,
		100% {
			transform: scale(0.8);
			opacity: 0.8;
		}
		50% {
			transform: scale(1.1);
			opacity: 1;
		}
	}

	@keyframes float {
		0%,
		100% {
			transform: translateY(0);
		}
		50% {
			transform: translateY(-10px);
		}
	}

	@keyframes shimmer {
		to {
			background-position: 200% center;
		}
	}
</style>
