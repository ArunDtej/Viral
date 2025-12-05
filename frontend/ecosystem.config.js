module.exports = {
  apps: [
    {
      name: 'frontend',
      script: 'build/index.js',
      exec_mode: 'fork',
      env: {
        PORT: 3001,
        HOST: '0.0.0.0'
      }
    }
  ]
};
