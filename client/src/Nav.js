import React from "react";
import logo from './logo.svg'
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
            <nav class="navbar navbar-expand-md mb-2 m-0 border-bottom" style={navstyle}>
                <img src={logo} class="logo" alt="logo" style={logostyle}/>
            </nav>
        )
    }
}