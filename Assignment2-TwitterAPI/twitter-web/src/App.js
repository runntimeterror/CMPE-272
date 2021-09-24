/**
 * File created by Soham Bhattacharjee
 */
import './App.css';
import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom'
import Home from './Home'
import Demo from './Demo'

export default function App() {
  return (
    <header className="App-header">
      <Router>
        <Switch>
          <Route path="/Demo">
            <>
              <HomeNav />
              <Demo />
            </>
          </Route>
          <Route exact path="/">
            <Home />
          </Route>
        </Switch>
      </Router>
    </header>
  );
}

function HomeNav() {
  return <Link className="home-nav-link" to="/">
    <img src="/Logo_white.svg"></img>
  </Link>
}
