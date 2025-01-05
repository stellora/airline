import { parseZonedDateTime } from '@internationalized/date'
import { describe, expect, test } from 'vitest'
import { deltaCalendarDays, formatFlightTime } from './datetime-helpers'

describe('formatFlightTime', () => {
	test('plusMinusDaysFrom', () => {
		const a = parseZonedDateTime('2025-01-02T03:05:01-10:00[America/Adak]')
		const b = parseZonedDateTime('2025-01-03T12:05:01-08:00[America/Los_Angeles]')
		const c = parseZonedDateTime('2024-12-31T12:05:01+00:00[Europe/London]')
		const d = parseZonedDateTime('2025-02-02T07:05:01+08:00[Asia/Singapore]')
		expect.soft(formatFlightTime(b, { plusMinusDaysFrom: a })).toBe('12:05+1')
		expect.soft(formatFlightTime(c, { plusMinusDaysFrom: a })).toBe('12:05-2')
		expect.soft(formatFlightTime(d, { plusMinusDaysFrom: b })).toBe('07:05+29')
	})
})

describe('deltaCalendarDays', () => {
	test('same date', () => {
		const a = parseZonedDateTime('2025-01-02T03:05:01-10:00[America/Adak]')
		const b = parseZonedDateTime('2025-01-02T12:05:01-08:00[America/Los_Angeles]')
		expect(deltaCalendarDays(a, b)).toBe(0)
	})

	test('different dates', () => {
		const a = parseZonedDateTime('2025-01-02T03:05:01-10:00[America/Adak]')
		const b = parseZonedDateTime('2025-01-03T12:05:01-08:00[America/Los_Angeles]')
		expect(deltaCalendarDays(a, b)).toBe(1)
	})

	test('different months', () => {
		const a = parseZonedDateTime('2025-01-29T03:05:01-10:00[America/Adak]')
		const b = parseZonedDateTime('2025-02-02T12:05:01-08:00[America/Los_Angeles]')
		expect(deltaCalendarDays(a, b)).toBe(4)
	})

	test('different years', () => {
		const a = parseZonedDateTime('2024-12-31T03:05:01-10:00[America/Adak]')
		const b = parseZonedDateTime('2025-01-01T12:05:01-08:00[America/Los_Angeles]')
		expect(deltaCalendarDays(a, b)).toBe(1)
	})
})
