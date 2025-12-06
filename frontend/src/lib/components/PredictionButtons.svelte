<script lang="ts">
	import { RefreshCw, Home } from 'lucide-svelte';
	export let slug: string;
</script>

<div class="button-group">
	<a href="/viral/{slug}" class="secondary-btn">
		<RefreshCw size={20} />
		<span>Try Again</span>
	</a>
	<a href="/" class="secondary-btn">
		<Home size={20} />
		<span>View More Predictions</span>
	</a>
</div>

<style>
	.button-group {
		display: flex;
		gap: 1rem;
		margin-top: 1.5rem;
	}

	.secondary-btn {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.75rem;
		width: 100%;
		padding: 0.875rem;
		border: 1px solid rgba(255, 255, 255, 0.2);
		border-radius: 12px;
		background: rgba(255, 255, 255, 0.05);
		color: white;
		font-size: 1rem;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.3s ease;
		text-decoration: none;
	}

	.secondary-btn:hover {
		background: rgba(255, 255, 255, 0.1);
		border-color: rgba(255, 255, 255, 0.3);
	}
</style>
