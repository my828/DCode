import React from "react";
import logo from './logo.svg'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

export default class Nav extends React.Component {
    render() {
        let logostyle = {
            height: "40px",
            position: "relative",
            left: "0"
        }
        let navstyle = {
            backgroundColor: "gay",
            padding: "10px"
        }
        return (
            <nav className="navbar navbar-expand-md mb-2 m-0 border-bottom justify-content-between" style={navstyle}>
                <Link to="/"><img src={logo} class="logo" alt="logo" style={logostyle} /></Link>
                <button>Extend Session</button>
            </nav>
        )
    }
}