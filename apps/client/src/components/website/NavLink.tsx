import { components } from '@nsv-interfaces/cms-service';
import { ExternalLink } from 'lucide-react';
import Link from 'next/link';

export default function NavLink({
  navItem,
  ...props
}: Partial<Parameters<typeof Link>[0]> & {
  navItem: components['schemas']['HeaderNavItemData'];
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
