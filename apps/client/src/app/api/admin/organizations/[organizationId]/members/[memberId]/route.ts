import { NextRequest, NextResponse } from 'next/server';
import {
  addUserToOrganization,
  removeUserFromOrganization,
} from '@/services/member.service';

export async function DELETE(
  request: NextRequest,
  { params }: { params: { organizationId: string; memberId: string } },
) {
  try {
    const result = await removeUserFromOrganization({
      memberId: params.memberId,
      organizationId: params.organizationId,
    });

    return NextResponse.json(result, { status: 200 });
  } catch (error) {
    console.error('Failed to remove member:', error);
    return NextResponse.json(
      { message: 'Failed to remove member from organization' },
      { status: 500 },
    );
  }
}

export async function POST(
  request: NextRequest,
  { params }: { params: Promise<{ organizationId: string; memberId: string }> },
) {
  try {
    const { memberId, organizationId } = await params;
    const result = await addUserToOrganization({
      userId: memberId,
      organizationId,
    });

    return NextResponse.json(result, { status: 200 });
  } catch (error) {
    console.error('Failed to add member:', error);
    return NextResponse.json(
      { message: 'Failed to add member to organization' },
      { status: 500 },
    );
  }
}
