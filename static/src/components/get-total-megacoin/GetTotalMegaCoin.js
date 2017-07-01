// @flow
import React from 'react'
import Base from '../base/Base'
import { connect } from 'react-redux'
import { mapStateToProps, matchDispatchToProps } from './mappings'
import {InvokeError} from '../../errors/http'
import uuid4 from 'uuid4'
import base64 from 'base-64'
import numeral from 'numeral'

export type PropTypes = {}

export type State = {
    balance: string
}

class GetTotalMegaCoin extends Base {
    propTypes: PropTypes
    state: State
    
    constructor(props={}) {
        super(props)
        this.state = {
            balance: ""
        }
    }
    
    componentWillMount () {
        this.getTotalMegaCoinBalance()
    }
    
    getTotalMegaCoinBalance(){
        this.invoke(uuid4(), "get-total-supply", []).then((data) => {
            this.setState({ balance: numeral(base64.decode(data.body)).format('0,0') })
        }).catch((e: InvokeError) => {
            alert(e.body.msg || e.message)
        })
    }
    
    
    paySalaries(){
        this.invoke(uuid4(), "pay-salaries", []).then((data) => {
            alert("Paid")
            setTimeout(() => {
                window.location.reload()
            }, 5000)
        }).catch((e: InvokeError) => {
            alert(e.body.msg || e.message)
        })
    }
    
    render(){
        return <div className="mid-panel">
            <div className="heading">
                TOTAL UN-ISSUED MEGACOIN
            </div>
            <div className="has-text-centered">
                <div className="amount"><h2 className="title is-1">{this.state.balance}</h2></div>
            </div>
            <div className="pay-salary-btn">
                <button onClick={this.paySalaries.bind(this)} className="button is-primary">Pay Salaries</button>
            </div>
        </div>
    }
}

export default connect(mapStateToProps, matchDispatchToProps)(GetTotalMegaCoin);