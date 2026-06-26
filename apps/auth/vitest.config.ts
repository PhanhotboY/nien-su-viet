import { defineConfig } from 'vitest/config';
import path from 'path';

export default defineConfig({
  resolve: {
    tsconfigPaths: true,
  },
  test: {
    // Environment
    environment: 'node',
    // Include below if you want code coverage
    coverage: {
      provider: 'v8', // or 'istanbul'
      reporter: ['text', 'json', 'html'],
    },
    root: path.resolve(__dirname, '../../'),
  },
});
