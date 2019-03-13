import React, { Component } from 'react';
import './App.css';
// import Nav from './Nav';
import Main from './Main';
import Splash from './SplashScreen'

import './editstyle.css';
import {HashRouter as Router, Switch, Redirect, Route} from "react-router-dom";


class App extends Component {
  render() {
    return (
      <Router>
        <Switch>
            <Route exact path="/" component={Splash}></Route>
            <Route path="/content" component={Main}></Route>
        </Switch>
        {/* <div className="App">
          <Nav />
          <Main />
        </div> */}
      </Router>
    );
  }
}

export default App;
