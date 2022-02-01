CHANGELOG 
===
This is a version of CHANGELOG less than v1.0.0

## [Unreleased]

### Support APIs
* `GET /2/compliance/jobs/:id`
* `GET /2/compliance/jobs`
* `POST /2/compliance/jobs`

v0.10.0 (2022-02-01)
===

### Features
* Handling API errors


v0.9.10 (2022-01-14)
===

### New Supported APIs
* `GET /2/spaces/:id/buyers`
* `GET /2/spaces/:id/tweets`

v0.9.9 (2022-01-14)
===

### Documentation
* add some tests

### Fixes
* use `io.ReadAll` instead of `ioutil.ReadAll` 
* comment for API (`users.UserLookupMe`)

v0.9.8 (2021-12-28)
===

### New Supported APIs
* `GET /2/users/me`

### Documentation
* add some tests

v0.9.7 (2021-12-24)
===

### Fixes
* remove some unnecessary processing

### Documentation
* add some tests

v0.9.6 (2021-12-24)
===

### Documentation
* add code coverage tool

v0.9.5 (2021-12-01)
===

### Fixes
* Creating OAuth 1.0 signature
* some tests

v0.9.4 (2021-11-22)
===

### Documentation
* add examples
* add pkg.go badge

v0.9.3 (2021-11-22)
===

### New Supported APIs
* `DELETE /2/tweets`
* `POST /2/tweets`

v0.9.2 (2021-11-22)
===

### New Supported APIs
* `GET /2/users/:id/list_memberships`
* `GET /2/users/:id/followed_lists`
* `GET /2/lists/:id/followers`
* `GET /2/users/:id/pinned_lists`
* `GET /2/lists/:id/tweets`

v0.9.1 (2021-11-18)
===

### New Supported APIs
* `GET /2/lists/:id/members`

v0.9.0 (2021-11-18)
===

### New Supported APIs
* `GET /2/users/:id/owned_lists`
* `GET /2/lists/:id`

### Fixes
* type of resource fields
* name of Lists resources


v0.8.2 (2021-11-16)
===

### New Supported APIs
* `GET /2/tweets/search/stream/rules`

v0.8.1 (2021-11-12)
===

### New Supported APIs
* `GET /2/spaces/by/creator_ids`
* `GET /2/spaces`
* `GET /2/spaces/:id`
* `GET /2/spaces/search`

### Fixes
* some fields

v0.8.0 (2021-10-29)
===

### Features
* Call API with context

### Fixes
* Use json decoder

v0.7.0 (2021-10-27)
===

### Fixes
* type of some fields

v0.6.0 (2021-10-27)
===

### New Supported APIs
* `POST DELETE /2/users/:id/pinned_lists`
* `POST DELETE /2/lists/:id/follows`
* `POST DELETE /2/lists/:id/members`
* `DELETE /2/lists/:id`
* `PUT /2/lists/:id`
* `POST /2/lists`
* `PUT /2/tweets/:id/hidden`

### Fixes
* type of JSON parameters
* not ok error struct

v0.5.2 (2021-10-21)
===

### New Supported APIs
* `DELETE /2/users/:id/retweets/:source_tweet_id`
* `POST /2/users/:id/retweets`
* `DELETE /2/users/:id/likes`
* `POST /2/users/:id/likes`
* `GET /2/tweets/:id/liked_tweets`
* `DELETE /2/users/:source_user_id/muting/:target_user_id`
* `POST /2/users/:id/muting`
* `DELETE /2/users/:source_user_id/blocking/:target_user_id`
* `POST /2/users/:id/blocking`
* `DELETE /2/users/:source_user_id/following/:target_user_id`
* `POST /2/users/:id/following`
* `GET /2/tweets/:id/liking_users`
* `GET /2/users/:id/retweeted_by`
* `GET /2/users/:id/muting`
* `GET /2/users/:id/mentions`
* `GET /2/users/:id/tweets`

v0.5.1 (2021-10-18)
===

### Fixes
* Creating new client
* client method, OAuth 1.0 method

v0.5.0 (2021-10-17)
===

### New Supported APIs
* `GET /2/users/:id/blocking`
* `GET /2/tweets/counts/all`
* `GET /2/tweets/counts/recent`
* `GET /2/tweets/search/all API`

### Features
* Support OAuth 1.0a

### Fixes
* Resolving endpoint
* ParameterMap method
* name of some structs

### Documentation

v0.4.2 (2021-10-14)
===

### Features
* Handling rate limit error

v0.4.1 (2021-10-13)
===

### Features
* Handling not 200 errors

v0.4.0 (2021-10-12)
===

### New Supported APIs
* GET /2/tweets/search/recent

v0.3.0 (2021-10-11)
===

### New Supported APIs
* `GET /2/users/:id/followers`
* `GET /2/users/:id/following`

### Fixes
* Name of some files


v0.2.0 (2021-10-09)
====

### New Supported APIs
* `GET /2/tweets/:id`
* `GET /2/tweets`

### Features
* Support partial error
* Handle non 200 error

### Fixes
* Calling Twitter API
* User Lookup APIs
* Directory struct

### Documentation
* add CREDITS
* add LICENCE

v0.1.0 (2021-10-08)
====

* dev release 🚀