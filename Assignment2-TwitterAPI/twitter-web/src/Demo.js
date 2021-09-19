import CreateTweet from './Create-Tweet'
import TweetList from './Tweet-List'
import { useState, useEffect } from 'react'
import './Demo.css'

function Demo() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);
    const [tweetText, setTweetText] = useState("What's happening?")

    const service_url = `http://localhost:8080/tweet`

    const refreshList = () => {
        fetch(service_url)
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result);
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }

    useEffect(() => {
        refreshList()
    }, [])

    const deleteTweet = (tweetId) => {
        fetch(`${service_url}/${tweetId}`, {
            method: `DELETE`
        }).then(() => {
            refreshList()
        })
    }
    const postTweet = () => {
        const status = document.getElementById('tweet-text').innerText
        setTweetText(status)
        fetch(service_url, {
            method: `POST`,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ status })
        }).then(() => {
            refreshList()
            setTweetText("What's happening?")
        })
    }


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
        <CreateTweet postTweet={postTweet}
            tweetText={tweetText} />
        <TweetList items={items}
            deleteTweet={deleteTweet}
            error={error}
            isLoaded={isLoaded} />
    </div>
}

export default Demo
