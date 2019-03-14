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

    render() {
        return(
            <div style={{textAlign: "center"}}>
                <header><b class="pb-5" style={{fontSize: "30px"}}>HELLO WELCOME TO DCODE :)</b></header>
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
