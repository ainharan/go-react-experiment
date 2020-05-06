import React, { Component } from 'react';
import "bootstrap/dist/css/bootstrap.css";
import { Button, Form } from 'react-bootstrap';
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
        <Form>
          <Form.Group controlId="formBasicEmail">
            <Form.Label>Email address</Form.Label>
            <Form.Control type="email" />
          </Form.Group>
          <Form.Group controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Password" />
          </Form.Group>
          <Form.Group controlId="formBasicCheckbox">
            <Form.Check type="checkbox" label="Manual" />
          </Form.Group>
          <Button variant="primary" type="submit">
            Submit
          </Button>
        </Form>
        {this.state.persons.map(person => <li key={person.id}>{person.name}</li>)}
      </div>
    );
  }
}

export default App;