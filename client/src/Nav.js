import React from "react";
import logo from './logo.svg'
import { BrowserRouter as Link } from "react-router-dom";
import './index.css';

export default class Nav extends React.Component {
    render() {
        let logostyle = {
            height: "40px",
            position: "relative",
            left: "0"
        }
        return (
            <nav className="navbar navbar-dark bg-dark">
                <Link to="/dcode"><span className={"lead text-white"}>DCode</span></Link>
                <button className={"btn btn-sm btn-outline-warning"}>Extend Session</button>
            </nav>
        )
    }
}