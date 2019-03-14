import React, { Component } from 'react';
import './App.css';
// import Nav from './Nav';
import Main from './Main';
import Splash from './SplashScreen'

import './editstyle.css';
import {HashRouter as Router, Switch, Link, Route} from "react-router-dom";

class App extends Component {

  handleNewSession = () => {
    fetch("http://localhost:4000/dcode/v1/new", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(
      res => {
        return res.json();
      }
    ).then(data => {
      // this.setState({
      //   info: data.body,
      //   sessionID: data.sessonID
      // })
      console.log(data)
    }).catch(
      
    )

    const socket = new WebSocket('ws://localhost:4000/ws/')
    this.setState({
      socket: socket
    })
  }

  componentDidMount() {
    
  }

  render() {
    return (
      <div>
        <Router>
        {/* <div style={{textAlign: "center"}}>
          <header>
            <b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b>
          </header>
          <Link to="/decode/v1/:sessionID"><button type="button" 
                  class="btn btn-primary" onClick={this.handleNewSession}>
                      Start DCode
                  </button>
          </Link>
        </div> */}
        <div>
          <Link to="/decode/v1/:sessionID"><button type="button" 
                    class="btn btn-primary" onClick={this.handleNewSession}>
                        Start DCode
                    </button>
            </Link>
          <Switch>

              <Route path="/decode/v1/:sessionID" render={(routerProps) => (
                // <Main {...routerProps} info={this.state.info} socket={this.state.socket}/>
                <Main />
              )}></Route>
          </Switch>
        </div>
        </Router>
      </div>
    );
  }
}

export default App;
