import React, { Component } from 'react';
import { Header } from 'semantic-ui-react'
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header as="h1" className="App-header">
          Welcome to Silly Point
        </Header>
      </div>
    );
  }
}

export default App;
