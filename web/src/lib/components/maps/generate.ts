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

async function run() {
	const geojsonFeatures = []

	const geojsonURLs = [
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_0_countries_lakes.geojson',
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_1_states_provinces_lakes.geojson'
	]
	for (const url of geojsonURLs) {
		const resp = await fetch(url)
		if (!resp.ok) {
			throw new Error(`Error fetching GeoJSON data: ${resp.status} ${resp.statusText}`)
		}
		const geojsonData: any = await resp.json()

		// Skip USA from countries data because we add the more granular state-level data in the 2nd
		// data file.
		if (url.endsWith('countries_lakes.geojson')) {
			const filteredFeatures = geojsonData.features.filter((feature: any) => {
				return feature.properties.GEOUNIT !== 'United States of America'
			})
			geojsonData.features = filteredFeatures
		}

		geojsonFeatures.push(...geojsonData.features)
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
		if (omitAntarctica && feature.properties.GEOUNIT === 'Antarctica') {
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
	console.log(svgContent)
}

if (typeof process !== 'undefined' && process.env.GENERATE) {
	run().catch((error) => {
		console.error(error)
		process.exit(1)
	})
}
