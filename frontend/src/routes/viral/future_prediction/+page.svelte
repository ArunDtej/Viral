<div class="page-container">
	<div class="form-container animate-slide-in-up">
		<div class="title-container animate-fade-in">
			<h1>My Future Prediction</h1>
			<p>Enter your details below to see what the future holds.</p>
		</div>

		<form on:submit|preventDefault={() => alert('Calculating your future...')}>
			<div class="form-group animate-fade-in anim-delay-100">
				<label for="name">Full Name</label>
				<input type="text" id="name" name="name" placeholder="e.g., Jane Doe" required />
			</div>

			<div class="form-group animate-fade-in anim-delay-200">
				<label for="dob">Date of Birth</label>
				<input type="date" id="dob" name="dob" required />
			</div>

			<div class="form-group animate-fade-in anim-delay-300">
				<label for="gender">Gender</label>
				<select id="gender" name="gender" required>
					<option value="" disabled selected>Select your gender</option>
					<option value="female">Female</option>
					<option value="male">Male</option>
					<option value="non-binary">Non-binary</option>
					<option value="other">Other</option>
					<option value="prefer-not-to-say">Prefer not to say</option>
				</select>
			</div>

			<button type="submit" class="submit-btn animate-fade-in anim-delay-300">
				Predict My Future
			</button>
		</form>
	</div>
</div>

<style>
	.form-container {
		width: 100%;
		max-width: 450px;
		background-color: var(--surface-color);
		padding: 3rem;
		border-radius: 16px;
		border: 1px solid var(--border-color);
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
		/* Make sure opacity is 0 to start for animations */
		opacity: 0;
	}

	.title-container {
		text-align: center;
		margin-bottom: 2.5rem;
		opacity: 0;
	}

	h1 {
		font-size: 2.25rem;
		font-weight: 700;
		color: #f9fafb;
		margin: 0;
	}

	.title-container p {
		margin-top: 0.5rem;
		font-size: 1rem;
		color: var(--text-muted-color);
	}

	.form-group {
		margin-bottom: 1.5rem;
		opacity: 0;
	}

	label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: 500;
		font-size: 0.875rem;
		color: #d1d5db;
	}

	input,
	select {
		width: 100%;
		padding: 0.875rem 1rem;
		border: 1px solid var(--border-color);
		border-radius: 8px;
		background-color: #334155; /* Slightly lighter input background */
		color: var(--text-color);
		font-size: 1rem;
		transition:
			border-color 0.3s ease,
			box-shadow 0.3s ease;
	}

	input::placeholder {
		color: #9ca3af;
	}

	input:focus,
	select:focus {
		outline: none;
		border-color: var(--primary-color);
		box-shadow: 0 0 0 4px rgba(139, 92, 246, 0.2);
	}

	/* Style for date input placeholder */
	input[type='date']:invalid {
		color: #9ca3af;
	}

	.submit-btn {
		display: block;
		width: 100%;
		padding: 1rem;
		border: none;
		border-radius: 8px;
		background-image: linear-gradient(to right, var(--primary-color), var(--primary-hover));
		color: white;
		font-size: 1rem;
		font-weight: 500;
		cursor: pointer;
		transition:
			transform 0.2s ease-out,
			box-shadow 0.2s ease-out;
		margin-top: 2rem;
		opacity: 0;
	}

	.submit-btn:hover {
		transform: translateY(-3px) scale(1.02);
		box-shadow: 0 7px 20px rgba(139, 92, 246, 0.25);
	}
</style>
