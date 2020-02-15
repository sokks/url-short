import React, { Component } from 'react';
import MainForm from '../MainForm/MainForm';
import Info from '../Info/Info';
import LastResult from '../LastResult/LastResult';
import ResultsList from '../ResultsList/ResultsList';
import './MainPage.css';

class MainPage extends Component {
  render() {
    return (
      <div className="main">
        <MainForm history={this.props.history} />
        <Info history={this.props.history} />
        <LastResult history={this.props.history} />
        <ResultsList history={this.props.history} />
      </div>
    );
  }
}

export default MainPage;
