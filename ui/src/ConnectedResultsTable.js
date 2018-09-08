import React from 'react';
import ResultsTable from './ResultsTable';
import Pusher from 'pusher-js';

const socket = new Pusher('ce62ea44549ce3be69ef', {
  cluster: 'us2',
  encrypted: true
});

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

      const channel = socket.subscribe('results');
      channel.bind('results', (data) => {
        this.setState(data);
      })
    } catch (error) {
       console.log(error);
    }
    
  }

  render() {
    return <ResultsTable results={this.state.results} />;
  }
}