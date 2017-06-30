// @flow
import {bindActionCreators} from 'redux';

export function mapStateToProps(state: any) {
    return {
    };
}

export function matchDispatchToProps(dispatch: any){
    return bindActionCreators({}, dispatch);
}

