<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import type { PageData } from './$types';

	export let data: PageData;
	let name: string = '';
	let loading = false;
	let mounted = false;
	let errorMessage = '';

	onMount(() => {
		mounted = true;
	});

	async function handleSubmit() {
		if (!mounted) return;

		// Clear previous errors
		errorMessage = '';
		loading = true;

		const minLoadingTime = 2000; // 2 seconds for better UX
		const startTime = Date.now();

		try {
			const response = await fetch('/api/predict', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					name,
					slug: 'your-2025-fortune-reading'
				})
			});

			const elapsedTime = Date.now() - startTime;
			const remainingTime = minLoadingTime - elapsedTime;

			if (remainingTime > 0) {
				await new Promise((resolve) => setTimeout(resolve, remainingTime));
			}

			if (response.ok) {
				const responseData = await response.json();
				// Navigate to the prediction result page
				goto(`/viral/your-2025-fortune-reading/${responseData.id}`);
			} else {
				const errorData = await response.json().catch(() => ({ message: 'An error occurred' }));
				errorMessage = errorData.message || 'Failed to generate prediction. Please try again.';
			}
		} catch (error) {
			console.error('Error submitting prediction:', error);
			errorMessage = 'Network error. Please check your connection and try again.';
		} finally {
			loading = false;
		}
	}
</script>

<div class="page-container relative z-10">
	<div class="form-container" class:animate-slide-in-up={mounted}>
		<div class="title-container" class:animate-fade-in={mounted}>
			<h1>{data.prediction.title}</h1>
			<p>Enter your details below to see what the future holds.</p>
		</div>

		<form on:submit|preventDefault={handleSubmit}>
			<div class="form-group" class:animate-fade-in={mounted} style="--anim-delay: 100ms">
				<label for="name">Full Name</label>
				<input
					type="text"
					id="name"
					name="name"
					placeholder="e.g., Jane Doe"
					required
					bind:value={name}
				/>
			</div>

			<button
				type="submit"
				class="submit-btn"
				class:animate-fade-in={mounted}
				style="--anim-delay: 300ms"
				disabled={loading}
			>
				{#if loading}
					<span>Calculating...</span>
				{:else}
					Predict My Future
				{/if}
			</button>

			{#if errorMessage}
				<div class="error-message">
					<span class="error-icon">⚠️</span>
					<p>{errorMessage}</p>
				</div>
			{/if}
		</form>

		{#if loading}
			<div class="loading-overlay">
				<div class="fancy-spinner">
					<div class="dot1"></div>
					<div class="dot2"></div>
				</div>
				<p class="loading-text">Generating your future...</p>
			</div>
		{/if}
	</div>
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
		max-width: 500px;
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
		opacity: 0;
	}

	label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: 500;
		font-size: 0.875rem;
		color: #d1d5db;
	}

	input,
	select {
		width: 100%;
		padding: 0.875rem 1rem;
		border: 1px solid rgba(255, 255, 255, 0.15);
		border-radius: 12px;
		background: rgba(255, 255, 255, 0.05);
		backdrop-filter: blur(10px);
		-webkit-backdrop-filter: blur(10px);
		color: var(--text-color);
		font-size: 1rem;
		transition:
			border-color 0.3s ease,
			background 0.3s ease,
			box-shadow 0.3s ease;
	}

	input::placeholder {
		color: rgba(255, 255, 255, 0.4);
	}

	input:focus,
	select:focus {
		outline: none;
		border-color: rgba(168, 85, 247, 0.6);
		background: rgba(255, 255, 255, 0.08);
		box-shadow:
			0 0 0 4px rgba(168, 85, 247, 0.15),
			0 8px 16px rgba(168, 85, 247, 0.2);
	}

	/* Style for date input placeholder */
	input[type='date']:invalid {
		color: #9ca3af;
	}

	.submit-btn {
		position: relative;
		display: block;
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
		opacity: 0;
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

	.error-message {
		margin-top: 1.5rem;
		padding: 1rem 1.25rem;
		background: rgba(239, 68, 68, 0.1);
		border: 1px solid rgba(239, 68, 68, 0.3);
		border-radius: 12px;
		backdrop-filter: blur(10px);
		-webkit-backdrop-filter: blur(10px);
		display: flex;
		align-items: center;
		gap: 0.75rem;
		animation:
			shake 0.5s ease-in-out,
			fade-in 0.3s ease-out;
	}

	.error-icon {
		font-size: 1.25rem;
		flex-shrink: 0;
	}

	.error-message p {
		margin: 0;
		color: #fca5a5;
		font-size: 0.875rem;
		line-height: 1.5;
	}

	@keyframes shake {
		0%,
		100% {
			transform: translateX(0);
		}
		10%,
		30%,
		50%,
		70%,
		90% {
			transform: translateX(-5px);
		}
		20%,
		40%,
		60%,
		80% {
			transform: translateX(5px);
		}
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

	/* Loading Overlay Styles */
	.loading-overlay {
		position: absolute;
		inset: 0;
		background-color: rgba(0, 0, 0, 0.8); /* Darker overlay */
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		border-radius: 16px;
		z-index: 10;
		color: white;
		font-size: 1.2rem;
		font-weight: 500;
		backdrop-filter: blur(5px); /* Subtle blur effect */
	}

	.fancy-spinner {
		width: 50px;
		height: 50px;
		position: relative;
		animation: rotate 2s linear infinite;
		margin-bottom: 1.5rem;
	}

	.fancy-spinner .dot1,
	.fancy-spinner .dot2 {
		width: 60%;
		height: 60%;
		display: inline-block;
		position: absolute;
		top: 0;
		background-color: var(--primary-color);
		border-radius: 100%;
		animation: bounce 2s infinite ease-in-out;
	}

	.fancy-spinner .dot2 {
		top: auto;
		bottom: 0;
		animation-delay: -1s;
		background-color: var(--primary-hover);
	}

	@keyframes rotate {
		100% {
			transform: rotate(360deg);
		}
	}

	@keyframes bounce {
		0%,
		100% {
			transform: scale(0);
		}
		50% {
			transform: scale(1);
		}
	}

	.loading-text {
		font-size: 1.5rem;
		font-weight: 600;
		color: #fff;
		animation: pulse-text 2s infinite ease-in-out;
	}

	@keyframes pulse-text {
		0%,
		100% {
			opacity: 1;
		}
		50% {
			opacity: 0.7;
		}
	}
</style>
