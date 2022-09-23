import React from 'react'
import { render } from '@testing-library/react'

import { Wrapper } from './wrapper'

test('it should render without error', () => {
  render(<Wrapper />)
})
