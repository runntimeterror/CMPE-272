/**
 * File created by Rajat Banerjee
 */
import Tweet from './Tweet'

function TweetList(props) {
    const { error, isLoaded, items, deleteTweet } = props

    if (error) {
        return <div>Error: {error.message}</div>
    } else if (!isLoaded || items.length === 0) {
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
