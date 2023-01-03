import React, { FunctionComponent } from 'react'
import { gql, useQuery } from '@apollo/client'

import { Card } from '../../components/card'

import { getColor } from '../../helper/color'
import { Gauge, Point, Trend } from '../../types'

interface IActiveMeetings {
  activeMeetings: {
    Gauge: Gauge
    Trend: Trend
    Sparkline: Point[]
  }
}

const GET_ACTIVE_MEETINGS = gql`
    {
        activeMeetings(start: "-1h", stop: "now()") {
            Gauge {
                value
            }
            Trend {
              value
            }
            Sparkline {
                value
                time
            }
        }
    }
`

export const CardWithSparkline: FunctionComponent = () => {
  const { data, loading } = useQuery<IActiveMeetings>(GET_ACTIVE_MEETINGS, {
    pollInterval: 10000,
    fetchPolicy: 'no-cache'
  })

  return <Card.DataWithSparkline
    title='Active meetings'
    value={data?.activeMeetings.Gauge.value.toString()}
    points={data?.activeMeetings.Sparkline}
    trend={data?.activeMeetings.Trend.value}
    sparklineColor={getColor('cyan')}
    loading={loading}
  />
}
