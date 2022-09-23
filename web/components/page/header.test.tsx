import React from 'react'
import { render } from '@testing-library/react'

import { Header } from './header'

test('it should render without error', () => {
  render(<Header preTitle="pre-title" title="title" />)
})
