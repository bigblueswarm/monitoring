import React from 'react'
import { render } from '@testing-library/react'

import { Page } from './page'

test('it should render without error', () => {
  render(<Page.Body />)
  render(<Page.Header preTitle="pre-title" title="title" />)
  render(<Page.Wrapper />)
})
