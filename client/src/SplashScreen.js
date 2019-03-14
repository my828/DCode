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
        //console.log(this.props);
    }

    componentWillReceiveProps(props) {
        //console.log(props);
        this.setState({
            sessionID: props.sessionID
        });
    }
    handleChange = (event) => {
        this.setState({value: event.target.value})
        console.log(event.target.value)
    }
    render() {
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
                <form onSubmit={this.handleSubmit}>
                    <label>
                    Enter SessionID: 
                    <input type="text" value={this.state.value} onChange={this.handleChange}></input>
                    </label>
                    <input type='submit' value="Submit"></input>
                </form>
                <h3>or</h3>
                <Link to={`/dcode/${this.state.sessionID}`}>
                    <button type="button" class="btn btn-primary" onClick={() => this.props.handleNewSession()}>
                            Start DCode
                    </button>
                </Link>
            </div>
        )
    }
}

export default Splash;
