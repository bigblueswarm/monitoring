import React from 'react'
import { render } from '@testing-library/react'

import { Server } from './server'

test('it should render without error', () => {
  render(<Server />)
})
