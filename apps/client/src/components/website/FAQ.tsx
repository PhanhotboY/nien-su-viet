'use client';

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import Link from '@/i18n/navigation';
import { useTranslations } from 'next-intl';

interface FAQItem {
  question: string;
  answer: string;
  value: string;
}

export const FAQ = ({ locale }: { locale: string }) => {
  const t = useTranslations('AboutPage.FAQ');

  const FAQList: FAQItem[] = [
    {
      question: t('items.item-1.question'),
      answer: t('items.item-1.answer'),
      value: 'item-1',
    },
    {
      question: t('items.item-2.question'),
      answer: t('items.item-2.answer'),
      value: 'item-2',
    },
    {
      question: t('items.item-3.question'),
      answer: t('items.item-3.answer'),
      value: 'item-3',
    },
    {
      question: t('items.item-4.question'),
      answer: t('items.item-4.answer'),
      value: 'item-4',
    },
    {
      question: t('items.item-5.question'),
      answer: t('items.item-5.answer'),
      value: 'item-5',
    },
  ];

  return (
    <section id="faq" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold mb-4">
        {t('title.prefix')}{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {t('title.highlight')}
        </span>
      </h2>

      <Accordion type="single" collapsible className="w-full AccordionRoot">
        {FAQList.map(({ question, answer, value }: FAQItem) => (
          <AccordionItem key={value} value={value} className="!border-b-1">
            <AccordionTrigger className="text-left text-base">
              {question}
            </AccordionTrigger>

            <AccordionContent className="text-base">{answer}</AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>

      <h3 className="font-medium mt-4">
        {t('footer.prefix')}{' '}
        <Link
          rel="noreferrer noopener"
          href={t('footer.href')}
          className="text-secondary transition-all border-primary hover:border-b-2"
        >
          {t('footer.linkText')}
        </Link>
      </h3>
    </section>
  );
};
