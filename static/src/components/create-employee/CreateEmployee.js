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
    accountID: string,
    position: string,
    employeeCreated: boolean,
    validPositions: Array<string>
}

class CreateAccount extends Base {
    propTypes: PropTypes
    state: State
    
    constructor(props={}) {
        super(props)
        this.state = {
            accountID: "",
            position: "ceo",
            employeeCreated: false,
            validPositions: ["ceo", "coo", "cto", "frontend_dev", "backend_dev"]
        }
    }
    
    create(){
        if (!this.state.accountID || !this.state.position) {
            return alert("Account ID and Position are required")
        }
        
        this.invoke(uuid4(), "create-employee", [this.state.accountID, this.state.position]).then((data) => {
            this.setState({ employeeCreated: true, accountID: "" })
        }).catch((e: InvokeError) => {
            alert(e.body.msg || e.message)
        })
    }
    
    
    render() {
        let notification = null
        if (this.state.employeeCreated) {
            notification = <div className="notification is-success">Employee Created! </div>
        }
        
        let positionOptions = this.state.validPositions.map((p: string) => {
            return <option key={p} value={p}>{p}</option>
        })
        
        return <div className="mid-panel">
            <div className="heading">
                CREATE MEGACORP EMPLOYEE
            </div>
            {notification}
            <div className="content">
                <label className="label">Account ID</label>
                <p className="control">
                    <input value={this.state.accountID} onChange={(e) => this.setState({ accountID: e.target.value})} className="input" type="text" />
                </p>
                <label className="label">Position</label>
                <p className="control">
                    <span className="select">
                        <select onChange={(e) => this.setState({ position: e.target.value })}>{positionOptions}</select>
                    </span>
                </p>
                <p className="control">
                    <button onClick={this.create.bind(this)} className="button is-primary">Create</button>
                </p>
            </div>
        </div>
    }
}

export default connect(mapStateToProps, matchDispatchToProps)(CreateAccount);
