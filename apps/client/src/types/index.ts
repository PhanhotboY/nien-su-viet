export type SeoType = {
  title: string;
  subTitle: string;
  absoluteTitle: string;
  ogTitle: string;
  author: { name: string; twitterUrl: string; twitterAddress: string };
  description: string;
  keywords: Array<string>;
  tags: Array<string>;
};
