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
      sessionID: "",
      code: "",
      figures: ""
    }
    this.updateSessionID = this.updateSessionID.bind(this);
    this.socket = null;
  }

  componentDidUpdate() {
    this.processMessages();
  }

  handleNewSession = () => {
    const API_WS = 'ws://localhost:4000/ws/'
    const DCODE_API = "http://localhost:4000/dcode"; // api.harshiakkaraju/decode 
    fetch(`${DCODE_API}/v1/new`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(res => {
       return res.text()
    })
    .then(sessionID => {
      const socket = new WebSocket(`${API_WS}${sessionID}`);
      this.setState({
        sessionID: sessionID,
        socket: socket
      }, () => {
        // check connection
        socket.onopen = () => {
            console.log("Socket Connect!");
        }
        
      });
    })
    .catch(err => {
      //window.alert("session does not exist!");
    })
  }

  passSocket(socket) {
    this.socket = socket
  }

  updatePageState(page) {
    this.setState({
      sessionID: page.sessionID,
      code: page.code,
      figures: page.figures
    }, () => {
      this.socket.send(this.state);
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
    console.log(this.state.sessionID);
    return <Splash sessionID={this.state.sessionID} handleNewSession={this.handleNewSession}></Splash>
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
              <Route exact path="/dcode" render={() => this.renderSplashPage()}></Route>
              <Route exact path={"/dcode/:sessionID"} render={() => this.renderPage()}></Route>
          </Switch>
        </Router>
      </div>
    );
  }
}

export default App;
