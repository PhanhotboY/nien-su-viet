import { NextRequest, NextResponse } from 'next/server';
import { getUsers } from '@/services/user.service';

export async function GET(request: NextRequest) {
  try {
    const searchParams = Object.fromEntries(
      request.nextUrl.searchParams.entries(),
    );

    const res = await getUsers(searchParams);

    return NextResponse.json(res);
  } catch (error: any) {
    return NextResponse.json(
      {
        message: error.message || 'Unexpected error occurred',
        statusCode: error.statusCode || 500,
      },
      { status: error.statusCode || 500 },
    );
  }
}
