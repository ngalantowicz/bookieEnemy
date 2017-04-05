import React from 'react';
import {render} from 'react-dom';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import Search from './components/Search/Search.jsx';
import Results from './components/Results/Results.jsx';

render((
  <div>
    <Router>
      <div>
        <h1>Bookie Enemy</h1>
        <Route exact path="/" component={Search} />
        <Route exact path="/result" component={Results} />
      </div>
    </Router>
  </div>
), document.getElementById('app'));
