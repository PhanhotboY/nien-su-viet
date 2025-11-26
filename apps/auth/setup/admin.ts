// Load environment variables BEFORE any imports
require('dotenv').config({ path: 'apps/auth/.env' });

export async function setupAdmin() {
  // Dynamic import after env is loaded
  const { auth } = await import('../src/lib/auth.js');
  const adminEmail = process.env.ADMIN_EMAIL;
  const adminName = process.env.ADMIN_NAME || 'Admin User';
  const adminPassword = process.env.ADMIN_PASSWORD;
  if (!adminEmail || !adminPassword) {
    return console.error(
      'ADMIN_EMAIL and ADMIN_PASSWORD must be set in environment variables.',
    );
  }

  const admin = await auth.api.createUser({
    body: {
      email: adminEmail,
      name: adminName,
      role: 'admin',
      password: adminPassword,
    },
  });

  console.log('Admin user created:', admin.user);
}
