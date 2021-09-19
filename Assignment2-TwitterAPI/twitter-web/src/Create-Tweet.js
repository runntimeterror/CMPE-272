import './Create-Tweet.css'

function CreateTweet() {
    return <div className="create-tweet-wrapper">
        <div className="user-icon">
            <img height="50" width="50" src="/user-icon.png"></img>
        </div>
        <div className="tweet-text">
            <div contentEditable="true">What's happening?</div>
            <div className="tweet-button">Tweet</div>
        </div>
    </div>
}

export default CreateTweet
