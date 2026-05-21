import type { AuthPluginBase } from '@better-auth-ui/core';
import type { ComponentType, ReactNode } from 'react';
import type { AuthProps } from './auth';

type AuthButtonProps = {
  view: 'signIn' | 'signUp';
};

type AppAuthPlugin = AuthPluginBase & {
  captchaComponent?: ReactNode;
  authButtons?: ComponentType<AuthButtonProps>[];
  views?: { auth?: Record<string, ComponentType<AuthProps>> };
  fallbackViews?: { auth?: { signIn?: ComponentType<AuthProps> } };
  viewPaths?: { auth?: Record<string, string>; settings?: Record<string, string> };
  accountCards?: ComponentType[];
  securityCards?: ComponentType[];
  userMenuItems?: ComponentType[];
};

declare module '@better-auth-ui/core' {
  interface AuthPluginRegister {
    nsv: AppAuthPlugin;
  }
}
