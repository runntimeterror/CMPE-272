/**
 * File created by Soham Bhattacharjee
 */
import logo from './logo.svg';
import './Home.css'
import { Link } from 'react-router-dom'

function Home() {
    return (<>
        <img src={logo} className="App-logo" alt="logo" />
        <h1>CMPE-272 - Assignment 2</h1>
        <p>
            <Link className="demo-link" to="/demo">Twitter Service Demo</Link>
        </p>
        <div className="student-emails">
            <a
                className="App-link"
                href="mailto:soham.bhattacharjee@sjsu.edu"
            >
                Soham Bhattacharjee
            </a>&nbsp;&amp;&nbsp;
            <a
                className="App-link"
                href="mailto:rajat.banerjee@sjsu.edu"
            >
                Rajat Banerjee
            </a>
        </div>
    </>)
}

export default Home
