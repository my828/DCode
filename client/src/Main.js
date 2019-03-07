import React from "react";
import Canvas from './Canvas'
export default class Main extends React.Component {
    constructor(prop) {
        super(prop)
        this.state = {
            landing: true
        }
    }
    handleClick = () => {
        console.log("it works!")
        this.setState({
            landing: !this.state.landing
        })
    }

    render() {

        let font = {
            fontSize: "30px"
        }
        return(
            <div>
                {
                    !this.state.landing && 
                    <div>
                        <header><b class="pb-5" style={font}>HELLO WELCOME TO DCODE :)</b></header>
                        <button type="button" class="btn btn-dark p-2" 
                            onClick={() => this.setState({landing: !this.state.landing})}>
                            Start DCode
                        </button>
                    </div>
                }

                 <Canvas />
                
            </div>
        )
    }
}
