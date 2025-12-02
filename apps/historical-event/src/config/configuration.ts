require('dotenv').config({ path: 'apps/historical-event/.env' });
import { commonUtils } from '@phanhotboy/nsv-common';
import type { Default, Production } from './config.interface';
import { config as defaultConfig } from './envs/default';
import { config as developmentConfig } from './envs/development';
import { config as productionConfig } from './envs/production';
import { config as testConfig } from './envs/test';

export const configuration = commonUtils.genConfiguration<Default, Production>(
  defaultConfig,
  {
    development: developmentConfig as Production,
    production: productionConfig,
    test: testConfig as Production,
  },
);
