import { 
  CLEAR_FORM, 
} from './actionTypes';

// const sleep = delay => new Promise(resolve => setTimeout(resolve, delay));

export function shortenURL(longURL) {
  return async (dispatch) => {
    let response = await callApi('POST', 's/new', {url: longURL} );
    if (response.json && response.json.status === 200) {
      await sleep(500);
      updateVmsList(dispatch);
    } else {
      dispatch(popupShowWithParams(getPopupErrorText(response), false, function(){}, 'Error'));
      console.log('ERROR:', response);
    }
    dispatch(clearForm());
  }
}

function clearForm() {
  return { type: CLEAR_FORM, payload: null };
}