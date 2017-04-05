
import React from 'react';

class Results extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div id="ghuser">
        <h5>{this.state.name}</h5>
        <img src={this.state.avatar} />
        <p>Number of repos: {this.state.repos}</p>
        <p>Git hub <a href={this.state.url}>Link</a></p>
      </div>
    );
  }
}

export default Results;
