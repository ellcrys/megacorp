// @flow
import React from 'react'
import Base from '../base/Base'
import { connect } from 'react-redux'
import { mapStateToProps, matchDispatchToProps } from './mappings'
import logo from '../../images/MegaCorp.svg'
import constants from '../../constants'

export type PropTypes = {
}

class Navigation extends Base {
    propTypes: PropTypes
    
    render() {
        return <div className="container navigation">
            <nav className="nav">
                <div className="nav-left">
                    <a href="/" className="nav-item">Overview</a>
                    <a href="create-account" className="nav-item">Create Account</a>
                    <a href="create-employee" className="nav-item">Create Employee</a>
                </div>
                <div className="nav-center">
                    <div className="nav-item">
                        <img src={constants.imgDir + "/media/MegaCorp.cfc0a723.svg"} className="logo" alt="logo" />
                    </div>
                </div>  
                <div className="nav-right">
                    <a href="get-account" className="nav-item">Get Account</a>
                    <a href="megacoin-balance" className="nav-item">Get MegaCoin Balance</a>
                    <a className="nav-item" href="https://github.com/ellcrys/megacorp">Github</a>
                </div>
            </nav>
        </div>
    }
}


export default connect(mapStateToProps, matchDispatchToProps)(Navigation);