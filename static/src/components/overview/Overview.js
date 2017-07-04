// @flow
import React from 'react'
import Base from '../base/Base'
import { connect } from 'react-redux'
import { mapStateToProps, matchDispatchToProps } from './mappings'
import uuid4 from 'uuid4'
import base64 from 'base-64'
import {InvokeError} from '../../errors/http'
import constants from '../../constants'

export type PropTypes = {
}

export type Account = {
    ID: string,
    Balance: number,
    FirstName: string,
    LastName: string
}

export type Employee = {
    AccountID: string,
    Position: string
}

export type State = {
    accounts: Array<Account>,
    employees: Array<Employee>
}

class Overview extends Base {
    propTypes: PropTypes
    state: State
    
    constructor(props = {}) {
        super(props)
        this.state = {
            accounts: [],
            employees: []
        }
    }
    
    componentDidMount () {
        this.getAllAccounts()
        this.getAllEmployees()
    }
    
    getAllAccounts() {
        this.invoke(uuid4(), "get-all-accounts", []).then((data) => {
            this.setState({ accounts: JSON.parse(base64.decode(data.body)) || [] })
        }).catch((e: InvokeError) => {
            alert("failed to create account: " + e.body.msg)
        })
    }
    
    getAllEmployees() {
        this.invoke(uuid4(), "get-all-employees", []).then((data) => {
            this.setState({ employees: JSON.parse(base64.decode(data.body)) || [] })
        }).catch((e: InvokeError) => {
            alert("failed to create account: " + e.body.msg)
        })
    }
    
    listAccounts(){
        return this.state.accounts.map((account: Account) => {
            return <article key={account.ID} className="media">
                <figure className="media-left">
                    <p className="image is-64x64">
                    <img src="http://bulma.io/images/placeholders/128x128.png" alt="{account.FirstName} {account.LastName}" />
                    </p>
                </figure>
                <div className="media-content">
                    <div className="content">
                    <div><strong>{account.FirstName} {account.LastName}</strong> <small>{account.ID}</small></div>
                    <p><b>Balance:</b> {account.Balance}</p>
                    </div>
                </div>
            </article>
        })
    }
    
    listEmployees(){
        return this.state.employees.map((employee: Employee) => {
            return <article key={employee.AccountID} className="media">
                <figure className="media-left">
                    <p className="image is-64x64">
                    <img src={constants.imgDir + "/media/128x128.svg"} alt={employee.AccountID} />
                    </p>
                </figure>
                <div className="media-content">
                    <div className="content">
                    <p>
                        <strong>{employee.Position}</strong> <small>{employee.AccountID}</small>
                    </p>
                    </div>
                </div>
            </article>
        })
    }
    
    render() {
        
        let createAcctBtn = <a href="create-account" className="button is-primary">Create</a>
        if (this.state.accounts.length > 0) {
            createAcctBtn = null
        }
        
        let createEmpBtn = <a href="create-employee" className="button is-primary">Create</a>
        if (this.state.employees.length > 0) {
            createEmpBtn = null
        }
        
        return <div className="overview">
            <div className="columns">
                <div className="column">
                    <h1 className="title is-4">Accounts</h1>
                    <h3 className="subtitle is-6">These are all existing MegaCorp accounts</h3>
                    <div className="content">
                        {createAcctBtn}
                        {this.listAccounts()}
                    </div>
                </div>
                <div className="column ">
                    <h1 className="title is-4">Employees</h1>
                    <h3 className="subtitle is-6">These are all existing MegaCorp employees</h3>
                    <div className="content">
                        {createEmpBtn}
                        {this.listEmployees()}
                    </div>
                </div>
            </div>
        </div>
    }
}


export default connect(mapStateToProps, matchDispatchToProps)(Overview);