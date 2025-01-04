import { DateFormatter } from '@internationalized/date'

export const dateFormatter = new DateFormatter('en-US', {
	dateStyle: 'full',
	timeStyle: 'medium',
	second: undefined,
	hour12: false,
})

export const localTimeFormatter = new DateFormatter('en-US', {
	hour: '2-digit',
	minute: '2-digit',
	hour12: false,
})

const flightDateTimeOptions: Intl.DateTimeFormatOptions = {
	month: 'short',
	day: 'numeric',
	weekday: 'short',
	hour: '2-digit',
	minute: '2-digit',
	hour12: false,
}

const sameYear = new DateFormatter('en-US', {
	...flightDateTimeOptions,
})

const differentYear = new DateFormatter('en-US', {
	...flightDateTimeOptions,
	year: 'numeric',
})

export function formatterForFlightDateTime(value: Date): DateFormatter {
	const thisYear = new Date().getFullYear()
	if (value.getFullYear() === thisYear) {
		return sameYear
	}
	return differentYear
}
