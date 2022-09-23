import React, { FunctionComponent } from 'react'
import { gql, useQuery } from '@apollo/client'

import { Card } from '../../components/card'

import { getColor } from '../../helper/color'
import { Gauge, Point, Trend } from '../../types'

interface IActiveUsers {
  activeUsers: {
    Gauge: Gauge
    Trend: Trend
    Sparkline: Point[]
  }
}

const GET_ACTIVE_USERS = gql`
    {
        activeUsers(start: "-1h", stop: "now()", every: "10s") {
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
  const { data, loading } = useQuery<IActiveUsers>(GET_ACTIVE_USERS, {
    pollInterval: 10000,
    fetchPolicy: 'no-cache'
  })

  return <Card.DataWithSparkline
    title='Active users'
    value={data?.activeUsers.Gauge.value.toString()}
    points={data?.activeUsers.Sparkline}
    trend={data?.activeUsers.Trend.value}
    sparklineColor={getColor('primary')}
    loading={loading}
  />
}
