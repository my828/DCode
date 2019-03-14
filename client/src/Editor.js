/* global require */
/* eslint-disable import/no-commonjs */
import React from 'react';
import Prism from 'prismjs';
import Editor from 'react-simple-code-editor';
// import { highlight, languages } from 'prismjs';
// import './prism/prism.css'

const code = "// Welcome to DCode! Copy and paste the link above to hare with friends. Type your code here:";
 
export default class App extends React.Component {
  constructor(props) {
    super(props);
    // this.state = {
    //   sessionID: this.props.state.sessionID,
    //   code: this.props.state.code,
    //   figures: this.props.state.figures
    // }
  }
  componentDidMount() {
    this.setState({code: this.props.state.code})
    console.log(this.state.code)
  }
  // componentWillReceiveProps(props) {
  //   this.setState({
  //     sessionID: props.state.sessionID,
  //     code: props.state.code,
  //     figures: props.state.figures
  //   });
  // }

  state = { code }
  render() {
    return (
        <Editor
        value={this.state.code}
        // onValueChange={code => this.updateCode(code)}
        onValueChange={code => this.setState({ code }, () => {
          this.props.update(this.state)
        })}
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