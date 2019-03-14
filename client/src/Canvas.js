import React from "react";
import CanvasDraw from "react-canvas-draw";
import Websocket from 'react-websocket';
import {SketchField, Tools} from 'react-sketch';

export default class Canvas extends React.Component {
    constructor(prop) {
        super(prop) 
        this.state = {
            width: window.screen.availWidth / 2 - window.screen.availWidth / 5,
        };
        this.canvasRef = React.createRef();
    }

    // componentDidUpdate() {
    //     const socket = this.props.socket
    //     socket.send(JSON.stringify(this.canvasRef.current.toJSON()))
    // }

    // componentDidMount() {
    //     this.setState({
    //         info: this.props.info
    //     })
        
    // }

    render() {
        return (
            <div>
                <button type="button" class="btn btn-danger ml-2">
                    Clear
                </button>
                <div class="d-flex">
                    <SketchField 
                            width='1024px'
                            ref={this.canvasRef}
                            height='768px' 
                            tool={Tools.Pencil} 
                            lineColor='black'
                            lineWidth={3}
                            // onChange={(e) => console.log(this.canvasRef.current.toJSON())}
                            onChange={(evt) => this.setState({path: this.canvasRef.current.toJSON()})}
                    />
                </div>
            </div>
        )
    }
}