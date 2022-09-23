import React, { FunctionComponent } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Header } from './components/header/header'
import Pages from './pages'

export const App: FunctionComponent = () => (
    <>
        {/* <React.StrictMode> */}
            <BrowserRouter>
                <Header />
                <Routes>
                    <Route index element={<Pages.Overview />} />
                    <Route path="servers" element={<Pages.Servers />} />
                    <Route path="config" element={<Pages.Config />} />
                </Routes>
            </BrowserRouter>
        {/* </React.StrictMode> */}
    </>
)
