import React, { FunctionComponent, PropsWithChildren } from 'react'

export const Body: FunctionComponent<PropsWithChildren> = ({ children }) => (
    <div className="card-body" data-testid="card-body">{children}</div>
)
