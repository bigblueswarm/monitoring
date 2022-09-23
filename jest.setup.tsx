import React from 'react'

class ResizeObserverMock {
  disconnect (): void {}
  observe (): void {}
  unobserve (): void {}
}

window.ResizeObserver = ResizeObserverMock

jest.mock('react-apexcharts', () => {
  return {
    __esModule: true,
    default: () => {
      // eslint-disable-next-line react/react-in-jsx-scope
      return <div />
    }
  }
})
