<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Viral Predictions</title>
</svelte:head>

<div class="page-container relative z-10">
	<div class="container mx-auto px-4 py-8 text-center bg-transparent">
		<h1 class="text-4xl font-extrabold text-gray-900 dark:text-white mb-4">Choose Your <span class="text-purple-600">Magical</span> Prediction</h1>
		<p class="text-lg text-gray-600 dark:text-gray-300 mb-8">Unveil your destiny! Select one of our mystical predictions below to see what the future holds.</p>
		<ul class="flex flex-col items-center gap-6">
			{#each data.predictions as prediction}
				<li>
					<a href="/viral/{prediction.slug}" class="card w-full max-w-md animate-card">
						<div class="card-header animate-header">
							<span class="text-5xl mb-4">{prediction.icon}</span>
							<h2 class="text-xl font-semibold text-gray-800 dark:text-white animate-text">{prediction.title}</h2>
							<div class="flex items-center justify-center mt-2">
								{#each Array(prediction.stars) as _, i}
									<span class="text-yellow-400 text-lg">⭐</span>
								{/each}
							</div>
						</div>
					</a>
				</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	.page-container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem;
		background-color: transparent; /* Ensure background is transparent */
	}

	/* Card styles adapted from detail page */
	.card {
		background: var(--surface-color);
		border-radius: 24px;
		overflow: hidden;
		border: 1px solid var(--border-color);
		box-shadow: 0 20px 50px rgba(0, 0, 0, 0.6);
		opacity: 0;
		transform: translateY(20px);
		transition:
			transform 0.2s ease-out,
			box-shadow 0.2s ease-out;
	}

	.card:hover {
		transform: translateY(-3px) scale(1.03);
		box-shadow: 0 8px 25px rgba(139, 92, 246, 0.3);
	}

	.card-header {
		padding: 3rem 3rem 2rem;
		text-align: center;
		background: rgba(255, 255, 255, 0.03);
		border-bottom: 1px solid var(--border-color);
		opacity: 0;
	}

	.animate-card {
		animation: slide-up 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
	}
	.animate-header {
		animation: fade-in 0.6s 0.4s ease-out forwards;
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