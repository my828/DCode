import React from "react";
import Editor from './Editor'
import Canvas from './Canvas'
import Nav from './Nav';

import { BrowserRouter as Router, Route, Link } from "react-router-dom";

export default class Main extends React.Component {
    constructor(prop) {
        super(prop)
        this.state = {
            landing: true
        }
    }

    render() {
        return(
            <div>
                <Nav />
                <div class="d-flex">
                    <Canvas />
                    <Editor />
                </div>
            </div>
        )
    }
}
