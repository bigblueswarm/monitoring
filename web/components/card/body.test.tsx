import React from 'react'
import { render } from '@testing-library/react'

import { Body } from './body'

describe('Card.Body', () => {
  test('it should render without error', () => {
    render(<Body />)
  })

  test('it should render children elements', () => {
    const text = 'Hello world'
    const component = render(<Body><span>{text}</span></Body>)
    expect(component.getByTestId('card-body').textContent).toEqual(text)
  })
})
