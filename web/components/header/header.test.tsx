import React from 'react'
import { render } from '@testing-library/react'

import { Header } from './header'
import { MemoryRouter } from 'react-router-dom'

test('it should render without error', () => {
  render(
        <MemoryRouter>
            <Header />
        </MemoryRouter>
  )
})
