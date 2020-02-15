import {
    SET_SHOW_RESULTS_LIST,
    SET_RESULTS_LIST
  } from '../actions/actionTypes.js';
  
  const initialState = {
    results: [],
    showResultsList: false,
  };
  
  export default function resultsListReducer(state = initialState, action) {
    switch (action.type) {
      case SET_SHOW_RESULTS_LIST:
        return {
          ...state,
          showResultsList: action.payload,
        };
      case SET_RESULTS_LIST:
        return {
          ...state,
          resultsList: action.payload,
        };
      default:
        return state;
    }
  }