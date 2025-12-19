import { NextRequest, NextResponse } from 'next/server';
import { getEvents } from '@/services/historical-event.service';

export async function GET(request: NextRequest) {
  try {
    const { searchParams } = new URL(request.nextUrl);
    const res = await getEvents(searchParams.toString());

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
