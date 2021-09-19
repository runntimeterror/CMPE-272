import CreateTweet from './Create-Tweet'
import TweetList from './Tweet-List'
import './Demo.css'

function Demo() {
    return <div className="demo-container">
        <div className="legend">
            <span>CMPE-272 - Assignment 2</span>
            <span className="students">
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
                </a></span>
        </div>
        <CreateTweet />
        <TweetList />
    </div>
}

export default Demo
