media upload example
===

This is sample code that uses the `media/upload` package to upload media files to Twitter.

# Steps to upload media

## 1. Initialize media upload

First, initialize the media upload with the following code.

```go
res, err := Initialize(client, &types.InitializeInput{
    MediaType:     types.MediaTypePNG,
    TotalBytes:    len(fileBytes),
    Shared:        false,
    MediaCategory: types.MediaCategoryTweetImage,
})
```

## 2. Append media data

You can append media data in two ways:

### Single segment upload
```go
appendRes, err := Append(client, &types.AppendInput{
    MediaID:      mediaID,
    Media:        bytes.NewReader(fileBytes),
    SegmentIndex: 0,
})
```

### Multi-segment upload

In this example, the file is divided into 10 chunks for multi-segment upload.


```go
chunkSize := len(fileBytes) / 10
segmentIndex := 0
for i := 0; i < len(fileBytes); i += chunkSize {
    end := min(i+chunkSize, len(fileBytes))
    chunk := fileBytes[i:end]
    appendRes, err := Append(client, &types.AppendInput{
        MediaID:      mediaID,
        Media:        bytes.NewReader(chunk),
        SegmentIndex: segmentIndex,
    })
    segmentIndex++
}
```

## 3. Finalize media upload

After appending all media data, finalize the upload.

```go
finalizeRes, err := Finalize(client, &types.FinalizeInput{
    MediaID: mediaID,
})
```

## 4. Post tweet with media (optional)

You can post a tweet with the uploaded media.

```go
postedID, err := PostWithMedia(client, "post with a media by using gotwi", mediaID)
```

# Run example code

1. Set environment variables.

    ```bash
    export GOTWI_API_KEY=your-api-key
    export GOTWI_API_KEY_SECRET=your-api-key-secret
    export GOTWI_ACCESS_TOKEN=your-access-token
    export GOTWI_ACCESS_TOKEN_SECRET=your-access-token-secret
    ```
    
2. Run the example with default settings (single segment upload).

    ```bash
    go run .
    ```
    
3. Run with multi-segment upload.

    ```bash
    go run . multi
    ```
    
4. Run with multi-segment upload and post a tweet.

    ```bash
    go run . multi post
    ```

# Notes

- The example uses a sample PNG file (`sample.png`) for demonstration.
- Make sure you have the correct file permissions and the file exists in the same directory.
- The media upload process consists of three steps: Initialize, Append, and Finalize.
- You can choose between single-segment and multi-segment upload based on your needs. 