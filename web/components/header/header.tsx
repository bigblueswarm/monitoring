import React, { FunctionComponent } from 'react'
import { Link } from 'react-router-dom'

import Logo from '../../statics/images/dark.logo.svg'
import Icon from '../../icons'

export const Header: FunctionComponent = () => (
    <header className="navbar navbar-expand-md navbar-dark navbar-overlap d-print-none">
        <div className="container-xl">
            <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu">
                <span className="navbar-toggler-icon"></span>
            </button>
            <h1 className="navbar-brand d-none-navbar-horizontal pe-0 pe-md-3">
                <img src={Logo} width="110" height="32" alt="Tabler" className="navbar-brand-image" />
                <span className="m-2">BigBlueSwarm</span>
            </h1>
            <div className="navbar-nav flex-row order-md-last align-items-center">
                <a href="/auth/logout" className="btn btn-indigo h-33">
                    <Icon.Logout />
                    Logout
                </a>
            </div>
            <div className="collapse navbar-collapse" id="navbar-menu">
                <div className="d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center">
                    <ul className="navbar-nav">
                        <li className="nav-item">
                            <Link to={'/'} className="nav-link">
                                <span className="nav-link-icon d-md-none d-lg-inline-block">
                                    <Icon.Dashboard />
                                </span>
                                <span className="nav-link-title">
                                    Overview
                                </span>
                            </Link>
                        </li>
                        <li className="nav-item">
                            <Link to={'/servers'} className="nav-link pe-none">
                                <span className="nav-link-icon d-md-none d-lg-inline-block">
                                    <Icon.Server />
                                </span>
                                <span className="nav-link-title">
                                    Servers
                                </span>
                            </Link>
                        </li>
                        <li className="nav-item">
                            <Link to={'/config'} className="nav-link pe-none">
                                <span className="nav-link-icon d-md-none d-lg-inline-block">
                                    <Icon.Configuration />
                                </span>
                                <span className="nav-link-title">
                                    Configuration
                                </span>
                            </Link>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </header>
)
