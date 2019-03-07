import React from "react";
import Editor from "./Editor"
export default class Main extends React.Component {
    constructor(prop) {
        super(prop)
        this.state = {
            landing: true
        }
    }
    render() {
        let mainstyle = {
            position: "relative",
            top: "300px"
        }
        return(
            <div>
            {
                this.state.landing && 
                <div class="justify-content-center" style={mainstyle}>  
                <button type="button" class="btn btn-dark p-2" onClick={()=>this.setState({landing: !this.state.landing})}>Start DCode</button>
                </div>
            }
            { 
                !this.state.landing && 
                <div class="justify-content-center" style={mainstyle}>  
                    <button type="button" class="btn btn-primary p-2" onClick={()=>this.setState({landing: !this.state.landing})}>Start DCode</button>
                    <Editor /> 
                </div>
            }
            </div>
        )
    }
}
