import React, { Component } from 'react';
import { Switch, Route } from 'react-router-dom';
import MainPage from '../MainPage/MainPage';
import AboutPage from '../AboutPage/AboutPage';

class Main extends Component {
  render() {
    return (
      <Switch>
        <Route path='/about' component={AboutPage} />
        <Route component={MainPage} />
      </Switch>
    );
  }
}

export default Main;
