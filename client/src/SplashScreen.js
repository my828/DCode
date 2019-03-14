import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

export default class Splash extends React.Component {
    constructor(props) {
        super(props)
    }

    handleNewSession = () => {
        fetch("http://localhost:4000/dcode/v1/new", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            }
        }).then(

        )
        const socket = new WebSocket('ws://localhost:4000/ws/')
        socket.onmessage = event => {
        console.log(event)
        // const code = JSON.parse()
        }
    }

    render() {
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
                <Link to="/decode/v1/"><button type="button" 
                        class="btn btn-primary" onClick={this.handleNewSession}>
                            Start DCode
                        </button>
                </Link>
            </div>
        )
    }
}
