import { genConfiguration } from '@phanhotboy/nsv-common';
import type { Default, Production } from './config.interface';

export const configuration = (
  envPath: string = './apps/historical-event/.env',
) => {
  require('dotenv').config({ path: envPath });

  const defaultConfig = require('./envs/default').config;
  const envConfigs: Record<string, Production & Default> = {
    development: require('./envs/development').config,
    production: require('./envs/production').config,
    test: require('./envs/test').config,
  };

  return genConfiguration<Default, Production>(defaultConfig, envConfigs);
};
