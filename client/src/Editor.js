/* @flow */
/* global require */
/* eslint-disable import/no-commonjs */
import React from 'react';
import Prism from 'prismjs';
import Editor from 'react-simple-code-editor';
// import { highlight, languages } from 'prismjs';
import './prism/prism.css'

const code = "// Type your code here:";

  
 
export default class App extends React.Component {
  state = { code };
 
  render() {
    return (
      <Editor
        value={this.state.code}
        onValueChange={code => this.setState({ code })}
        highlight={code => Prism.highlight(code, Prism.languages.javascript, 'javascript')}
        padding={10}
        style={{
          fontFamily: '"Fira code", "Fira Mono", monospace',
          fontSize: 15,
        }}
      />
    );
  }
}