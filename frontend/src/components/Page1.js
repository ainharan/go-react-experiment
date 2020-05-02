import React, {Component} from 'react';

import axios from 'axios';

export default class Page1 extends Component {
  state = {
    persons: []
  }

  componentDidMount() {
    axios.get(`/vehicles`)
      .then(res => {
        const persons = res.data;
        this.setState({ persons });
        console.log(persons)
      })
  }

  render() {
    return (
      <ul>
        { this.state.persons.map(person => <li>{person.name}</li>)}
      </ul>
    )
  }
}