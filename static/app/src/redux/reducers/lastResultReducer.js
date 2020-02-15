import {
    SET_SHOW_LAST_RESULT,
    SET_LAST_RESULT
  } from '../actions/actionTypes.js';
  
  const initialState = {
    lastResult: null,
    showLastResult: false,
  };
  
  export default function lastResultReducer(state = initialState, action) {
    switch (action.type) {
      case SET_SHOW_LAST_RESULT:
        return {
          ...state,
          showLastResult: action.payload,
        };
      case SET_LAST_RESULT:
        return {
          ...state,
          lastResult: action.payload,
        };
      default:
        return state;
    }
  }