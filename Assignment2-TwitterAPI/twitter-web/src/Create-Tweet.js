import './Create-Tweet.css'

function CreateTweet(props) {
    return <div className="create-tweet-wrapper">
        <div className="user-icon">
            <img height="50" width="50" src="/user-icon.png"></img>
        </div>
        <div className="tweet-text">
            <div className="textarea" id="tweet-text" contentEditable="true">{props.tweetText}</div>
            <div onClick={props.postTweet} className="tweet-button">Tweet</div>
        </div>
    </div>
}

export default CreateTweet
