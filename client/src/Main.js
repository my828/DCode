import React from "react";
import Editor from './Editor'
import Canvas from './Canvas'
import Nav from './Nav';
// import { constants } from "http2";

export default class Main extends React.Component {
    constructor(props) {
        super(props)
        this.socket = null
    }
    componentDidMount() {
        const API_WS = 'ws://localhost:4000/ws/'
        const DCODE_API = "http://localhost:4000/dcode/";
        fetch(`${DCODE_API}${this.props.state.sessionID}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(res => {
            return res.text()
        })
        .then((res) => {
            const socket = new WebSocket(`${DCODE_API}${this.props.state.sessionID}`);
            this.socket = socket
            socket.onopen = () => {
                console.log("CONNECT")
            }
            console.log(res)
            this.props.socket(socket)
            // this.setState({
            //     sessionID: sessionID,
                
            // }, () => {
            //     // check connection
            //     socket.onopen = () => {
            //         console.log("Socket Connect!");
            //     }
                
            // });
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
