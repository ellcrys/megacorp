// @flow
import { Component } from 'react';
import constants from '../../constants'
import {HTTPError,InvokeError} from '../../errors/http'

type State = any
type Props = any

// FetchHandleResponse handles response from fetch. It returns 
// a resolved promise if response code is between 200-299 or a rejected
// rejected promise if otherwise.
export async function FetchHandleResponse(response: Response): Promise<any> { 
    let resp = await response.json()
    if (response.ok) {
        return Promise.resolve(resp)
    }
    return Promise.reject(new HTTPError(response.status, response.statusText, resp))
}

// Base is the base component
class Base extends Component {
    state: State
    props: Props
    
     /**
     * Sends an invoke request to the cocoon.
     * 
     * @param {String} id       A unique id for this request. UUID4 Expected
     * @param {String} func     The function to invoke
     * @param {Array}  params   A list of parameters to send to the cocoon code
     * @returns Promise
     * @memberof Cocoon
     */
    invoke(id: string, func: string, params: Array<string>|void) {
        return new Promise((resolve, reject) => {
            const invokeEndpoint = constants.contractHost + "/v1/invoke"
            fetch(new Request(invokeEndpoint, { 
                method: 'POST', 
                headers: new Headers({ "Content-Type": "application/json" }),
                body: JSON.stringify({ id: id, "function": func, params: params })
            })).then(FetchHandleResponse).then(function (data) {
                resolve(data)
            }).catch(function(err: HTTPError){
                reject(new InvokeError(err.status, err.body))
            })
        })
    }
}

export default Base