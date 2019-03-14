import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

class Splash extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            sessionID: "",
        }
    }

    handleNewSession = () => {
        console.log("reaching here..");
        fetch("http://localhost:4000/dcode/v1/new", {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(res => {
           return res.text()
        })
        .then(sessionID => {
          this.setState({
            sessionID: sessionID
          });
          this.props.history.push(`/dcode/${sessionID}`);
        })
        .catch(err => {
          this.props.history.push(`/dcode`);
        })
      }

    render() {
        const starter = "dcode";
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
                <Link to={`/dcode/${starter}`} onClick={() => {}}>
                    <button type="button" class="btn btn-primary" onClick={() => this.handleNewSession()}>
                            Start DCode
                    </button>
                </Link>
            </div>
        )
    }
}

export default Splash;
