// @flow
import 'whatwg-fetch';
import Promise from 'promise-polyfill'; 
import React, { Component } from 'react';
import Navigation from './navigation/Navigation';
import Overview from './overview/Overview'
import CreateAccount from './create-account/CreateAccount'
import CreateEmployee from './create-employee/CreateEmployee'
import GetAccount from './get-account/GetAccount'
import GetTotalMegaCoin from './get-total-megacoin/GetTotalMegaCoin'
import { Route } from 'react-router'

if (!window.Promise) {
	window.Promise = Promise;
}

class App extends Component {
	render() {
		return (
		<div>
			<Navigation />
			<div className="container">
				<Route exact path="/" component={Overview}/>
				<Route exact path="/create-account" component={CreateAccount}/>
				<Route exact path="/create-employee" component={CreateEmployee}/>
				<Route exact path="/get-account" component={GetAccount}/>
				<Route exact path="/megacoin-balance" component={GetTotalMegaCoin}/>
			</div>
		</div>
		);
	}
}

export default App;
