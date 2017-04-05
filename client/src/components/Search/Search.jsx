import React from 'react';

class Search extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return(
      <form action="/" method="POST">
        <input type="submit" value="search" />
      </form>
    );
  }
}

export default Search;
