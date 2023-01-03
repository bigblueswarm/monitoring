import React, { FunctionComponent } from 'react'
import { gql, useQuery } from '@apollo/client'

import { Card } from '../../components/card'

import { getColor } from '../../helper/color'
import { Gauge, Point, Trend } from '../../types'

interface IActiveRecordings {
  activeRecordings: {
    Gauge: Gauge
    Trend: Trend
    Sparkline: Point[]
  }
}

const GET_ACTIVE_RECORDINGS = gql`
    {
        activeRecordings(start: "-1h", stop: "now()") {
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
  const { data, loading } = useQuery<IActiveRecordings>(GET_ACTIVE_RECORDINGS, {
    pollInterval: 10000,
    fetchPolicy: 'no-cache'
  })

  return <Card.DataWithSparkline
    title='Active recordings'
    value={data?.activeRecordings.Gauge.value.toString()}
    points={data?.activeRecordings.Sparkline}
    trend={data?.activeRecordings.Trend.value}
    sparklineColor={getColor('azure')}
    loading={loading}
  />
}
