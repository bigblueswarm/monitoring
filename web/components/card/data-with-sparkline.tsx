import React, { FunctionComponent } from 'react'
import Chart from 'react-apexcharts'

import { Wrapper } from './wrapper'
import { Body } from './body'
import Icons from '../../icons'

import { Point } from '../../types'
import { NewAreaChart } from '../../helper/chart'

interface DataWithSparklineProps {
  title: string
  value: string | undefined
  trend?: number | undefined
  points: Point[] | undefined
  sparklineColor: string
  loading: boolean
}

export const Trend: FunctionComponent<{ value: number }> = ({ value }) => (
  <div className="me-auto" data-testid="trend">
    <span className={`${value > 0 ? 'text-green' : 'text-red'} d-inline-flex align-items-center lh-1`}>
      {value}% {value > 0 ? <Icons.Trending.Up /> : <Icons.Trending.Down />}
    </span>
  </div>
)

export const DataWithSparkline: FunctionComponent<DataWithSparklineProps> = ({ title, value, trend, points, sparklineColor, loading }) => {
  return (
    <Wrapper>
        <Body>
            <div className='d-flex align-items-center'>
                <div className='subheader' data-testid="subheader">{title}</div>
            </div>
            {
              loading
                ? <div data-testid="loading-value">
                  <div className='placeholder placeholder-xs col-10'/ >
                  <div className='placeholder placeholder-xs col-11'/ >
                </div>
                : <div className='d-flex align-items-baseline'>
                <div className='h1 mb-3 me-2' data-testid="value">{value}</div>
                {
                  (trend != null && trend !== 0) && <Trend value={trend} />
                }
              </div>
            }
        </Body>
        {
          loading
            ? <div data-testid="loading-chart" className='chart-sm placeholder'></div>
            : <div className='chart-sm'>
          { (points != null) && <Chart
              height={40}
              options={NewAreaChart(points, sparklineColor)}
              series={NewAreaChart(points, sparklineColor).series}
              type="area"
           />}
      </div>
        }
    </Wrapper>
  )
}
