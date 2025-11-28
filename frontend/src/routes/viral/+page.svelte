<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Viral Predictions</title>
</svelte:head>

<div class="container">
	<h1>Choose Your Prediction</h1>
	<p>What do you want to know about your future? Select one of our super viral predictions below!</p>
	<ul class="prediction-list">
		{#each data.predictions as prediction}
			<li>
				<a href="/viral/{prediction.slug}">{prediction.title}</a>
			</li>
		{/each}
	</ul>
</div>

<style>
	.container {
		max-width: 600px;
		margin: 2rem auto;
		padding: 2rem;
		text-align: center;
	}

	h1 {
		margin-bottom: 1rem;
	}

	p {
		margin-bottom: 2rem;
		color: #666;
	}

	.prediction-list {
		list-style: none;
		padding: 0;
		display: grid;
		gap: 1rem;
	}

	.prediction-list a {
		display: block;
		padding: 1.5rem;
		background-color: #fff;
		border: 1px solid #eee;
		border-radius: 8px;
		text-decoration: none;
		color: #333;
		font-size: 1.1rem;
		font-weight: 500;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
		transition:
			transform 0.2s,
			box-shadow 0.2s;
	}

	.prediction-list a:hover {
		transform: translateY(-3px);
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	}
</style>
