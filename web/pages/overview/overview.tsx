import React, { FunctionComponent } from 'react'

import { Page } from '../../components/page/page'

export const Overview: FunctionComponent = () => (
    <Page.Wrapper>
        <Page.Header preTitle='Overview' title='Monitor your cluster' />
        <Page.Body>
            <div className='text-white'>Overview</div>
        </Page.Body>
    </Page.Wrapper>
)
