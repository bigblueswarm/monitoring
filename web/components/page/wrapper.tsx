import React, { FunctionComponent, PropsWithChildren } from 'react'

export const Wrapper: FunctionComponent<PropsWithChildren> = ({ children }) => (
    <div className='page-wrapper'>
        {children}
    </div>
)
