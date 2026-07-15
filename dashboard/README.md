# Democritus Dashboard

A real-time 3D visualization dashboard for the Democritus distributed Monte Carlo simulation engine. Built with Next.js, React Three Fiber, and Tailwind CSS.

![Dashboard Preview](../Democritus.PNG)

## Features

- **3D Particle Cloud**: WebGL-powered visualization using Three.js/React Three Fiber
- **Real-time Metrics**: Live task completion counter from Prometheus
- **Auto-rotating Camera**: Orbital controls with smooth animation
- **Responsive Design**: Works on desktop and mobile browsers

## Tech Stack

- **Framework**: [Next.js](https://nextjs.org/) (App Router)
- **3D Graphics**: [React Three Fiber](https://docs.pmnd.rs/react-three-fiber) + [Three.js](https://threejs.org/)
- **Styling**: [Tailwind CSS](https://tailwindcss.com/)
- **Animation**: [Framer Motion](https://www.framer.com/motion/)
- **Language**: TypeScript

## Getting Started

### Prerequisites

- Node.js 18+ and npm
- The Democritus backend running (for live metrics)

### Installation

```bash
# Navigate to dashboard directory
cd dashboard

# Install dependencies
npm install

# Start development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) to view the dashboard.

## Project Structure

```
dashboard/
├── app/
│   ├── layout.tsx        # Root layout with metadata
│   ├── page.tsx          # Main dashboard page
│   ├── globals.css       # Global styles
│   └── api/
│       └── metrics/
│           └── route.ts  # Proxy endpoint for Prometheus metrics
├── components/
│   └── ParticleCloud.tsx # 3D WebGL particle visualization
├── lib/
│   └── api.ts            # API utilities
└── public/               # Static assets
```

## API Endpoint

### `GET /api/metrics`

Proxies metrics from the Democritus Scheduler's Prometheus endpoint.

**Response:**
```json
{
  "tasksCompleted": 8910,
  "activeWorkers": 5
}
```

**Notes:**
- Fetches from `http://localhost:2112/metrics` (Scheduler metrics endpoint)
- Returns fallback values if the backend is unavailable
- Parses the `tasks_completed_total` Prometheus metric

## Components

### ParticleCloud

The main visualization component renders a 3D point cloud representing simulation results.

**Props:**
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `count` | `number` | `10000` | Number of particles to render |

**Features:**
- Spherical distribution with center clustering
- Additive blending for glow effect
- Auto-rotating orbital camera
- Violet color theme (#8b5cf6)

## Development

```bash
# Run development server with hot reload
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Run linting
npm run lint
```

## Connecting to Backend

The dashboard expects the Democritus backend to be running:

```bash
# From project root
docker compose up --build --scale worker=5
```

This starts:
- Scheduler on port 50051 (gRPC) and 2112 (metrics)
- Prometheus on port 9090
- 5 Worker instances

## Customization

### Changing Particle Count
Edit the `count` prop in `app/page.tsx`:
```tsx
<ParticleCloud count={20000} />
```

### Changing Particle Color
Edit the `color` value on the `PointMaterial` in `components/ParticleCloud.tsx`:
```tsx
<PointMaterial color="#00ff00" />
```

## Troubleshooting

### "Failed to fetch metrics"
- Ensure the Democritus backend is running
- Check that port 2112 is accessible
- The dashboard will show fallback values (0 tasks) if backend is down

### Particles not rendering
- Ensure WebGL is enabled in your browser
- Try a different browser (Chrome/Firefox recommended)
- Check browser console for Three.js errors

## License

MIT License - see [LICENSE](../LICENSE) for details.
