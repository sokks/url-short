import {
    SET_SHOW_INFO
  } from '../actions/actionTypes.js';
  
  const initialState = {
    showInfo: false,
  };
  
  export default function infoReducer(state = initialState, action) {
    switch (action.type) {
      case SET_SHOW_INFO:
        return {
          ...state,
          showInfo: action.payload,
        };
      default:
        return state;
    }
  }