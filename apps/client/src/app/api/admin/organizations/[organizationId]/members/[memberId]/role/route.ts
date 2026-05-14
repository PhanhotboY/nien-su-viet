import { NextRequest, NextResponse } from 'next/server';
import { headers } from 'next/headers';
import { retryFetcher } from '@/services';
import { revalidateTag } from 'next/cache';

export async function PUT(
  request: NextRequest,
  { params }: { params: Promise<{ organizationId: string; memberId: string }> },
) {
  try {
    const { organizationId, memberId } = await params;
    const body = await request.json();
    const { role } = body;

    if (!role) {
      return NextResponse.json(
        { message: 'Role is required' },
        { status: 400 },
      );
    }

    const reqHeaders = new Headers(await headers());
    reqHeaders.delete('content-length');
    reqHeaders.set('content-type', 'application/json');

    // Call your backend API to update member role
    // Note: You may need to adjust this endpoint based on your backend implementation
    const result = await retryFetcher(
      `/auth/admin/organizations/${organizationId}/members/${memberId}/role`,
      {
        method: 'PUT',
        headers: reqHeaders,
        body: JSON.stringify({ role }),
      },
    );

    revalidateTag('organization-member', 'max');

    return NextResponse.json(result, { status: 200 });
  } catch (error) {
    console.error('Failed to update member role:', error);
    return NextResponse.json(
      { message: 'Failed to update member role' },
      { status: 500 },
    );
  }
}
