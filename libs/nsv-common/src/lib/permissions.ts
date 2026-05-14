import { post } from 'axios';
import { createAccessControl } from 'better-auth/plugins/access';
import {
  defaultStatements,
  adminAc,
  userAc,
} from 'better-auth/plugins/admin/access';

const crudActions = ['create', 'read', 'update', 'delete'] as (
  | 'create'
  | 'read'
  | 'update'
  | 'delete'
)[];

const eventStatements = {
  historicalEvent: crudActions,
  eventCategory: crudActions,
  eventEdit: crudActions,
};

const postStatements = {
  ping: crudActions,
  post: crudActions,
};

const organizationStatements = {
  organization: crudActions,
  organizationMember: crudActions,
};

const statements = {
  // convert readonly arrays to mutable arrays
  user: [...defaultStatements.user],
  session: [...defaultStatements.session],
  ...eventStatements,
  ...postStatements,
  ...organizationStatements,
};

/**
 * @default
 * user: ["create", "list", "set-role", "ban", "impersonate", "delete", "set-password", "get", "update"]
 * session: ["list", "revoke", "delete"]
 */
const ac = createAccessControl(statements);

const admin = ac.newRole({
  ...eventStatements,
  ...postStatements,
  ...organizationStatements,
  ...(adminAc.statements as any),
});

const editor = ac.newRole({
  ...eventStatements,
  eventEdit: ['create', 'read', 'update'],
  ...postStatements,
  ping: [],
  ...(userAc.statements as any),
});

const user = ac.newRole({
  historicalEvent: ['read'],
  eventCategory: ['read'],
  eventEdit: ['read'],
  post: ['read'],
  ...userAc.statements,
} as unknown as any);

const resources = Object.keys(statements) as (keyof typeof statements)[];
const roles = { admin, user, editor } as const;

export { admin, editor, user, ac, roles, resources, statements };
