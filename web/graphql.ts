import { ApolloClient, InMemoryCache } from '@apollo/client'

export const client = new ApolloClient({
  uri: '/query',
  cache: new InMemoryCache()
})
