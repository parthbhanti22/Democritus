'use client';

import { useMemo, useRef } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { OrbitControls, Points, PointMaterial } from '@react-three/drei';
import * as THREE from 'three';

interface ParticleCloudProps {
    count?: number;
}

function Particles({ count = 5000 }: { count: number }) {
    const points = useRef<THREE.Points>(null!);

    // Generate Gaussian distributed points
    const positions = useMemo(() => {
        const data = new Float32Array(count * 3);
        for (let i = 0; i < count; i++) {
            // Gaussian approximation: (r1 + r2 + r3 + r4 - 2) * scale
            // A simple way to get a clump in the middle is just summing randoms.
            const x = (Math.random() - 0.5) * 10;
            const y = (Math.random() - 0.5) * 10;
            const z = (Math.random() - 0.5) * 10;

            // We want strict Gaussian? 
            // Let's use a simpler "Galaxy" distribution
            const theta = Math.random() * Math.PI * 2;
            const phi = Math.acos(2 * Math.random() - 1);
            const r = Math.pow(Math.random(), 0.3) * 10; // Cluster near center

            data[i * 3] = r * Math.sin(phi) * Math.cos(theta);
            data[i * 3 + 1] = r * Math.sin(phi) * Math.sin(theta);
            data[i * 3 + 2] = r * Math.cos(phi);
        }
        return data;
    }, [count]);

    useFrame((state, delta) => {
        // Optional: Subtle rotation or pulsation here if desired manually
        // points.current.rotation.x -= delta / 10;
        // points.current.rotation.y -= delta / 15;
    });

    return (
        <group rotation={[0, 0, Math.PI / 4]}>
            <Points ref={points} positions={positions} stride={3} frustumCulled={false}>
                <PointMaterial
                    transparent
                    color="#8b5cf6" // Violet-500
                    size={0.05}
                    sizeAttenuation={true}
                    depthWrite={false}
                    blending={THREE.AdditiveBlending}
                />
            </Points>
        </group>
    );
}

export default function ParticleCloud({ count = 10000 }: ParticleCloudProps) {
    return (
        <div className="absolute inset-0 z-0">
            <Canvas camera={{ position: [0, 0, 15], fov: 60 }}>
                <color attach="background" args={['#000000']} />
                <Particles count={count} />
                <OrbitControls autoRotate autoRotateSpeed={0.5} enableZoom={false} />
            </Canvas>
        </div>
    );
}
