import React, { FunctionComponent, PropsWithChildren } from 'react'

export const Wrapper: FunctionComponent<PropsWithChildren> = ({ children }) => (
    <div className="card" data-testid="card-wrapper">
        {children}
    </div>
)
