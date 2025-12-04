<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Viral Predictions</title>
</svelte:head>

<div class="page-container relative z-10">
	<div class="container mx-auto px-4 py-8 text-center">
		<div class="header-section">
			<h1 class="main-title">
				Choose Your <span class="gradient-text">Magical</span> Prediction
			</h1>
			<p class="subtitle">
				Unveil your destiny! Select one of our mystical predictions below to see what the future
				holds.
			</p>
		</div>

		<ul class="cards-grid">
			{#each data.predictions as prediction, i}
				<li style="--card-index: {i}">
					<a href="/viral/{prediction.slug}" class="glass-card">
						<div class="card-glow"></div>
						<div class="card-content">
							<span class="icon">{prediction.icon}</span>
							<h2 class="card-title">{prediction.title}</h2>
							<div class="stars-container">
								{#each Array(prediction.stars) as _, i}
									<span class="star">⭐</span>
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
		padding: 2rem 1rem;
	}

	.header-section {
		margin-bottom: 3rem;
		animation: fadeInDown 0.8s cubic-bezier(0.16, 1, 0.3, 1);
	}

	.main-title {
		font-size: clamp(2rem, 5vw, 3.5rem);
		font-weight: 800;
		color: #ffffff;
		margin-bottom: 1rem;
		line-height: 1.2;
	}

	.gradient-text {
		background: linear-gradient(135deg, #a855f7 0%, #ec4899 50%, #f59e0b 100%);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		animation: shimmer 3s ease-in-out infinite;
		background-size: 200% 200%;
	}

	@keyframes shimmer {
		0%,
		100% {
			background-position: 0% 50%;
		}
		50% {
			background-position: 100% 50%;
		}
	}

	.subtitle {
		font-size: 1.125rem;
		color: rgba(255, 255, 255, 0.7);
		max-width: 600px;
		margin: 0 auto;
		line-height: 1.6;
	}

	.cards-grid {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.cards-grid li {
		width: 100%;
		max-width: 450px;
		opacity: 0;
		animation: slideUp 0.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
		animation-delay: calc(0.1s * var(--card-index));
	}

	.glass-card {
		position: relative;
		display: block;
		background: rgba(255, 255, 255, 0.03);
		backdrop-filter: blur(20px);
		-webkit-backdrop-filter: blur(20px);
		border-radius: 24px;
		border: 1px solid rgba(255, 255, 255, 0.1);
		overflow: hidden;
		transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
		text-decoration: none;
	}

	.glass-card::before {
		content: '';
		position: absolute;
		inset: 0;
		border-radius: 24px;
		padding: 2px;
		background: linear-gradient(
			135deg,
			rgba(168, 85, 247, 0.4) 0%,
			rgba(236, 72, 153, 0.4) 50%,
			rgba(245, 158, 11, 0.4) 100%
		);
		-webkit-mask:
			linear-gradient(#fff 0 0) content-box,
			linear-gradient(#fff 0 0);
		-webkit-mask-composite: xor;
		mask-composite: exclude;
		opacity: 0;
		transition: opacity 0.4s ease;
	}

	.glass-card:hover::before {
		opacity: 1;
	}

	.card-glow {
		position: absolute;
		inset: -100px;
		background: radial-gradient(circle at center, rgba(168, 85, 247, 0.15) 0%, transparent 70%);
		opacity: 0;
		transition: opacity 0.4s ease;
	}

	.glass-card:hover .card-glow {
		opacity: 1;
	}

	.glass-card:hover {
		transform: translateY(-8px) scale(1.02);
		background: rgba(255, 255, 255, 0.05);
		box-shadow:
			0 20px 60px rgba(168, 85, 247, 0.3),
			0 0 80px rgba(168, 85, 247, 0.1);
	}

	.card-content {
		position: relative;
		z-index: 1;
		padding: 3rem 2rem;
		text-align: center;
	}

	.icon {
		display: block;
		font-size: 4rem;
		margin-bottom: 1.5rem;
		filter: drop-shadow(0 4px 12px rgba(168, 85, 247, 0.4));
		animation: float 3s ease-in-out infinite;
	}

	@keyframes float {
		0%,
		100% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-10px);
		}
	}

	.card-title {
		font-size: 1.5rem;
		font-weight: 700;
		color: #ffffff;
		margin: 0 0 1rem 0;
		text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
	}

	.stars-container {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.25rem;
	}

	.star {
		font-size: 1.25rem;
		filter: drop-shadow(0 2px 4px rgba(245, 158, 11, 0.5));
		animation: twinkle 2s ease-in-out infinite;
		animation-delay: calc(var(--star-index, 0) * 0.2s);
	}

	@keyframes twinkle {
		0%,
		100% {
			opacity: 1;
			transform: scale(1);
		}
		50% {
			opacity: 0.6;
			transform: scale(0.9);
		}
	}

	@keyframes fadeInDown {
		from {
			opacity: 0;
			transform: translateY(-30px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	@keyframes slideUp {
		from {
			opacity: 0;
			transform: translateY(40px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
</style>