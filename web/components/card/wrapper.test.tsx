import React from 'react'
import { render } from '@testing-library/react'

import { Wrapper } from './wrapper'

describe('Card.Wrapper', () => {
  test('it should render without error', () => {
    render(<Wrapper />)
  })

  test('it should render children elements', () => {
    const text = 'Hello world'
    const component = render(<Wrapper><span>{text}</span></Wrapper>)
    expect(component.getByTestId('card-wrapper').textContent).toEqual(text)
  })
})
