'use client';

// import { Draft } from "@/types/collection";
import { ColumnDef } from '@tanstack/react-table';
import { format } from 'date-fns';
import { DataTableColumnHeader } from './data-table-column-header';
import { DataTableRowActions } from './data-table-row-actions';
import { statuses } from './data/data';
import { PostBriefResponseDto } from '@/types/collection';
import Link from '@/i18n/navigation';

export const columns: ColumnDef<PostBriefResponseDto>[] = [
  {
    accessorKey: 'thumbnail',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Thumbnail" />
    ),
    cell: ({ row }) => {
      const thumbnail: string = row.getValue('thumbnail');
      if (!thumbnail) {
        return null;
      }
      return (
        <div className="w-28 h-full items-center aspect-square">
          <img
            src={thumbnail}
            alt="Thumbnail"
            className="w-full h-full object-contain"
          />
        </div>
      );
    },
    enableHiding: false,
    enableSorting: false,
    enableColumnFilter: false,
    maxSize: 100,
  },
  {
    accessorKey: 'title',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Title" />
    ),
    cell: ({ row }) => {
      return (
        <Link
          href={`/cmsdesk/posts/${row.getValue('id')}`}
          className="flex space-x-2 hover:underline"
        >
          <div className="flex space-x-2">
            <span className="max-w-[500px] truncate font-medium">
              {row.getValue('title')}
            </span>
          </div>
        </Link>
      );
    },
    enableHiding: false,
  },
  {
    accessorKey: 'published',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Status" />
    ),
    cell: ({ row }) => {
      const status = statuses.find(
        (status) => status.value === row.getValue('published')?.toString(),
      );

      if (!status) {
        return null;
      }

      return (
        <div className="flex w-[100px] items-center">
          {status.icon && (
            <status.icon className="mr-2 h-4 w-4 text-muted-foreground" />
          )}
          <span>{status.label}</span>
        </div>
      );
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id)?.toString());
    },
  },
  {
    accessorKey: 'createdAt',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Created" />
    ),
    cell: ({ row }) => {
      const date = format(
        new Date(row.getValue('createdAt') || Date.now()),
        'MM/dd/yyyy',
      );

      if (!date) {
        return null;
      }

      return (
        <div className="flex w-[100px] items-center">
          <span>{date}</span>
        </div>
      );
    },
  },
  {
    accessorKey: 'id',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Actions" />
    ),
    cell: ({ row }) => <DataTableRowActions row={row} />,
    enableHiding: false,
    enableSorting: false,
  },
];
