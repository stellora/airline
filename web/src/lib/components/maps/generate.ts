import type { BBox, FeatureCollection } from 'geojson'

async function run() {
	const output: FeatureCollection = { type: 'FeatureCollection', features: [] }

	const geojsonURLs = [
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_0_countries_lakes.geojson',
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_1_states_provinces_lakes.geojson'
	]
	for (const url of geojsonURLs) {
		const resp = await fetch(url)
		if (!resp.ok) {
			throw new Error(`Error fetching GeoJSON data: ${resp.status} ${resp.statusText}`)
		}
		const geojsonData: FeatureCollection = await resp.json()

		// Skip USA from countries data because we add the more granular state-level data in the 2nd
		// data file.
		if (url.endsWith('countries_lakes.geojson')) {
			const filteredFeatures = geojsonData.features.filter((feature: any) => {
				return feature.properties.GEOUNIT !== 'United States of America'
			})
			geojsonData.features = filteredFeatures
		}

		output.features.push(...geojsonData.features)
	}

	// Filter out unneeded data to cut bundle size.
	for (const feature of output.features) {
		if (feature.type === 'Feature') {
			if (feature.properties) {
				feature.properties = {
					ISO_A2: feature.properties.ISO_A2
				}
			}
		}
	}

	// Reduce precision to cut bundle size.
	function round(n: number): number {
		return Number(n.toFixed(6))
	}
	for (const feature of output.features) {
		const coordinates =
			feature.geometry.type === 'Polygon'
				? [feature.geometry.coordinates]
				: feature.geometry.type === 'MultiPolygon'
					? feature.geometry.coordinates
					: []
		for (const shape of coordinates) {
			for (const [i, point] of shape.entries()) {
				shape[i] = point.map((coord) => coord.map(round))
			}
		}

		if (feature.bbox) {
			feature.bbox = feature.bbox.map(round) as BBox
		}
	}

	process.stdout.write(JSON.stringify(output))
	process.stdout.write('\n')
}

run().catch((error) => {
	console.error(error)
	process.exit(1)
})
