import React from "react";
import { Link } from "react-router-dom";
import './index.css';

export default class Nav extends React.Component {
    render() {
        return (
            <nav className="navbar navbar-dark bg-dark">
                <Link to="/dcode"><span className={"lead text-white"}>DCode</span></Link>
                <button className={"btn btn-sm btn-outline-warning"}>Extend Session</button>
            </nav>
        )
    }
}