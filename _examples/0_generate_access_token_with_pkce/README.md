generate-access-token-with-pkce
===

This is a simple web site to generate access token with PKCE for Twitter App. This site is intended to run in a local environment.

# Usage

1. Execute following command to run the site.

    ```
    docker-compose up
    ```

2. Access to `http://127.0.0.1:8888/auth/` , and follow the instructions in that page.

    #### Step 0: Create Twitter App and create `appconfig.json`

    Access to Twitter Developers Portal and create a new Twitter app. Then, please input following url to OAuth redirect URL form.
    
    ```
    http://127.0.0.1:8888/auth/step2-3/
    ```

    ##### Create `appconfig.json` file
    
    Please create `appconfig.json` file with following format, and put it to `/auth` directory.
    
    ```
    cp auth/appconfig.json.sample auth/appconfig.json
    ```
    
    and edit it.

    #### Step 1: The authorization request
    
    Access the link displayed on the page to request authorization.
    
    #### Step 2: Arrow app, and get authorization result
    
    When you authenticate on the Twitter app authentication screen, you will be redirected to `http://127.0.0.1:8888/auth/step2-3/` and the authentication results will be displayed.
    
    #### Step 3: Exchange token
    
    Execute the curl command shown on the screen to obtain an access token. The result will be output to the `token.json` file.

3. (option) Refresh access token.

    If you have `refresh_token`, the following command can be executed to refresh the access token.
    
    ```bash
    curl -u ${CLIENT_ID}:${CLIENT_SECRET} \
    -X POST 'https://api.twitter.com/2/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode "refresh_token=${REF_TOKEN}" \
    --data-urlencode 'grant_type=refresh_token' \
    --data-urlencode "client_id=${CLIENT_ID}" \
    | jq
    ```