<script lang="ts">
	import type { Point } from '$lib/types'
	import WorldMapSVG from './world-map.svg?raw'

	const { origin, destination }: { origin: Point; destination: Point } = $props()

	let mapWrapper: HTMLDivElement

	const width = 960
	const height = 500

	/**
	 * Returns the [x, y] coordinates of the given longitude and latitude using a simple equirectangular
	 * projection.
	 */
	export function project(lon: number, lat: number): [number, number] {
		const x = (lon + 180) * (width / 360)
		const y = (90 - lat) * (height / 180)
		return [x, y]
	}

	function coordsToSVGPath(coords: number[][][]): string {
		let path = ''
		for (const ring of coords) {
			for (const [j, coord] of ring.entries()) {
				const [x, y] = project(coord[0], coord[1]).map((n) => n.toFixed(1))
				path += `${j === 0 ? 'M' : 'L'}${x},${y}`
			}
			path += 'Z'
		}
		return path
	}

	const svgPaths = []
	const omitAntarctica = false
	for (const feature of geojsonFeatures) {
		if (omitAntarctica && feature.properties?.GEOUNIT === 'Antarctica') {
			continue
		}
		const polygons =
			feature.geometry.type === 'Polygon'
				? [feature.geometry.coordinates]
				: feature.geometry.type === 'MultiPolygon'
					? feature.geometry.coordinates
					: []
		for (const polygon of polygons) {
			const path = coordsToSVGPath(polygon)
			svgPaths.push(
				`<path d="${path}" fill="var(--land-color)" stroke="var(--border-color)" stroke-width="0.5"/>`
			)
		}
	}

	const svgContent = `
	<svg width="${width}" height="${height}" viewBox="0 0 ${width} ${height}" xmlns="http://www.w3.org/2000/svg">
		${svgPaths.join('\n')}
	</svg>
`

	$effect(() => {
		const svg = mapWrapper?.querySelector('svg')
		if (svg) {
			const path = document.createElementNS('http://www.w3.org/2000/svg', 'path')
			const [x1, y1] = project(origin.longitude, origin.latitude)
			const [x2, y2] = project(destination.longitude, destination.latitude)

			// Add start and end points
			for (const [x, y] of [
				[x1, y1],
				[x2, y2]
			]) {
				const point = document.createElementNS('http://www.w3.org/2000/svg', 'circle')
				point.setAttribute('cx', x.toString())
				point.setAttribute('cy', y.toString())
				point.setAttribute('r', '2')
				point.setAttribute('fill', 'var(--map-point)')
				svg.appendChild(point)
			}

			// Calculate intermediate points for great circle route.
			const numPoints = 100
			let pathData = [`M ${x1} ${y1}`]
			for (let i = 1; i <= numPoints; i++) {
				const f = i / numPoints
				const lat1 = (origin.latitude * Math.PI) / 180
				const lon1 = (origin.longitude * Math.PI) / 180
				const lat2 = (destination.latitude * Math.PI) / 180
				const lon2 = (destination.longitude * Math.PI) / 180

				const d = Math.acos(
					Math.sin(lat1) * Math.sin(lat2) + Math.cos(lat1) * Math.cos(lat2) * Math.cos(lon2 - lon1)
				)
				const A = Math.sin((1 - f) * d) / Math.sin(d)
				const B = Math.sin(f * d) / Math.sin(d)

				const x = A * Math.cos(lat1) * Math.cos(lon1) + B * Math.cos(lat2) * Math.cos(lon2)
				const y = A * Math.cos(lat1) * Math.sin(lon1) + B * Math.cos(lat2) * Math.sin(lon2)
				const z = A * Math.sin(lat1) + B * Math.sin(lat2)

				const lat = (Math.atan2(z, Math.sqrt(x * x + y * y)) * 180) / Math.PI
				const lon = (Math.atan2(y, x) * 180) / Math.PI

				const [px, py] = project(lon, lat)
				pathData.push(`L ${px} ${py}`)
			}

			path.setAttribute('d', pathData.join(' '))
			path.setAttribute('fill', 'none')
			path.setAttribute('stroke', 'var(--map-line)')
			path.setAttribute('stroke-width', '0.75')
			svg.appendChild(path)

			const zoom = true
			if (zoom) {
				const padding = 25
				const width = Math.abs(x2 - x1) + padding * 2
				const height = Math.abs(y2 - y1) + padding * 2
				const minX = Math.min(x1, x2) - padding
				const minY = Math.min(y1, y2) - padding
				svg.setAttribute('viewBox', `${minX} ${minY} ${width} ${height}`)
			}
		}
	})
</script>

<div class="map-wrapper h-auto" bind:this={mapWrapper}>
	{@html WorldMapSVG}
</div>

<style>
	.map-wrapper :global(svg) {
		width: 100%;
		height: auto;

		--land-color: hsl(var(--map-land));
		--border-color: hsl(var(--map-border));

		> :global(path) {
			stroke-linejoin: bevel;
		}
	}
</style>
