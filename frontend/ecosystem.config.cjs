module.exports = {
  apps: [
    {
      name: 'viral-frontend',
      cwd: '/home/g200561_arun/Viral/frontend',   // VERY IMPORTANT
      script: 'build/index.js',
      exec_mode: 'fork',
      instances: 1,
      env: {
        NODE_ENV: 'production',        // REQUIRED for SvelteKit
        PORT: 3001,
        HOST: '0.0.0.0',
        PUBLIC_API_URL: 'https://necrobyte.in',
        PRIVATE_API_URL: 'http://127.0.0.1:8080'
      }
    }
  ]
};
