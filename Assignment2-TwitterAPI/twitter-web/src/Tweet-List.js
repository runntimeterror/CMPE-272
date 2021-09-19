import { useState, useEffect } from 'react'
import Tweet from './Tweet'

function TweetList() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

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

    if (error) {
        return <div>Error: {error.message}</div>
    } else if (!isLoaded) {
        return <img src="/loading.gif" width="30" />
    } else {
        return (
            <div className="tweet-list-container">

                {items.map(item => (
                    <Tweet deleteTweet={deleteTweet} {...item} />
                ))}
            </div>
        )
    }
}

export default TweetList
