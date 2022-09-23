import React from 'react'
import * as ReactDOM from 'react-dom/client'
import { ApolloProvider } from '@apollo/client'

import { App } from './app'

import { client } from './graphql'

import '@tabler/core/dist/css/tabler.min.css'

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
const root = ReactDOM.createRoot(document.getElementById('app')!)

root.render(<ApolloProvider client={client}><App /></ApolloProvider>)
