async function run() {
	const geojsonURL =
		'https://raw.githubusercontent.com/nvkelso/natural-earth-vector/refs/heads/master/geojson/ne_110m_admin_0_countries.geojson'
	const resp = await fetch(geojsonURL)
	if (!resp.ok) {
		throw new Error(`Error fetching GeoJSON data: HTTP ${resp.status} ${resp.statusText}`)
	}
	const geojsonData = await resp.json()

	// Define map projection parameters
	const width = 960
	const height = 500

	// Create projection matrix
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

	// Convert coordinates to SVG path
	function coordsToSVGPath(coords, projection) {
		let path = ''
		coords.forEach((ring, i) => {
			ring.forEach((coord, j) => {
				const [x, y] = projection.project(coord[0], coord[1])
				path += `${j === 0 ? 'M' : 'L'}${x},${y}`
			})
			path += 'Z' // Close the path
		})
		return path
	}

	// Generate SVG
	function generateSVG(geoData, projection) {
		let svgPaths = ''

		geoData.features.forEach((feature) => {
			if (feature.geometry.type === 'Polygon') {
				const path = coordsToSVGPath(feature.geometry.coordinates, projection)
				svgPaths += `<path d="${path}" fill="#ccc" stroke="#fff" stroke-width="0.5"/>`
			} else if (feature.geometry.type === 'MultiPolygon') {
				feature.geometry.coordinates.forEach((polygon) => {
					const path = coordsToSVGPath(polygon, projection)
					svgPaths += `<path d="${path}" fill="#ccc" stroke="#fff" stroke-width="0.5"/>`
				})
			}
		})

		return `
        <svg width="${width}" height="${height}" viewBox="0 0 ${width} ${height}" xmlns="http://www.w3.org/2000/svg">
            ${svgPaths}
        </svg>
    `
	}

	// Create the map
	const projection = createProjectionMatrix()
	const svgContent = generateSVG(geojsonData, projection)

	console.log(svgContent)
}

run().catch((error) => {
	console.error(error)
	process.exit(1)
})
