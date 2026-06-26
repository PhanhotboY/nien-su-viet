import { genConfiguration } from '@phanhotboy/nsv-common';
import type { Default, Production } from './config.interface';
import { config as defaultConfig } from './envs/default';
import { config as developmentConfig } from './envs/development';
import { config as productionConfig } from './envs/production';
import { config as testConfig } from './envs/test';

export const configuration = (envPath: string = './apps/auth/.env') => {
  require('dotenv').config({ path: envPath });

  const defaultCfg: Default = defaultConfig();
  const envConfigs: Record<string, Production & Default> = {
    development: developmentConfig(),
    production: productionConfig(),
    test: testConfig(),
  };

  return genConfiguration(defaultCfg, envConfigs);
};
