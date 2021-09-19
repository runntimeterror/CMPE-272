import './Tweet.css'

function Tweet(props) {
    const { user: { profile_image_url_https,
        name,
        screen_name },
        text,
        created_at,
        deleteTweet,
        id_str
    } = props
    return <div className="tweet-container">
        <div className="twitter-user">
            <div className="user-icon">
                <img height="50" width="50" src={profile_image_url_https}></img>
            </div>
            <div className="user-name">
                <div className="name">{name}</div>
                <div className="twitter-handle">{`@${screen_name}`}</div>
            </div>
        </div>
        <div className="tweet-text-posted">{text}</div>
        <div className="meta-info">
            <div className="created-at">Created: {created_at}</div>
            <i onClick={() => deleteTweet(id_str)} class="delete-tweet"><img src="/trash-alt-regular.svg" width="20"></img></i>
        </div>
    </div>
}

export default Tweet