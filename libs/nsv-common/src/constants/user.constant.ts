export const USER_EVENT = {
  REGISTERED: 'user.registered',
  DELETED: 'user.deleted',
  UPDATED: 'user.updated',
} as const;

export const USER = {
  STATUS: {
    ACTIVE: 'ACTIVE',
    PENDING: 'PENDING',
    DELETED: 'DELETED',
  },
  SEX: {
    MALE: 'MALE',
    FEMALE: 'FEMALE',
    OTHER: 'OTHER',
  },
} as const;
