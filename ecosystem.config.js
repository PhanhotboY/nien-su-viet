module.exports = {
  apps: [
    {
      name: 'auth-service',
      script: './dist/apps/auth/main.js', // Path to the compiled JavaScript file
      instances: 'max', // Scales app to the number of CPU cores
      exec_mode: 'cluster', // Enables clustering mode
      watch: false, // Disable watching for production
      env: {
        NODE_ENV: 'production', // Set the environment
      },
    },
    {
      name: 'historical-event-service',
      script: './dist/apps/historical-event/main.js', // Path to the compiled JavaScript file
      instances: 'max', // Scales app to the number of CPU cores
      exec_mode: 'cluster', // Enables clustering mode
      watch: false, // Disable watching for production
      env: {
        NODE_ENV: 'production', // Set the environment
      },
    },
    {
      name: 'gateway-service',
      script: './dist/apps/gateway/main.js', // Path to the compiled JavaScript file
      instances: 'max', // Scales app to the number of CPU cores
      exec_mode: 'cluster', // Enables clustering mode
      watch: false, // Disable watching for production
      env: {
        NODE_ENV: 'production', // Set the environment
      },
    },
  ],
};
