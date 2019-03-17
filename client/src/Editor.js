/* global require */
/* eslint-disable import/no-commonjs */
import React from 'react';
import Prism from 'prismjs';
import Editor from 'react-simple-code-editor';
import { highlight, languages } from 'prismjs';
import "./index.css";

const code = "// Welcome to DCode! Copy and paste the link above to hare with friends. Type your code here:";
 
export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      pageState: this.props.state
    }
  }

  componentWillReceiveProps(props) {
    this.setState({
      pageState: props.state
    });
  }

  handleEditorChange(code) {
    let state = this.state.pageState;
    state.code = code;
    this.props.updateState(state);
  }

  state = { code }
  render() {
    
    return (
        <Editor
          className={"editor"}
          value={this.state.pageState.code}
          onValueChange={code => {this.handleEditorChange(code)}}
          highlight={code => Prism.highlight(code, Prism.languages.javascript, 'javascript')}
          padding={5}
          style={{
            fontFamily: 'Fira code, Fira Mono, monospace',
            fontSize: '16px',
            color: '#fff',
          }}
      />
    );
  }
}