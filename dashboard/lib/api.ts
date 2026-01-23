export async function fetchMetrics(): Promise<{ tasksCompleted: number; activeWorkers: number }> {
    try {
        const res = await fetch('/api/metrics');
        if (!res.ok) throw new Error('Failed to fetch metrics');
        return res.json();
    } catch (error) {
        console.error(error);
        return { tasksCompleted: 0, activeWorkers: 0 };
    }
}
