gotwi examples
===
## Preparation

1. Create `.env` file.

  ```bash
  cp .env.sample .env
  ```

  And, replace values of each environment variables.

2. Load environment variables.

  ```bash
  source .env
  ```
  
## Examples

### 1. Get followers' recent activity

```bash
go run  ./1_get_recent_activity/. 'your-account-id'
```

### 2. Post a tweet, delete a tweet

- post a tweet

  ```bash
  go run ./2_post_tweet/. 'post' 'This is a test tweet.'
  ```

- delete a tweet

  ```bash
  go run ./2_post_tweet/. 'delete' 'tweet-id'
  ```