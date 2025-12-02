import { NextResponse } from 'next/server';
import { getEvent } from '@/services/historical-event.service';

export async function GET({
  params,
}: {
  params: Promise<{ eventId: string }>;
}) {
  try {
    const { eventId } = await params;
    const res = await getEvent(eventId);

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
