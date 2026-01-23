'use client';

import { useEffect, useState } from 'react';
import ParticleCloud from '@/components/ParticleCloud';
import { fetchMetrics } from '@/lib/api';
import { Activity, Server, Atom } from 'lucide-react';

export default function Home() {
  const [metrics, setMetrics] = useState({ tasksCompleted: 0, activeWorkers: 0 });

  useEffect(() => {
    // Initial fetch
    fetchMetrics().then(setMetrics);

    // Polling interval
    const interval = setInterval(() => {
      fetchMetrics().then(setMetrics);
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  return (
    <main className="relative w-screen h-screen overflow-hidden bg-black text-white">
      {/* 3D Background */}
      <ParticleCloud count={10000 + metrics.tasksCompleted} />

      {/* Glassmorphism Overlay */}
      <div className="absolute top-8 left-8 z-10">
        <div className="backdrop-blur-xl bg-white/10 border border-white/20 p-6 rounded-2xl shadow-2xl w-80">
          <div className="flex items-center gap-3 mb-6">
            <div className="relative">
              <div className="w-3 h-3 bg-green-500 rounded-full animate-ping absolute inset-0"></div>
              <div className="w-3 h-3 bg-green-500 rounded-full relative"></div>
            </div>
            <h1 className="text-xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-purple-400">
              Democritus Grid
            </h1>
          </div>

          <div className="space-y-4">
            {/* Metric 1: Tasks */}
            <div className="flex items-center justify-between p-3 rounded-lg bg-black/20 hover:bg-black/30 transition-colors">
              <div className="flex items-center gap-2 text-gray-300">
                <Atom size={18} />
                <span className="text-sm font-medium">Tasks Processed</span>
              </div>
              <span className="text-2xl font-mono font-bold text-blue-400">
                {metrics.tasksCompleted.toLocaleString()}
              </span>
            </div>

            {/* Metric 2: Workers */}
            <div className="flex items-center justify-between p-3 rounded-lg bg-black/20 hover:bg-black/30 transition-colors">
              <div className="flex items-center gap-2 text-gray-300">
                <Server size={18} />
                <span className="text-sm font-medium">Active Nodes</span>
              </div>
              <span className="text-2xl font-mono font-bold text-green-400">
                {metrics.activeWorkers}
              </span>
            </div>

            {/* Metric 3: Status */}
            <div className="flex items-center justify-between p-3 rounded-lg bg-black/20 hover:bg-black/30 transition-colors">
              <div className="flex items-center gap-2 text-gray-300">
                <Activity size={18} />
                <span className="text-sm font-medium">System Status</span>
              </div>
              <span className="text-xs font-bold px-2 py-1 rounded-full bg-green-500/20 text-green-300 border border-green-500/30">
                OPERATIONAL
              </span>
            </div>

          </div>

          <div className="mt-6 text-xs text-gray-500 text-center">
            Distributed Monte Carlo Engine v1.0
          </div>
        </div>
      </div>
    </main>
  );
}
