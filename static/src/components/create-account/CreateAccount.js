// @flow
import React from 'react'
import Base from '../base/Base'
import { connect } from 'react-redux'
import { mapStateToProps, matchDispatchToProps } from './mappings'
import uuid4 from 'uuid4'
import {InvokeError} from '../../errors/http'

export type PropTypes = {
}

export type State = {
    firstName: string,
    lastName: string,
    accountCreated: boolean
}

class CreateAccount extends Base {
    propTypes: PropTypes
    state: State
    
    constructor(props={}) {
        super(props)
        this.state = {
            firstName: "",
            lastName: "",
            accountCreated: false
        }
    }
    
    create(){
        if (!this.state.firstName || !this.state.lastName) {
            return alert("First Name and Last Name are required")
        }
        
        this.invoke(uuid4(), "create-account", [this.state.firstName, this.state.lastName]).then((data) => {
            this.setState({ accountCreated: true, firstName: "", lastName: "" })
        }).catch((e: InvokeError) => {
            alert("failed to create account: " + e.body.msg)
        })
    }
    
    render() {
        let notification = null
        if (this.state.accountCreated) {
            notification = <div className="notification is-success">Account Created!</div>
        }
        return <div className="mid-panel">
            <div className="heading">
                CREATE MEGACORP ACCOUNT
            </div>
            {notification}
            <div className="content">
                <label className="label">First Name</label>
                <p className="control">
                    <input value={this.state.firstName} onChange={(e) => this.setState({ firstName: e.target.value})} className="input" type="text" />
                </p>
                <label className="label">Last Name</label>
                <p className="control">
                    <input value={this.state.lastName} onChange={(e) => this.setState({ lastName: e.target.value})} className="input" type="text" />
                </p>
                <p className="control">
                    <button onClick={this.create.bind(this)} className="button is-primary">Create</button>
                </p>
            </div>
        </div>
    }
}

export default connect(mapStateToProps, matchDispatchToProps)(CreateAccount);
