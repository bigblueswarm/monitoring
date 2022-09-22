import React, { FunctionComponent, PropsWithChildren } from 'react'

export const Body: FunctionComponent<PropsWithChildren> = ({ children }) => (
    <div className="page-body">
        <div className="container-xl">
            <div className="row row-deck row-cards">
                {children}
            </div>
        </div>
    </div>
)
