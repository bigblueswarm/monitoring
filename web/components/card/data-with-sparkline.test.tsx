import { render } from '@testing-library/react'
import React from 'react'

import { DataWithSparkline, Trend } from './data-with-sparkline'

describe('Trend', () => {
  test('it should render without error', () => {
    render(<Trend value={12.5}/>)
  })

  test('a positive value should render Trend with green text and trending up icon', () => {
    const component = render(<Trend value={12.5}/>)
    const text = component.getByText('12.5%')
    const icon = component.getByTestId('trending-up-icon')
    expect(text.classList.contains('text-green'))
    expect(icon).toBeDefined()
  })

  test('a positive value should render Trend with red text and tending down icon', () => {
    const component = render(<Trend value={-12.5}/>)
    const text = component.getByText('-12.5%')
    const icon = component.getByTestId('trending-down-icon')
    expect(text.classList.contains('text-red'))
    expect(icon).toBeDefined()
  })
})

describe('DataWithSparkline', () => {
  test('it should render without error', () => {
    render(<DataWithSparkline value='value' title='title' points={[]} sparklineColor="#000000" trend={12.5} loading={false}/>)
  })

  test('header and subheader should be visible', () => {
    const value = (5438).toString()
    const title = 'Active users'
    const component = render(<DataWithSparkline value={value} title={title} points={[]} sparklineColor="#000000" trend={12.5} loading={false}/>)
    expect(component.getByTestId('subheader').textContent).toEqual(title)
    expect(component.getByTestId('value').textContent).toEqual(value)
  })

  test('a trend equals to 0 should not be visible', () => {
    const component = render(<DataWithSparkline value={'54342'} title={'active users'} points={[]} sparklineColor="#000000" trend={0} loading={false}/>)
    expect(component.queryByText('0%')).toBeNull()
  })

  test('a null trend should not be visible', () => {
    const component = render(<DataWithSparkline value={'54342'} title={'active users'} points={[]} sparklineColor="#000000" loading={false} />)
    expect(component.queryByText('0%')).toBeNull()
  })

  test('a valid trend should be visible', () => {
    const trend = 53
    const component = render(<DataWithSparkline value={'54342'} title={'active users'} points={[]} sparklineColor="#000000" trend={trend} loading={false} />)
    expect(component.getByTestId('trend').textContent?.trim()).toEqual(`${trend}%`)
  })

  test('a loading state should display loading placeholder', () => {
    const component = render(<DataWithSparkline value={''} title={'active users'} points={[]} sparklineColor="#000000" trend={0} loading={true} />)
    expect(component.getByTestId('loading-value')).not.toBeNull()
    expect(component.getByTestId('loading-chart')).not.toBeNull()
  })
})
