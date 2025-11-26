export const config = {
  db: {
    url: process.env.DATABASE_URL || '',
    directUrl: process.env.DIRECT_DATABASE_URL || '',
    // entities: [`${__dirname}/../../entity/**/*.{js,ts}`],
    // subscribers: [`${__dirname}/../../subscriber/**/*.{js,ts}`],
    // migrations: [`${__dirname}/../../migration/**/*.{js,ts}`],
  },
};
