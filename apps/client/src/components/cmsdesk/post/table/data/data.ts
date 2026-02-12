import {
  Pencil2Icon as DraftIcon,
  CheckCircledIcon as PublishedIcon,
} from '@radix-ui/react-icons';
import { Calendar } from 'lucide-react';

export const statuses = [
  {
    value: 'true',
    label: 'Published',
    icon: PublishedIcon,
  },
  {
    value: 'false',
    label: 'Draft',
    icon: DraftIcon,
  },
];

export const categories = [
  {
    value: '4db30a13-2797-4c7d-a0ce-e0c127287a39',
    label: 'History',
    icon: Calendar,
  },
];
