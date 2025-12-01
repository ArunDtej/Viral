<script lang="ts">
	import { onMount } from 'svelte';

	let canvas: HTMLCanvasElement;
	let ctx: CanvasRenderingContext2D | null;
	let animationFrameId: number;

	interface Star {
		x: number;
		y: number;
		size: number;
		vx: number;
		vy: number;
		brightness: number;
		twinkleSpeed: number;
		twinklePhase: number;
	}

	let stars: Star[] = [];

	function createStars(count: number) {
		stars = [];
		for (let i = 0; i < count; i++) {
			stars.push({
				x: Math.random() * canvas.width,
				y: Math.random() * canvas.height,
				size: Math.random() * 1.5 + 0.5, // Smaller, sharper stars (0.5 to 2 pixels)
				vx: Math.random() * 0.3 - 0.15,
				vy: Math.random() * 0.5 + 0.1,
				brightness: Math.random() * 0.3 + 0.7, // Higher base brightness (0.7 to 1)
				twinkleSpeed: Math.random() * 0.04 + 0.02,
				twinklePhase: Math.random() * Math.PI * 2
			});
		}
	}

	function resizeCanvas() {
		if (!canvas) return;
		canvas.width = window.innerWidth;
		canvas.height = window.innerHeight;
		createStars(200); // More stars for better coverage
	}

	function drawStar(star: Star) {
		if (!ctx) return;

		const currentBrightness = star.brightness + Math.sin(star.twinklePhase) * 0.3;

		// Draw a sharp cross/star shape for more realistic stars
		ctx.save();
		ctx.translate(star.x, star.y);

		// Subtle glow (much smaller than before)
		const glowSize = star.size * 2;
		const gradient = ctx.createRadialGradient(0, 0, 0, 0, 0, glowSize);
		gradient.addColorStop(0, `rgba(255, 255, 255, ${currentBrightness * 0.8})`);
		gradient.addColorStop(0.3, `rgba(220, 230, 255, ${currentBrightness * 0.3})`);
		gradient.addColorStop(1, 'rgba(200, 220, 255, 0)');

		ctx.fillStyle = gradient;
		ctx.beginPath();
		ctx.arc(0, 0, glowSize, 0, Math.PI * 2);
		ctx.fill();

		// Bright sharp center point
		ctx.fillStyle = `rgba(255, 255, 255, ${currentBrightness})`;
		ctx.fillRect(-star.size / 2, -star.size / 2, star.size, star.size);

		// Add cross sparkle effect for brighter stars
		if (star.size > 1) {
			ctx.strokeStyle = `rgba(255, 255, 255, ${currentBrightness * 0.6})`;
			ctx.lineWidth = 0.5;
			ctx.beginPath();
			ctx.moveTo(-star.size * 2, 0);
			ctx.lineTo(star.size * 2, 0);
			ctx.moveTo(0, -star.size * 2);
			ctx.lineTo(0, star.size * 2);
			ctx.stroke();
		}

		ctx.restore();
	}

	function updateStars() {
		stars.forEach((star) => {
			// Move the star
			star.x += star.vx;
			star.y += star.vy;

			// Twinkle effect
			star.twinklePhase += star.twinkleSpeed;

			// Wrap around screen
			if (star.x < -10) star.x = canvas.width + 10;
			if (star.x > canvas.width + 10) star.x = -10;
			if (star.y < -10) star.y = canvas.height + 10;
			if (star.y > canvas.height + 10) star.y = -10;
		});
	}

	function animate() {
		if (!ctx || !canvas) return;

		// Clear canvas completely for sharper rendering
		ctx.fillStyle = '#000000';
		ctx.fillRect(0, 0, canvas.width, canvas.height);

		updateStars();
		stars.forEach(drawStar);

		animationFrameId = requestAnimationFrame(animate);
	}

	onMount(() => {
		ctx = canvas.getContext('2d');
		resizeCanvas();

		// Initial clear
		if (ctx) {
			ctx.fillStyle = '#000000';
			ctx.fillRect(0, 0, canvas.width, canvas.height);
		}

		animate();

		window.addEventListener('resize', resizeCanvas);

		return () => {
			window.removeEventListener('resize', resizeCanvas);
			if (animationFrameId) {
				cancelAnimationFrame(animationFrameId);
			}
		};
	});
</script>

<canvas bind:this={canvas} class="starry-canvas"></canvas>

<style>
	.starry-canvas {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		z-index: 0;
		pointer-events: none;
	}
</style>
