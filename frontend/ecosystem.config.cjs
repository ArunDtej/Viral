module.exports = {
  apps: [
    {
      name: 'frontend',
      script: 'build/index.js',
      exec_mode: 'fork',
      env: {
        PORT: 3001,
        HOST: '0.0.0.0',
        PUBLIC_API_URL: 'https://viral.necrobyte.in',
        PRIVATE_API_URL: 'http://127.0.0.1:8080'
      }
    }
  ]
};
