import { DateFormatter, ZonedDateTime } from '@internationalized/date'

export const dateFormatter = new DateFormatter('en-US', {
	dateStyle: 'full',
	timeStyle: 'medium',
	second: undefined,
	hour12: false,
})

export function formatFlightDate(value: ZonedDateTime): string {
	const f = new DateFormatter('en-US', {
		year: 'numeric',
		month: 'short',
		day: '2-digit',
		timeZone: value.timeZone,
	})
	return f.format(value.toDate())
}

export function formatFlightTime(
	value: ZonedDateTime,
	{ plusMinusDaysFrom }: { plusMinusDaysFrom?: ZonedDateTime } = {},
): string {
	const f = new DateFormatter('en-US', {
		hour: '2-digit',
		minute: '2-digit',
		hour12: false,
		timeZone: value.timeZone,
	})
	let result = f.format(value.toDate())

	if (plusMinusDaysFrom) {
		// TODO!(sqs): support wrapping around months/years
		const daysDelta = value.day - plusMinusDaysFrom.day
		if (daysDelta !== 0) {
			result += `${daysDelta < 0 ? '-' : '+'}${Math.abs(daysDelta)}`
		}
	}

	return result
}

export function formatFlightDuration(start: ZonedDateTime, end: ZonedDateTime): string {
	const ms = end.toDate().getTime() - start.toDate().getTime()
	const hours = Math.floor(ms / 3600000)
	const minutes = Math.floor((ms % 3600000) / 60000)
	return `${hours}h ${minutes}m`
}
