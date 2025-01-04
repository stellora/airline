import { DateFormatter } from '@internationalized/date'

export const dateFormatter = new DateFormatter('en-US', {
	dateStyle: 'full',
	timeStyle: 'medium',
	second: undefined,
	hour12: false,
})

const flightDateOptions: Intl.DateTimeFormatOptions = {
	year: 'numeric',
	month: 'short',
	day: '2-digit',
}

const differentYear = new DateFormatter('en-GB', flightDateOptions)

const omitCurrentYear = false

const currentYear = omitCurrentYear
	? new DateFormatter('en-GB', {
			...flightDateOptions,
			year: undefined, // omit the current year
		})
	: differentYear

function formatterForFlightDateTime(value: Date): DateFormatter {
	const thisYear = new Date().getFullYear()
	if (value.getFullYear() === thisYear) {
		return currentYear
	}
	return differentYear
}

export function formatFlightDate(value: Date): string {
	const f = formatterForFlightDateTime(value)
	return f.format(value)
}

export const flightTimeFormatter = new DateFormatter('en-US', {
	hour: '2-digit',
	minute: '2-digit',
	hour12: false,
})

export function formatFlightTime(value: Date): string {
	return flightTimeFormatter.format(value)
}
