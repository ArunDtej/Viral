<script lang="ts">
	import './layout.css';
	import '../lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="bg-black fixed inset-0 z-0">
    <div class="stars"></div>
    <div class="twinkling"></div>
</div>

<div class="relative z-10 min-h-screen flex flex-col">
	{@render children()}
</div>
