import React from "react";
import CanvasDraw from "react-canvas-draw";

export default class Canvas extends React.Component {
    constructor(prop) {
        super(prop) 
        this.state = {
            width: 700,
            height: 900,
            brushRadius: 6
          };
    }

    componentDidMount() {
         // let's change the color randomly every 2 seconds. fun!
         window.setInterval(() => {
            this.setState({
              color: "#" + Math.floor(Math.random() * 16777215).toString(16)
            });
          }, 2000);
         console.log("did mount")
    }
    render() {
        return (
            <div>
                <CanvasDraw canvasWidth={this.state.width} canvasHeight={this.state.height} brushRadius={this.state.brushRadius}/>
            </div>
        )
    }
}