<script lang="ts">
	import { onMount } from 'svelte';

	onMount(() => {
		const script = document.createElement('script');
		script.src = '//pl28197650.effectivegatecpm.com/c91ff27a6caf5934b29d68567b330bf8/invoke.js';
		script.async = true;
		script.dataset.cfasync = 'false';
		document.body.appendChild(script);

		return () => {
			// Cleanup logic if needed, but usually ad scripts modify global state
			if (document.body.contains(script)) {
				document.body.removeChild(script);
			}
		};
	});
</script>

<!-- The specific container ID required by the script -->
<div id="container-c91ff27a6caf5934b29d68567b330bf8" class="flex justify-center my-4"></div>
