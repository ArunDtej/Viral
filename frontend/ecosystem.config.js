module.exports = {
  apps: [
    {
      name: 'frontend',
      script: '.svelte-kit/output/server/index.js',
      exec_mode: 'fork',
      env: {
        // Environment variables for your application
        PORT: 3000, // Change this to your desired port
        HOST: '0.0.0.0',
      },
    },
  ],
};
