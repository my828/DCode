import React, { Component } from 'react';
import './App.css';
// import Nav from './Nav';
import Main from './Main';
import Splash from './SplashScreen'

import './editstyle.css';
import {HashRouter as Router, Switch, Link, Route, Redirect} from "react-router-dom";

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      sessionID: "",
      code: "// Welcom to Dcode!",
      figures: ""
    }
    this.updateSessionID = this.updateSessionID.bind(this);
    this.socket = null;
  }

  componentDidUpdate() {
    console.log(`app state was updated...`);
    if (this.state.sessionID) {
        const WSS_ENDPOINT = `ws://localhost:4000/ws`;

        let requestURL = `${WSS_ENDPOINT}/${this.state.sessionID}`;
        console.log(`wss request url: ${requestURL}`);
        this.socket = new WebSocket(`${WSS_ENDPOINT}/${this.state.sessionID}`);

        // handle errors first
        this.socket.onerror = (err) => {
            console.log(`error opening websocket connection`, err);
        }

        // connection is open
        this.socket.onopen = () => {
          console.log(`websocket connection opened for session: ${this.state.sessionID}`);
        }

        // @TODO: handle messages
    }
  }

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
      })
      .catch(err => {
          console.log(`Error: ${err}`);
      });
  }

  passSocket = (socket) => {
    this.socket = socket
  }

  updatePageState = (page) => {
    this.setState({
      sessionID: page.sessionID,
      code: page.code,
      figures: page.figures
    }, () => {
      // this.socket.send(this.state);
      console.log(page.code)
      console.log(page.figures)
    });
  }

  processMessages() {
    this.socket.onmessage = (event) => {
      var message = JSON.parse(event.data);
      console.log(message)
      this.setState({
        sessionID: message.sessionID,
        code: message.code,
        figures: message.figures
      })
    }
  }

  updateSessionID(sessionID) {
    this.setState({
      sessionID: sessionID
    });
  }

  renderPage() {
    return <Main state={this.state} update={this.updatePageState} socket={this.passSocket}></Main>
  }

  renderSplashPage() {
    return <Splash sessionID={this.state.sessionID} getSessionID={this.getSessionID}></Splash>
  }

  render() {
    return (
      <div>
        <Router>
          <Switch>
              <Route exact path="/dcode" render={() => this.renderSplashPage()}></Route>
              <Route exact path={"/dcode/:sessionID"} render={() => this.renderPage()}></Route>
              <Redirect to="/dcode" />
          </Switch>
        </Router>
      </div>
    );
  }
}

export default App;
