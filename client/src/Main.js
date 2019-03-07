import React from "react";

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
        let mainstyle = {
            position: "relative",
            top: "300px"
        }
        let font = {
            fontSize: "30px"
        }
        return(
            <div style={mainstyle}>
                {
                    this.state.landing && 
                    <div>
                        <header><b class="pb-5" style={font}>HELLO WELCOME TO DCODE :)</b></header>
                        <button type="button" class="btn btn-dark p-2" 
                            onClick={() => this.setState({landing: !this.state.landing})}>
                            Start DCode
                        </button>
                    </div>
                }
                {
                    !this.state.landing && 
                    <div>
                        <button type="button" class="btn btn-primary p-2" 
                        onClick={() => this.setState({landing: !this.state.landing})}>
                            IT WORKS
                        </button>
                        <div>
                        </div>
                        <div>
                            
                        </div>
                    </div>
                }
            </div>
        )
    }
}
