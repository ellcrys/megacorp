// @flow
import React from 'react'
import Base from '../base/Base'
import { connect } from 'react-redux'
import { mapStateToProps, matchDispatchToProps } from './mappings'
import {InvokeError} from '../../errors/http'
import uuid4 from 'uuid4'
import base64 from 'base-64'

export type PropTypes = {}

export type Account = {
    ID: string,
    Balance: number,
    FirstName: string,
    LastName: string
}

export type State = {
    accountID: string,
    account?: Account
}


class GetAccount extends Base {
    propTypes: PropTypes
    state: State
    
    constructor(props={}) {
        super(props)
        this.state = {
            accountID: "",
        }
    }
    
    getAccount(){
        if (!this.state.accountID) {
            return alert("Account ID is required")
        }
        
        this.invoke(uuid4(), "get-account", [this.state.accountID]).then((data) => {
            this.setState({ accountID: "", account: JSON.parse(base64.decode(data.body)) })
        }).catch((e: InvokeError) => {
            alert(e.body.msg || e.message)
        })
    }
    
    render(){
        
        let accountInfo = null
        if (this.state.account) {
            accountInfo = <div>
                <span>First Name: <b>{this.state.account.FirstName}</b></span><br/>
                <span>Last Name: <b>{this.state.account.LastName}</b></span><br/>
                <span>Balance: <b>{this.state.account.Balance}</b></span><br/>
            </div>
        }
        
        return <div className="mid-panel">
            <div className="heading">
                GET MEGACORP ACCOUNT
            </div>
            <div className="content">
                <div className="columns">
                    <div className="column is-10">
                        <p className="control">
                            <input value={this.state.accountID} onChange={(e) => this.setState({ accountID: e.target.value})} className="input" type="text" placeholder="Enter Account ID" />
                        </p>
                    </div>
                    <div className="column">
                        <p className="control">
                            <button onClick={this.getAccount.bind(this)} className="button is-primary">Get</button>
                        </p>
                    </div>
                </div>
                {accountInfo}
            </div>
        </div>
    }
}

export default connect(mapStateToProps, matchDispatchToProps)(GetAccount);