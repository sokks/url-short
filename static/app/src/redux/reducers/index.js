import { combineReducers } from 'redux';
import mainFormReducer from './mainFormReducer';
import infoReducer from './infoReducer';
import lastResultReducer from './lastResultReducer';
import resultsListReducer from './resultsListReducer';

export default combineReducers({
    mainForm: mainFormReducer,
    info: infoReducer,
    lastRes: lastResultReducer,
    resList: resultsListReducer,
});
