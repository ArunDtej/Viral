<script lang="ts">
	import './layout.css';
	import '../lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';
	import StarryBackground from '$lib/components/StarryBackground.svelte';
	import AdsterraBanner from '$lib/components/AdsterraBanner.svelte';
	import AdsterraBottom from '$lib/components/AdsterraBottom.svelte';
	import { predictionList } from '$lib/predictions';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="bg-black fixed inset-0 z-0">
	<StarryBackground />
</div>

<div class="relative z-10 min-h-screen flex flex-col">
	<AdsterraBanner />
	{@render children()}
	<AdsterraBottom />
</div>

<div style="display: none;" aria-hidden="true">
	<ul>
		{#each predictionList as prediction}
			<li><a href="/viral/{prediction.slug}">{prediction.title}</a></li>
		{/each}
		<li><a href="https://apigate.in">ApiGate</a></li>
		<li><a href="https://necrobyte.in">Necrobyte</a></li>
		<li><a href="https://necrobyte.in">Viral Necrobyte</a></li>
	</ul>
</div>
