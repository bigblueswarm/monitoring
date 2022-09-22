import React, { FunctionComponent } from 'react'

interface HeaderProps {
  preTitle: string
  title: string
}

export const Header: FunctionComponent<HeaderProps> = ({ preTitle, title }) => (
    <div className='page-header d-print-none text-white'>
        <div className='container-xl'>
            <div className='row g-2 align-items-center'>
                <div className="page-pretitle">
                    {preTitle}
                </div>
                <h2 className="page-title">
                    {title}
                </h2>
            </div>
        </div>
    </div>
)
