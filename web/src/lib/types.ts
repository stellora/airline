import type { components } from './airline.openapi'

export type Aircraft = components['schemas']['Aircraft']
export type AircraftType = components['schemas']['AircraftType']
export type Airport = components['schemas']['Airport']
export type Point = components['schemas']['Point']
export type Airline = components['schemas']['Airline']
export type AirlineSpec = components['schemas']['AirlineSpec']
export type Schedule = components['schemas']['Schedule']
export type Flight = components['schemas']['Flight']
export type DaysOfWeek = components['schemas']['DaysOfWeek']
export type Passenger = components['schemas']['Passenger']
export type Itinerary = components['schemas']['Itinerary']
export type SeatAssignment = components['schemas']['SeatAssignment']
export type SeatNumber = components['schemas']['SeatNumber']
export type Fleet = components['schemas']['Fleet']
export type AirlineLoyalty = {
  id: number;
  airline: Airline;
  passengerId: number;
  mileageBalance: number;
}
