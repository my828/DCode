import React from "react";

export default class Main extends React.Component {
    render() {
        let mainstyle = {
            position: "relative",
            top: "300px"
        }
        return(
            <div class="justify-content-center" style={mainstyle}>
                
                <button type="button" class="btn btn-dark p-2">Start DCode</button>
            </div>
        )
    }
}
