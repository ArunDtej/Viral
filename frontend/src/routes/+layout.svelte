<script lang="ts">
	import './layout.css';
	import '../lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';
	import StarryBackground from '$lib/components/StarryBackground.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="bg-black fixed inset-0 z-0">
	<StarryBackground />
</div>

<div class="relative z-10 min-h-screen flex flex-col">
	{@render children()}
</div>
