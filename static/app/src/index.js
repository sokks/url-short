import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import configureStore from './redux/configureStore';
import DevTools from './components/DevTools/DevTools';
import App from './components/App/App';

const store = configureStore();

const component = (
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>
);

ReactDOM.render(component, document.getElementById('root'));

if (process.env.NODE_ENV !== 'production') {
  ReactDOM.render(<DevTools store={store} />, document.getElementById('dev-tools'));
}
