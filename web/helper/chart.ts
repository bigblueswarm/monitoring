import { Point } from '../types'

import { ApexOptions } from 'apexcharts'

export const NewAreaChart = (points: Point[], color: string): ApexOptions => {
  const data = Array.from(points, v => v.value)
  const labels = Array.from(points, v => v.time)

  return {
    chart: {
      type: 'area',
      fontFamily: 'inherit',
      sparkline: {
        enabled: true
      },
      animations: {
        enabled: false
      }
    },
    dataLabels: {
      enabled: false
    },
    fill: {
      opacity: 0.16,
      type: 'solid'
    },
    stroke: {
      width: 2,
      lineCap: 'round',
      curve: 'smooth'
    },
    series: [{
      data
    }],
    grid: {
      strokeDashArray: 4
    },
    xaxis: {
      labels: {
        show: false
      },
      tooltip: {
        enabled: false
      },
      axisBorder: {
        show: false
      },
      type: 'datetime'
    },
    yaxis: {
      labels: {
        show: false
      },
      tooltip: {
        enabled: false
      }
    },
    labels,
    colors: [color],
    legend: {
      show: false
    },
    tooltip: {
      enabled: false
    }
  }
}
