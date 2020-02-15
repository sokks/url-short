import {
    CLEAR_FORM,
  } from '../actions/actionTypes.js';
  
  const initialState = {
    longURL: '',
    showHello: true,
    lastResult: null,
    showLastResult: false,
    previousResults
    : [],
    showResultsList: false,
  };
  
  export default function mainFormReducer(state = initialState, action) {
    var newState = Object.assign({}, state);;
    switch (action.type) {
      case CLEAR_FORM:
        return {
          ...state,
          longURL: '',
        };
      default:
        return newState;
    }
  }