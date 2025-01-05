import { DateFormatter, ZonedDateTime } from '@internationalized/date'

const dateFormatterOptions: Intl.DateTimeFormatOptions = {
	year: 'numeric',
	month: 'short',
	day: 'numeric',
	weekday: 'short',
	hour: 'numeric',
	hour12: false,
	minute: '2-digit',
	timeZoneName: 'short',
}

/**
 * A {@link DateFormatter} that shows dates in the client's current timezone.
 */
export const clientTimezoneDateFormatter = new DateFormatter(
	navigator.language,
	dateFormatterOptions,
)

const cachedDateFormatters: { [timeZoneName: string]: DateFormatter } = {}

function dateFormatterFor(value: ZonedDateTime): DateFormatter {
	let formatter = cachedDateFormatters[value.timeZone]
	if (!formatter) {
		formatter = cachedDateFormatters[value.timeZone] = new DateFormatter(navigator.language, {
			...dateFormatterOptions,
			timeZone: value.timeZone,
		})
	}
	return formatter
}

export function formatDateFull(value: ZonedDateTime): string {
	const f = dateFormatterFor(value)
	return `${f.format(value.toDate())} (${value.timeZone})`
}

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
		const daysDelta = deltaCalendarDays(plusMinusDaysFrom, value)
		if (daysDelta !== 0) {
			result += `${daysDelta < 0 ? '-' : '+'}${Math.abs(daysDelta)}`
		}
	}

	return result
}

/**
 * The difference in the number of calendar
 */
export function deltaCalendarDays(a: ZonedDateTime, b: ZonedDateTime): number {
	if (a.calendar.identifier !== b.calendar.identifier) {
		throw new Error('calendars must match')
	}
	if (a.era !== b.era) {
		throw new Error('eras must match')
	}

	if (a.year === b.year && a.month === b.month) {
		return b.day - a.day // fast path
	}
	return a.calendar.toJulianDay(b) - a.calendar.toJulianDay(a)
}

export function formatFlightDuration(start: ZonedDateTime, end: ZonedDateTime): string {
	const ms = end.toDate().getTime() - start.toDate().getTime()
	const hours = Math.floor(ms / 3600000)
	const minutes = Math.floor((ms % 3600000) / 60000)
	return `${hours}h ${minutes}m`
}
