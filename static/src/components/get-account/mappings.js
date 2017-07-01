// @flow
import {bindActionCreators} from 'redux';

export function mapStateToProps(state: {}) {
    return {
    };
}

export function matchDispatchToProps(dispatch: any){
    return bindActionCreators({}, dispatch);
}

