import React from "react";
import { Link, withRouter } from "react-router-dom";

class Splash extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            sessionID: this.props.state.sessionID,
            pageState: this.props.state
        }
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
        this.props.history.push(`/dcode/${sessionID}`);
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
            this.props.getSessionID(body);
            this.props.history.push(`/dcode/${body}`);
        })
        .catch(err => {
            console.log(`Error: ${err}`);
        });
    }

    render() {
        let style = {
            textAlign: "center"
        };

        return(
            <div className={"jumbotron jumbotron-fluid"} style={style}>
                <header><b className="pb-5" style={{fontSize: "30px"}}>DCODE</b></header>
                <Link to={`/dcode/${this.state.sessionID}`}>
                    <button type="button" className="btn btn-primary" onClick={() => this.fetchSessionID()}>
                        Start DCode
                    </button>
                </Link>

                <h3>or</h3>

                <form onSubmit={(evt) => this.handleSubmit()}>
                    <label>
                        Enter SessionID:
                        <input type="text" value={this.state.value} onChange={this.handleChange}></input>
                    </label>
                    
                    <br/><br/>
                    
                    <button type='submit' value="Submit" className={"btn btn-outline-primary"}>Go to page</button>
                </form>
            </div>
        )
    }
}

export default withRouter(Splash);
