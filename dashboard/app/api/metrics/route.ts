import { NextResponse } from 'next/server';

export async function GET() {
    try {
        const response = await fetch('http://localhost:2112/metrics', {
            cache: 'no-store',
        });

        if (!response.ok) {
            // Fallback for demo purposes if the server isn't running
            return NextResponse.json({ tasksCompleted: 0, activeWorkers: 5 });
        }

        const text = await response.text();

        // Parse "tasks_completed_total 8910"
        const match = text.match(/tasks_completed_total\s+(\d+)/);
        const tasksCompleted = match ? parseInt(match[1], 10) : 0;

        return NextResponse.json({
            tasksCompleted,
            activeWorkers: 5
        });
    } catch (error) {
        console.error("Proxy fetch error:", error);
        return NextResponse.json({ tasksCompleted: 0, activeWorkers: 5 }, { status: 500 });
    }
}
