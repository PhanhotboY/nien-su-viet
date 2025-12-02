import { setupAdmin } from './admin';

setupAdmin()
  .catch((err) => {
    console.error('Error setting up admin user:', err);
  })
  .finally(() => {
    process.exit(0);
  });
