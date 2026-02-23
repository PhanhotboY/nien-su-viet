import { ExternalLink } from 'lucide-react';
import Link from 'next/link';

import { headerNavItems } from '@/content/menus/header-nav-items';

export default function NavLink({
  navItem,
  ...props
}: Partial<Parameters<typeof Link>[0]> & {
  navItem: (typeof headerNavItems)[number];
}) {
  return (
    <Link
      rel="noreferrer noopener"
      target={navItem.link_new_tab ? '__blank' : ''}
      {...props}
      href={navItem.link_url}
    >
      {navItem.link_label}
      {navItem.link_new_tab ? <ExternalLink /> : ''}
    </Link>
  );
}
