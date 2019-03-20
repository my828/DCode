import React, { Component } from 'react';
import './App.css';
// import Nav from './Nav';
import Main from './Main';
import Splash from './SplashScreen'

import './editstyle.css';
import {HashRouter as Router, Switch, Route, Redirect} from "react-router-dom";

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      pageState: {
        sessionID: "",
        code: "// Welcome to Dcode!",
        figures: ""
      }
    }
    this.socket = null;
  }

  // initializes the websocket connection -- once we have access to the sessionID
  initializeSocket = (sessionID) => {
    if (sessionID) {
      const WSS_ENDPOINT = `ws://localhost:4000/ws`;
      let requestURL = `${WSS_ENDPOINT}/${sessionID}`;
      console.log(`wss request url: ${requestURL}`);

      this.socket = new WebSocket(`${WSS_ENDPOINT}/${sessionID}`);

      // handle errors first
      this.socket.onerror = (err) => {
        console.log(`error opening websocket connection`, err);
      }

      // connection is open
      this.socket.onopen = () => {
        console.log(`websocket connection opened for session: ${sessionID}`)
      };

      // listen for messages
      this.socket.onmessage = (evt) => {
        let message = JSON.parse(evt.data);
        let figures = JSON.parse(message.figures);
        let pageState = {
          sessionID: message.sessionID,
          figures: figures,
          code: message.code
        };
        this.setState({
          pageState: pageState
        });
      }
    }
  }
  
  // returns the websocket connection if it exists
  getSocket = () => {
    if (this.socket) {
      return this.socket;
    }
  }

  // makes a http request to the api to make sure that sessionID is valid
  getSessionID = (sessionID) => {
      const DCODE_API = `http://localhost:4000/dcode`;

      let requestURL = `${DCODE_API}/v1/${sessionID}`;
      // send request to API
      fetch(requestURL, {
          method: "GET",
      })
      .then(res => {
          return res.text();
      })
      .then(body => {
          this.setState({
            sessionID: sessionID
          });
          console.log(`Got sessionID at the specific resource: ${body}`);
          this.initializeSocket(body);
      })
      .catch(err => {
          console.log(`Error: ${err}`);
      });
  }

  // passed to Editor and Canvas -- invoked when the state changes
  updatePageState = (page) => {
    this.setState({
      pageState: page
    }, () => {
      let socket = this.getSocket();
      let modified = {
        sessionID: page.sessionID,
        figures: JSON.stringify(page.figures),
        code: page.code
      }
      socket.send(JSON.stringify(modified));
    });
  }

  renderPage() {
    return <Main state={this.state.pageState} updateState={this.updatePageState} sessionID={this.state.pageState.sessionID} getSocket={this.getSocket} initializeSocket={this.initializeSocket}></Main>
  }

  renderSplashPage() {
      return <Splash state={this.state.pageState} sessionID={this.state.sessionID} getSessionID={this.getSessionID} initializeSocket={this.initializeSocket}></Splash>
  }

  render() {
    return (
      <div>
        <Router>
          <Switch>
              <Route exact path="/dcode" render={() => this.renderSplashPage()}></Route>
              <Route exact path={"/dcode/:sessionID"} render={() => this.renderPage()}></Route>
          </Switch>
        </Router>
      </div>
    );
  }
}

export default App;
