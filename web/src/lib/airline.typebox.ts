/* eslint eslint-comments/no-unlimited-disable: off */
/* eslint-disable */
// This document was generated automatically by openapi-box

/**
 * @typedef {import('@sinclair/typebox').TSchema} TSchema
 */

/**
 * @template {TSchema} T
 * @typedef {import('@sinclair/typebox').Static<T>} Static
 */

/**
 * @typedef {import('@sinclair/typebox').SchemaOptions} SchemaOptions
 */

/**
 * @typedef {{
 *  [Path in keyof typeof schema]: {
 *    [Method in keyof typeof schema[Path]]: {
 *      [Prop in keyof typeof schema[Path][Method]]: typeof schema[Path][Method][Prop] extends TSchema ?
 *        Static<typeof schema[Path][Method][Prop]> :
 *        undefined
 *    }
 *  }
 * }} SchemaType
 */

/**
 * @typedef {{
 *  [ComponentType in keyof typeof _components]: {
 *    [ComponentName in keyof typeof _components[ComponentType]]: typeof _components[ComponentType][ComponentName] extends TSchema ?
 *      Static<typeof _components[ComponentType][ComponentName]> :
 *      undefined
 *  }
 * }} ComponentType
 */

import { Type as T, TypeRegistry, Kind, CloneType } from '@sinclair/typebox'
import { Value } from '@sinclair/typebox/value'

/**
 * @typedef {{
 *  [Kind]: 'Binary'
 *  static: string | File | Blob | Uint8Array
 *  anyOf: [{
 *    type: 'object',
 *    additionalProperties: true
 *  }, {
 *    type: 'string',
 *    format: 'binary'
 *  }]
 * } & TSchema} TBinary
 */

/**
 * @returns {TBinary}
 */
const Binary = () => {
  /**
   * @param {TBinary} schema
   * @param {unknown} value
   * @returns {boolean}
   */
  function BinaryCheck(schema, value) {
    const type = Object.prototype.toString.call(value)
    return (
      type === '[object Blob]' ||
      type === '[object File]' ||
      type === '[object String]' ||
      type === '[object Uint8Array]'
    )
  }

  if (!TypeRegistry.Has('Binary')) TypeRegistry.Set('Binary', BinaryCheck)

  return /** @type {TBinary} */ ({
    anyOf: [
      {
        type: 'object',
        additionalProperties: true
      },
      {
        type: 'string',
        format: 'binary'
      }
    ],
    [Kind]: 'Binary'
  })
}

const ComponentsSchemasAircraftRegistration = T.String({
  pattern: '^[A-Z][A-Z0-9-]{1,12}$',
  minLength: 2,
  maxLength: 13
})
const ComponentsSchemasAircraftTypeIcaoCode = T.String({
  pattern: '^[A-Z][A-Z0-9]{1,3}$',
  minLength: 2,
  maxLength: 4
})
const ComponentsSchemasAirlineIataCode = T.String({
  pattern: '^[A-Z0-9]{2}$',
  minLength: 2,
  maxLength: 2
})
const ComponentsSchemasAirline = T.Object({
  id: T.Integer(),
  iataCode: CloneType(ComponentsSchemasAirlineIataCode),
  name: T.String()
})
const ComponentsSchemasAircraft = T.Object({
  id: T.Integer(),
  registration: CloneType(ComponentsSchemasAircraftRegistration),
  aircraftType: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
  airline: CloneType(ComponentsSchemasAirline)
})
const ComponentsSchemasAirlineId = T.Integer()
const ComponentsSchemasAirlineSpec = T.Union([
  CloneType(ComponentsSchemasAirlineId),
  CloneType(ComponentsSchemasAirlineIataCode)
])
const ComponentsSchemasAircraftId = T.Integer()
const ComponentsSchemasAircraftSpec = T.Union([
  CloneType(ComponentsSchemasAircraftId),
  CloneType(ComponentsSchemasAircraftRegistration)
])
const ComponentsParametersAircraftSpec = T.Any()
const ComponentsParametersAirlineSpec = T.Any()
const ComponentsSchemasAircraftType = T.Object({
  icaoCode: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
  name: T.String()
})
const ComponentsSchemasAirportIataCode = T.String({
  pattern: '^[A-Z]{3}$',
  minLength: 3,
  maxLength: 3
})
const ComponentsSchemasPoint = T.Object({
  longitude: T.Number({ format: 'double' }),
  latitude: T.Number({ format: 'double' })
})
const ComponentsSchemasAirport = T.Object({
  id: T.Integer(),
  name: T.String(),
  iataCode: CloneType(ComponentsSchemasAirportIataCode),
  country: T.String(),
  region: T.String(),
  point: CloneType(ComponentsSchemasPoint),
  timezoneID: T.String()
})
const ComponentsSchemasAirportId = T.Integer()
const ComponentsSchemasAirportSpec = T.Union([
  CloneType(ComponentsSchemasAirportId),
  CloneType(ComponentsSchemasAirportIataCode)
])
const ComponentsParametersAirportSpec = T.Any()
const ComponentsSchemasFlightNumber = T.String({
  pattern: '^[0-9]{1,4}$',
  minLength: 1,
  maxLength: 4
})
const ComponentsSchemasLocalDate = T.String({
  pattern: '^\\d{4}-\\d{2}-\\d{2}$'
})
const ComponentsSchemasDaysOfWeek = T.Array(
  T.Union([
    T.Literal(0),
    T.Literal(1),
    T.Literal(2),
    T.Literal(3),
    T.Literal(4),
    T.Literal(5),
    T.Literal(6)
  ]),
  { minItems: 0, maxItems: 7, uniqueItems: true }
)
const ComponentsSchemasTimeOfDay = T.String({ pattern: '[0-2]?\\d:[0-5]\\d' })
const ComponentsSchemasFlightSchedule = T.Object({
  id: T.Integer(),
  airline: CloneType(ComponentsSchemasAirline),
  number: CloneType(ComponentsSchemasFlightNumber),
  originAirport: CloneType(ComponentsSchemasAirport),
  destinationAirport: CloneType(ComponentsSchemasAirport),
  distanceMiles: T.Number({ format: 'double' }),
  aircraftType: CloneType(ComponentsSchemasAircraftType),
  startDate: CloneType(ComponentsSchemasLocalDate),
  endDate: CloneType(ComponentsSchemasLocalDate),
  daysOfWeek: CloneType(ComponentsSchemasDaysOfWeek),
  departureTime: CloneType(ComponentsSchemasTimeOfDay),
  durationSec: T.Integer(),
  published: T.Boolean()
})
const ComponentsSchemasZonedDateTime = T.String({ format: 'date-time' })
const ComponentsSchemasFlightInstance = T.Object({
  id: T.Integer(),
  scheduleID: T.Optional(T.Integer()),
  scheduleInstanceDate: T.Optional(CloneType(ComponentsSchemasLocalDate)),
  airline: CloneType(ComponentsSchemasAirline),
  number: CloneType(ComponentsSchemasFlightNumber),
  originAirport: CloneType(ComponentsSchemasAirport),
  destinationAirport: CloneType(ComponentsSchemasAirport),
  distanceMiles: T.Number({ format: 'double' }),
  aircraftType: CloneType(ComponentsSchemasAircraftType),
  aircraft: T.Optional(CloneType(ComponentsSchemasAircraft)),
  departureDateTime: CloneType(ComponentsSchemasZonedDateTime),
  arrivalDateTime: CloneType(ComponentsSchemasZonedDateTime),
  notes: T.String(),
  published: T.Boolean({ default: false })
})
const ComponentsSchemasSeatNumber = T.String({
  pattern: '^[0-9]{1,2}[A-Z]$',
  minLength: 2,
  maxLength: 3
})
const ComponentsSchemasSeatAssignment = T.Object({
  id: T.Integer(),
  itineraryID: T.Integer(),
  passengerID: T.Integer(),
  flightInstanceID: T.Integer(),
  seat: CloneType(ComponentsSchemasSeatNumber)
})
const ComponentsSchemasPassenger = T.Object({
  id: T.Integer(),
  name: T.String()
})
const ComponentsSchemasRecordLocator = T.String({
  pattern: '^[A-Z0-9]{6}$',
  minLength: 6,
  maxLength: 6
})
const ComponentsSchemasItinerary = T.Object({
  id: T.Integer(),
  recordID: CloneType(ComponentsSchemasRecordLocator),
  flights: T.Array(CloneType(ComponentsSchemasFlightInstance)),
  passengers: T.Array(CloneType(ComponentsSchemasPassenger), { minItems: 1 })
})
const ComponentsSchemasItinerarySpec = T.Union([
  T.Integer(),
  CloneType(ComponentsSchemasRecordLocator)
])
const ComponentsParametersItinerarySpec = T.Any()
const ComponentsSchemasRoute = T.Object({
  originAirport: CloneType(ComponentsSchemasAirport),
  destinationAirport: CloneType(ComponentsSchemasAirport),
  distanceMiles: T.Number({ format: 'double' }),
  flightSchedulesCount: T.Integer()
})

const schema = {
  '/health': {
    GET: {
      args: T.Void(),
      data: T.Object(
        {
          ok: T.Optional(T.Boolean())
        },
        {
          'x-status-code': '200',
          'x-content-type': 'application/json'
        }
      ),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/aircraft': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasAircraft), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            registration: CloneType(ComponentsSchemasAircraftRegistration),
            aircraftType: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
            airline: CloneType(ComponentsSchemasAirlineSpec)
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAircraft, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    },
    DELETE: {
      args: T.Void(),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/aircraft/{aircraftSpec}': {
    GET: {
      args: T.Object({
        params: T.Object({
          aircraftSpec: CloneType(ComponentsSchemasAircraftSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: CloneType(ComponentsSchemasAircraft, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          aircraftSpec: CloneType(ComponentsSchemasAircraftSpec, {
            'x-in': 'path'
          })
        }),
        body: T.Object(
          {
            registration: T.Optional(
              CloneType(ComponentsSchemasAircraftRegistration)
            ),
            aircraftType: T.Optional(
              CloneType(ComponentsSchemasAircraftTypeIcaoCode)
            ),
            airline: T.Optional(CloneType(ComponentsSchemasAirlineSpec))
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAircraft, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          aircraftSpec: CloneType(ComponentsSchemasAircraftSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/airlines/{airlineSpec}/aircraft': {
    GET: {
      args: T.Object({
        params: T.Object({
          airlineSpec: CloneType(ComponentsSchemasAirlineSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Array(CloneType(ComponentsSchemasAircraft), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/aircraft-types': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasAircraftType), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/airports': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasAirport), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            iataCode: CloneType(ComponentsSchemasAirportIataCode)
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAirport, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    },
    DELETE: {
      args: T.Void(),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/airports/{airportSpec}': {
    GET: {
      args: T.Object({
        params: T.Object({
          airportSpec: CloneType(ComponentsSchemasAirportSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: CloneType(ComponentsSchemasAirport, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          airportSpec: CloneType(ComponentsSchemasAirportSpec, {
            'x-in': 'path'
          })
        }),
        body: T.Object(
          {
            iataCode: T.Optional(CloneType(ComponentsSchemasAirportIataCode))
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAirport, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          airportSpec: CloneType(ComponentsSchemasAirportSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/airports/{airportSpec}/flight-schedules': {
    GET: {
      args: T.Object({
        params: T.Object({
          airportSpec: CloneType(ComponentsSchemasAirportSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Array(CloneType(ComponentsSchemasFlightSchedule), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/airlines': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasAirline), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            iataCode: CloneType(ComponentsSchemasAirlineIataCode),
            name: T.String()
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAirline, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    },
    DELETE: {
      args: T.Void(),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/airlines/{airlineSpec}': {
    GET: {
      args: T.Object({
        params: T.Object({
          airlineSpec: CloneType(ComponentsSchemasAirlineSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: CloneType(ComponentsSchemasAirline, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          airlineSpec: CloneType(ComponentsSchemasAirlineSpec, {
            'x-in': 'path'
          })
        }),
        body: T.Object(
          {
            iataCode: T.Optional(CloneType(ComponentsSchemasAirlineIataCode)),
            name: T.Optional(T.String())
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasAirline, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          airlineSpec: CloneType(ComponentsSchemasAirlineSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/airlines/{airlineSpec}/flight-schedules': {
    GET: {
      args: T.Object({
        params: T.Object({
          airlineSpec: CloneType(ComponentsSchemasAirlineSpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Array(CloneType(ComponentsSchemasFlightSchedule), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/flight-schedules': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasFlightSchedule), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            airline: CloneType(ComponentsSchemasAirlineSpec, {
              'x-in': 'path'
            }),
            number: CloneType(ComponentsSchemasFlightNumber),
            originAirport: CloneType(ComponentsSchemasAirportSpec, {
              'x-in': 'path'
            }),
            destinationAirport: CloneType(ComponentsSchemasAirportSpec, {
              'x-in': 'path'
            }),
            aircraftType: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
            startDate: CloneType(ComponentsSchemasLocalDate),
            endDate: CloneType(ComponentsSchemasLocalDate),
            daysOfWeek: CloneType(ComponentsSchemasDaysOfWeek),
            departureTime: CloneType(ComponentsSchemasTimeOfDay),
            durationSec: T.Integer(),
            published: T.Optional(T.Boolean({ default: false }))
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasFlightSchedule, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    },
    DELETE: {
      args: T.Void(),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/flight-schedules/{id}': {
    GET: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: CloneType(ComponentsSchemasFlightSchedule, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        }),
        body: T.Object(
          {
            airline: T.Optional(
              CloneType(ComponentsSchemasAirlineSpec, { 'x-in': 'path' })
            ),
            number: T.Optional(CloneType(ComponentsSchemasFlightNumber)),
            originAirport: T.Optional(
              CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' })
            ),
            destinationAirport: T.Optional(
              CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' })
            ),
            aircraftType: T.Optional(
              CloneType(ComponentsSchemasAircraftTypeIcaoCode)
            ),
            startDate: T.Optional(CloneType(ComponentsSchemasLocalDate)),
            endDate: T.Optional(CloneType(ComponentsSchemasLocalDate)),
            daysOfWeek: T.Optional(CloneType(ComponentsSchemasDaysOfWeek)),
            departureTime: T.Optional(CloneType(ComponentsSchemasTimeOfDay)),
            durationSec: T.Optional(T.Integer()),
            published: T.Optional(T.Boolean())
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasFlightSchedule, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/flight-schedules/{id}/instances': {
    GET: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: T.Array(CloneType(ComponentsSchemasFlightInstance), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/flight-instances': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasFlightInstance), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            airline: CloneType(ComponentsSchemasAirlineSpec, {
              'x-in': 'path'
            }),
            number: CloneType(ComponentsSchemasFlightNumber),
            originAirport: CloneType(ComponentsSchemasAirportSpec, {
              'x-in': 'path'
            }),
            destinationAirport: CloneType(ComponentsSchemasAirportSpec, {
              'x-in': 'path'
            }),
            aircraftType: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
            aircraft: T.Optional(
              CloneType(ComponentsSchemasAircraftSpec, { 'x-in': 'path' })
            ),
            departureDateTime: CloneType(ComponentsSchemasZonedDateTime),
            arrivalDateTime: CloneType(ComponentsSchemasZonedDateTime),
            notes: T.String(),
            published: T.Optional(T.Boolean({ default: false }))
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasFlightInstance, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    }
  },
  '/flight-instances/{id}': {
    GET: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: CloneType(ComponentsSchemasFlightInstance, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        }),
        body: T.Object(
          {
            airline: T.Optional(
              CloneType(ComponentsSchemasAirlineSpec, { 'x-in': 'path' })
            ),
            number: T.Optional(CloneType(ComponentsSchemasFlightNumber)),
            originAirport: T.Optional(
              CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' })
            ),
            destinationAirport: T.Optional(
              CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' })
            ),
            aircraftType: T.Optional(
              CloneType(ComponentsSchemasAircraftTypeIcaoCode)
            ),
            aircraft: T.Optional(
              CloneType(ComponentsSchemasAircraftSpec, { 'x-in': 'path' })
            ),
            departureDateTime: T.Optional(
              CloneType(ComponentsSchemasZonedDateTime)
            ),
            arrivalDateTime: T.Optional(
              CloneType(ComponentsSchemasZonedDateTime)
            ),
            notes: T.Optional(T.String()),
            published: T.Optional(T.Boolean({ default: false }))
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasFlightInstance, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([
        T.Any({ 'x-status-code': '400' }),
        T.Any({ 'x-status-code': '404' })
      ])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([
        T.Any({ 'x-status-code': '400' }),
        T.Any({ 'x-status-code': '404' })
      ])
    }
  },
  '/flight-instances/{flightInstanceID}/seat-assignments': {
    GET: {
      args: T.Object({
        params: T.Object({
          flightInstanceID: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: T.Array(CloneType(ComponentsSchemasSeatAssignment), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    POST: {
      args: T.Object({
        params: T.Object({
          flightInstanceID: T.Integer({ 'x-in': 'path' })
        }),
        body: T.Object(
          {
            itineraryID: T.Integer(),
            passengerID: T.Integer(),
            seat: CloneType(ComponentsSchemasSeatNumber)
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasSeatAssignment, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([
        T.Any({ 'x-status-code': '400' }),
        T.Any({ 'x-status-code': '404' })
      ])
    }
  },
  '/passengers': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasPassenger), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            name: T.String({ minLength: 1, maxLength: 255 })
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasPassenger, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    }
  },
  '/passengers/{id}': {
    GET: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: CloneType(ComponentsSchemasPassenger, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    PATCH: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        }),
        body: T.Object(
          {
            name: T.Optional(T.String())
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasPassenger, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([
        T.Any({ 'x-status-code': '400' }),
        T.Any({ 'x-status-code': '404' })
      ])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          id: T.Integer({ 'x-in': 'path' })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/itineraries': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasItinerary), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    },
    POST: {
      args: T.Object({
        body: T.Object(
          {
            flightInstanceIDs: T.Array(T.Integer(), { minLength: 1 }),
            passengerIDs: T.Array(T.Integer(), { minLength: 1 })
          },
          {
            'x-content-type': 'application/json'
          }
        )
      }),
      data: CloneType(ComponentsSchemasItinerary, {
        'x-status-code': '201',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '400' })])
    }
  },
  '/itineraries/{itinerarySpec}': {
    GET: {
      args: T.Object({
        params: T.Object({
          itinerarySpec: CloneType(ComponentsSchemasItinerarySpec, {
            'x-in': 'path'
          })
        })
      }),
      data: CloneType(ComponentsSchemasItinerary, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    },
    DELETE: {
      args: T.Object({
        params: T.Object({
          itinerarySpec: CloneType(ComponentsSchemasItinerarySpec, {
            'x-in': 'path'
          })
        })
      }),
      data: T.Any({ 'x-status-code': '204' }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  },
  '/routes': {
    GET: {
      args: T.Void(),
      data: T.Array(CloneType(ComponentsSchemasRoute), {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': 'default' })])
    }
  },
  '/routes/{route}': {
    GET: {
      args: T.Object({
        params: T.Object({
          route: T.String({ pattern: '^[A-Z]{3}-[A-Z]{3}$', 'x-in': 'path' })
        })
      }),
      data: CloneType(ComponentsSchemasRoute, {
        'x-status-code': '200',
        'x-content-type': 'application/json'
      }),
      error: T.Union([T.Any({ 'x-status-code': '404' })])
    }
  }
}

const _components = {
  parameters: {
    aircraftSpec: CloneType(ComponentsSchemasAircraftSpec, { 'x-in': 'path' }),
    airportSpec: CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' }),
    airlineSpec: CloneType(ComponentsSchemasAirlineSpec, { 'x-in': 'path' }),
    itinerarySpec: CloneType(ComponentsSchemasItinerarySpec, { 'x-in': 'path' })
  },
  schemas: {
    ZonedDateTime: CloneType(ComponentsSchemasZonedDateTime),
    AircraftID: CloneType(ComponentsSchemasAircraftId),
    AircraftRegistration: CloneType(ComponentsSchemasAircraftRegistration),
    AircraftSpec: CloneType(ComponentsSchemasAircraftSpec, { 'x-in': 'path' }),
    Aircraft: CloneType(ComponentsSchemasAircraft),
    AircraftTypeICAOCode: CloneType(ComponentsSchemasAircraftTypeIcaoCode),
    AircraftType: CloneType(ComponentsSchemasAircraftType),
    AirlineID: CloneType(ComponentsSchemasAirlineId),
    AirlineIATACode: CloneType(ComponentsSchemasAirlineIataCode),
    AirlineSpec: CloneType(ComponentsSchemasAirlineSpec, { 'x-in': 'path' }),
    Airline: CloneType(ComponentsSchemasAirline),
    AirportID: CloneType(ComponentsSchemasAirportId),
    AirportIATACode: CloneType(ComponentsSchemasAirportIataCode),
    AirportSpec: CloneType(ComponentsSchemasAirportSpec, { 'x-in': 'path' }),
    Airport: CloneType(ComponentsSchemasAirport),
    Point: CloneType(ComponentsSchemasPoint),
    LocalDate: CloneType(ComponentsSchemasLocalDate),
    DaysOfWeek: CloneType(ComponentsSchemasDaysOfWeek),
    TimeOfDay: CloneType(ComponentsSchemasTimeOfDay),
    FlightSchedule: CloneType(ComponentsSchemasFlightSchedule),
    FlightNumber: CloneType(ComponentsSchemasFlightNumber),
    FlightInstance: CloneType(ComponentsSchemasFlightInstance),
    Route: CloneType(ComponentsSchemasRoute),
    RecordLocator: CloneType(ComponentsSchemasRecordLocator),
    ItinerarySpec: CloneType(ComponentsSchemasItinerarySpec, {
      'x-in': 'path'
    }),
    Itinerary: CloneType(ComponentsSchemasItinerary),
    Passenger: CloneType(ComponentsSchemasPassenger),
    SeatNumber: CloneType(ComponentsSchemasSeatNumber),
    SeatAssignment: CloneType(ComponentsSchemasSeatAssignment)
  }
}

export { schema, _components as components }
