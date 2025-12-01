<script lang="ts">
  // You can add Svelte logic here later for interactivity
</script>

<div class="page-container relative z-10">
	<div class="flex flex-col items-center justify-center min-h-[calc(100vh-theme(spacing.16))] p-4 sm:p-6 lg:p-8">
		<h1 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold text-gray-900 dark:text-white text-center mb-4 sm:mb-6 leading-tight">
			Predict the Next <span class="text-blue-600">Viral Sensation!</span>
		</h1>
		<p class="text-lg sm:text-xl text-gray-600 dark:text-gray-300 text-center max-w-2xl mb-8 sm:mb-10">
			Unleash your inner trend-spotter. What's going to break the internet next?
		</p>

		<div class="w-full max-w-md bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6 sm:p-8 border border-gray-200 dark:border-gray-700">
			<div class="mb-6">
				<label for="prediction-input" class="block text-gray-700 dark:text-gray-200 text-sm font-medium mb-2">
					Enter your idea, topic, or keyword:
				</label>
				<input
					type="text"
					id="prediction-input"
					placeholder="e.g., 'AI-powered pet feeders', 'retro gaming comeback'"
					class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400"
				/>
			</div>

			<button
				class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded-lg transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105 shadow-md"
			>
				Get My Prediction!
			</button>

			<div class="mt-8 text-center text-gray-500 dark:text-gray-400">
				<p>✨ Powered by cutting-edge trend analysis ✨</p>
			</div>
		</div>
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
</style>
