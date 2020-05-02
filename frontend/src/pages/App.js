import React, { Component } from 'react';
import axios from 'axios';


class App extends Component {

  state = {
    persons: []
  }

  componentDidMount() {
    axios.get(`vehicles`)
      .then(res => {
        const persons = res.data;
        this.setState({ persons });
        console.log(persons)
      })
  }

  render() {
    return (
      <div>
          <h1 className="App-title">Welcome to React</h1>
      </div>
    );
  }
}

export default App;