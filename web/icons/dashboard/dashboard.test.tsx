import React from 'react'
import { render } from '@testing-library/react'

import { Dashboard } from './dashboard'

test('it should render without error', () => {
  render(<Dashboard />)
})
