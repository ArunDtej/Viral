<script lang="ts">
	import { onMount } from 'svelte';

	let { data } = $props<{ data: { id: string } }>();
	let apiData = $state<any>(null);
	let isLoading = $state(true);
	let error = $state<string | null>(null);

	onMount(async () => {
		try {
			// Using JSONPlaceholder as a dummy API
			const response = await fetch(`https://jsonplaceholder.typicode.com/posts/${data.id}`);

			if (!response.ok) {
				throw new Error(`API call failed with status: ${response.status}`);
			}

			apiData = await response.json();
		} catch (e: any) {
			error = e.message;
		} finally {
			isLoading = false;
		}
	});
</script>

<div class="container">
	<h1>Prediction Details</h1>
	<p>
		Fetching details for prediction ID: <strong>{data.id}</strong>
	</p>

	{#if isLoading}
		<p>Loading data from dummy API...</p>
	{:else if error}
		<div class="error">
			<p>Error fetching data:</p>
			<pre>{error}</pre>
		</div>
	{:else if apiData}
		<div class="api-data">
			<h2>{apiData.title}</h2>
			<p>{apiData.body}</p>
		</div>
	{/if}
</div>

<style>
	.container {
		max-width: 600px;
		margin: 2rem auto;
		padding: 2rem;
		border: 1px solid #eaeaea;
		border-radius: 8px;
	}
	.error {
		color: #d8000c;
		background-color: #ffbaba;
		padding: 1rem;
		border-radius: 4px;
	}
	.api-data {
		margin-top: 2rem;
		padding: 1rem;
		background-color: #f9f9f9;
		border: 1px solid #eee;
		border-radius: 4px;
	}
	.api-data h2 {
		margin-top: 0;
	}
</style>
