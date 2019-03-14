import React, { Component } from 'react';
import './App.css';
// import Nav from './Nav';
import Main from './Main';
import Splash from './SplashScreen'

import './editstyle.css';
import {HashRouter as Router, Switch, Link, Route} from "react-router-dom";

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      sessionID: ""
    }
    // this.sendCanvasData = this.sendCanvasData.bind(this);
  }

  componentDidMount() {
    
  }

  // sendCanvasData(event) {
  //   // use to send info through socket
  // }

  renderPage() {
    return <Main></Main>
  }

  render() {
    // "/decode/v1/:sessionID"
    // "/decode/v1/sessionID"
    // const sessionID = this.state.sessionID
    // const link = "/dcode/v1/"+sessionID
    return (
      <div>
        <Router>
          <Switch>
              <Route exact path="/dcode" component={Splash}></Route>
              <Route exact path={"/dcode/:sessionID"} render={() => this.renderPage()}></Route>
          </Switch>
        </Router>
      </div>
    );
  }
}

export default App;
