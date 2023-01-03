import React, { FunctionComponent } from 'react'

import { Page } from '../../components/page/page'

import {
  ActiveUsers
} from '../../containers/active-users'

import { ActiveMeetings } from '../../containers/active-meetings'

export const Overview: FunctionComponent = () => (
  <Page.Wrapper>
    <Page.Header preTitle='Overview' title='Monitor your cluster' />
    <Page.Body>
      <div className='col-sm-6 col-lg-4'>
        <ActiveUsers.CardWithSparkline />
      </div>
      <div className='col-sm-6 col-lg-4'>
        <ActiveMeetings.CardWithSparkline />
      </div>
    </Page.Body>
  </Page.Wrapper>
)
