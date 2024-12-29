async function run() {
	const geojsonURL =
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_0_countries.geojson'
	const resp = await fetch(geojsonURL)
	if (!resp.ok) {
		throw new Error(`Error fetching GeoJSON data: HTTP ${resp.status} ${resp.statusText}`)
	}
	const geojsonData = await resp.json()

	const width = 960
	const height = 500

	function createProjectionMatrix() {
		return {
			project: function (lon, lat) {
				// Simple equirectangular projection
				const x = (lon + 180) * (width / 360)
				const y = (90 - lat) * (height / 180)
				return [x, y]
			}
		}
	}

	function coordsToSVGPath(coords, projection) {
		let path = ''
		for (const ring of coords) {
			for (const [j, coord] of ring.entries()) {
				const [x, y] = projection.project(coord[0], coord[1])
				path += `${j === 0 ? 'M' : 'L'}${x},${y}`
			}
			path += 'Z' // Close the path
		}
		return path
	}


	// Create the map
	const projection = createProjectionMatrix()

	const svgPaths = []

	for (const feature of geojsonData.features) {
		if (feature.geometry.type === 'Polygon') {
			const path = coordsToSVGPath(feature.geometry.coordinates, projection)
			svgPaths.push(`<path d="${path}" fill="#ccc" stroke="#fff" stroke-width="0.5"/>`)
		} else if (feature.geometry.type === 'MultiPolygon') {
			for (const polygon of feature.geometry.coordinates) {
				const path = coordsToSVGPath(polygon, projection)
				svgPaths.push(`<path d="${path}" fill="#ccc" stroke="#fff" stroke-width="0.5"/>`)
			}
		}
	}

	const svgContent= `
	<svg width="${width}" height="${height}" viewBox="0 0 ${width} ${height}" xmlns="http://www.w3.org/2000/svg">
		${svgPaths.join('\n')}
	</svg>
`
	console.log(svgContent)
}

run().catch((error) => {
	console.error(error)
	process.exit(1)
})
