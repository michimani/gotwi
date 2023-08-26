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

Plans that can run this example: **Basic** , **Pro**

```bash
cd 1_get_recent_activity \
&& go run . 'your-account-id'
```

### 2. Post a tweet, delete a tweet

Plans that can run this example: **Free**,  **Basic** , **Pro**

```bash
cd 2_post_delete_tweet
```

- post a tweet

  ```bash
  go run . 'post' 'This is a test tweet.'
  ```

- delete a tweet

  ```bash
  go run . 'delete' 'tweet-id'
  ```
  
### 3. Sampling tweets

Plans that can run this example: **Pro** ?

```bash
cd 3_sample_stream \
&& go run . 5
```

### 4. Streaming tweets

Plans that can run this example: **Pro**

see [4_filtered_stream/README.md](./4_filtered_stream/README.md)