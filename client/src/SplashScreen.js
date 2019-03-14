import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

class Splash extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            sessionID: this.props.sessionID,
        }
    }

    componentDidMount() {
        
    }

    componentWillReceiveProps(props) {
        this.setState({
            sessionID: props.sessionID
        });
    }

    // onChange event handler for SessionID input field
    handleChange = (event) => {
        this.setState({
            formValue: event.target.value
        });
    }

    // handleSubmit handles clicks on the submit button after entering sessionID
    handleSubmit(evt) {
        let sessionID = this.state.formValue;
        // request the sessionID end point again
        this.props.getSessionID(sessionID);
    }

    // fetchSessionID makes a http request to generate a new sessionID
    fetchSessionID() {
        const DCODE_API = `http://localhost:4000/dcode`;
        let requestURL = `${DCODE_API}/v1/new`;
        // send request to API
        fetch(requestURL, {
            method: "GET",
        })
        .then(res => {
            return res.text();
        })
        .then(body => {
            // @TODO: change this
            console.log(`Got sessionID: ${body}`);
            this.props.getSessionID(body);
        })
        .catch(err => {
            console.log(`Error: ${err}`);
        });
    }

    render() {
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
                <form onSubmit={(evt) => this.handleSubmit()}>
                    <label>
                    Enter SessionID: 
                    <input type="text" value={this.state.value} onChange={this.handleChange}></input>
                    </label>
                    <input type='submit' value="Submit"></input>
                </form>
                <h3>or</h3>
                <button type="button" class="btn btn-primary" onClick={() => this.fetchSessionID()}>
                    Start DCode
                </button>
            </div>
        )
    }
}

export default Splash;
