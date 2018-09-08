import React from 'react';
import ResultsTable from './ResultsTable';

export default class ConnectedResultsTable extends React.Component {
  state = { 
    results: []
  }

  async componentDidMount() {
    try {
      const response = await fetch('http://localhost:8080/results');
      if (!response.ok) {
        throw new Error(response.statusText);
      }
      const json = await response.json();
      this.setState(json);
    } catch (error) {
       console.log(error);
    }
    
  }

  render() {
    return <ResultsTable results={this.state.results} />;
  }
}