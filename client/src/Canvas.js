import React from "react";
import {SketchField, Tools} from 'react-sketch';

export default class Canvas extends React.Component {
    constructor(props) {
        super(props) 
        this.state = {
            width: window.screen.availWidth / 2 - window.screen.availWidth / 5,
            pageState: props.state
        };
        this.canvasRef = React.createRef();
        this.mounted = true;
    }

    componentWillReceiveProps(props) {
        this.setState({
            pageState: props.state
        });
    }

    // handles changes to canvas
    handleCanvasChange(evt) {
        if (evt && (evt.type === "mouseup" || evt.type === "mouseout")) {
            let figures = this.canvasRef.current.toJSON();
            let pageState = this.state.pageState;
            let path = window.location.href;
            let components = path.split("/");
            let sessionID = components[components.length - 1];
            pageState.sessionID = sessionID;
            pageState.figures = figures;
            this.props.updateState(pageState);
        }
    }

    render() {
        return (
            <div>
                <button type="button" className="btn btn-danger ml-2 canvas-clear">
                    clear
                </button>
                <div className="d-flex">
                    <SketchField 
                            width='1024px'
                            ref={this.canvasRef}
                            height='768px'
                            tool={Tools.Pencil} 
                            lineColor='blue'
                            lineWidth={3}
                            onChange={(evt) => {this.handleCanvasChange(evt)}}
                            value={this.state.pageState.figures}
                    />
                </div>
            </div>
        )
    }
}