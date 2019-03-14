import React from "react";
import Editor from './Editor'
import Canvas from './Canvas'
import Nav from './Nav';


export default class Main extends React.Component {
    render() {
        return(
            <div>
                <Nav />
                <div class="d-flex">
                    {/* <Canvas socket={this.props.socket} info={this.props.info}/>
                    <Editor socket={this.props.socket} info={this.props.info}/> */}
                    <Canvas />
                    <Editor />
                </div>
            </div>
        )
    }
}
