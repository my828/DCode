import React from "react";
import Editor from './Editor'
import Canvas from './Canvas'
import Nav from './Nav';
import { withRouter } from "react-router-dom";
import './index.css';

class Main extends React.Component {
    componentDidMount() {
        let socket = this.props.getSocket();
        if (!socket) {
            let path = window.location.href;
            let components = path.split("/");
            let sessionID = components[components.length - 1];
            this.props.initializeSocket(sessionID);
        }
    }

    render() {
        return(
            <div>
                <Nav/>
                <div className={"main-container"}>
                    <div className={"main-canvas-container"}>
                        <Canvas state={this.props.state} updateState={this.props.updateState}/>
                    </div>
                    <div className={"main-editor-container"}>
                        <Editor state={this.props.state} updateState={this.props.updateState} />
                    </div>
                </div>
            </div>
        )
    }
}

export default withRouter(Main);
