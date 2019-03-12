import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

export default class Splash extends React.Component {
    render() {
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
                {/* <button type="button" class="btn btn-dark p-2" 
                    onClick={() => this.setState({landing: !this.state.landing})}>
                    Start DCode
                </button> */}
                <Link to="/decode/v1/"><button type="button" 
                        class="btn btn-primary">
                            Start DCode
                        </button>
                </Link>
            </div>
        )
    }
}
