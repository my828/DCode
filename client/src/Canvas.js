import React from "react";
import CanvasDraw from "react-canvas-draw";

export default class Canvas extends React.Component {
    constructor(prop) {
        super(prop) 
        this.state = {
            width: window.screen.availWidth / 2 - 80,
            height: 900,
            brushRadius: 5
          };
    }

    componentDidMount() {
         // let's change the color randomly every 2 seconds. fun!
        //  window.setInterval(() => {
        //     this.setState({
        //       color: "#" + Math.floor(Math.random() * 16777215).toString(16)
        //     });
        //   }, 2000);
        //  console.log("did mount")
    }

    render() {

        return (
            <div>
                <button type="button" class="btn btn-danger ml-2" onClick={() => this.saveableCanvas.clear()}>
                    Clear
                </button>
                <button type="button" class="btn btn-secondary ml-3" onClick={() => this.saveableCanvas.undo()}>
                    Undo
                </button>
                <CanvasDraw ref={canvasDraw => (this.saveableCanvas = canvasDraw)} canvasWidth={this.state.width} canvasHeight={this.state.height} 
                    brushRadius={this.state.brushRadius} onChange={console.log("change")}/>
            </div>
        )
    }
}