import React from "react";
import Editor from './Editor'
import Canvas from './Canvas'
import Nav from './Nav';
// import { constants } from "http2";

export default class Main extends React.Component {
    componentDidMount() {
        const API_WS = 'ws://localhost:4000/ws/'
        const DCODE_API = "http://localhost:4000/dcode";
        fetch(`${DCODE_API}${this.props.state.sessionID}`, {
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
             window.alert("Session does not exits!")
        })
    }

    render() {
        return(
            <div>
                <Nav />
                <div class="d-flex">
                    {/* <Canvas socket={this.props.socket} info={this.props.info}/>
                    <Editor socket={this.props.socket} info={this.props.info}/> */}
                    <Canvas state={this.props.state} update={this.props.update}/>
                    <Editor state={this.props.state} update={this.props.update}/>
                </div>
            </div>
        )
    }
}
