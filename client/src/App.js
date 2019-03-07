import React, { Component } from 'react';
import './App.css';
import Nav from './Nav';
import Main from './Main';
import './editstyle.css'

class App extends Component {
  render() {
    return (
      <div className="App">
        <Nav />
        <header><b class="pb-5">HELLO WELCOME TO DCODE</b></header>
        <Main />
      </div>
    );
  }
}

export default App;
