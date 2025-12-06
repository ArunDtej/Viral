<script lang="ts">
	import { onMount } from 'svelte';

	let container: HTMLDivElement;

	onMount(() => {
		if (container) {
			const conf = `
        <script type="text/javascript">
          atOptions = {
            'key' : '65525e6525529f450d27d10527d2c8ff',
            'format' : 'iframe',
            'height' : 50,
            'width' : 320,
            'params' : {}
          };
        <\/script>
        <script type="text/javascript" src="//www.highperformanceformat.com/65525e6525529f450d27d10527d2c8ff/invoke.js"><\/script>
      `;

			const iframe = document.createElement('iframe');
			iframe.title = 'Adsterra Banner';
			iframe.width = '320';
			iframe.height = '50';
			iframe.style.border = 'none';
			iframe.style.overflow = 'hidden';
			iframe.style.display = 'block';
			iframe.style.margin = '0 auto';

			container.appendChild(iframe);

			const doc = iframe.contentWindow?.document;
			if (doc) {
				doc.open();
				doc.write(
					'<html><head><base target="_top" /></head><body style="margin:0;padding:0;">' +
						conf +
						'</body></html>'
				);
				doc.close();
			}
		}
	});
</script>

<div bind:this={container} class="ad-container flex justify-center my-4" style="min-height: 50px;">
	<!-- Ad will be injected here -->
</div>
