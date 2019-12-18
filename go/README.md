# Golang sync-code-samples

## Table of Contents
1. [Requirements](#requirements)
2. [Docker](#docker)
3. [Use cases](#use-cases)  
3.1 [Users](#users)  
3.2 [Sessions](#sessions)  
3.3 [Catalogs](#catalogs)  
3.4 [Credentials](#credentials)  
3.5 [Accounts](#accounts)  
3.6 [Transactions](#transactions)  
3.7 [Attachments](#Attachments)  

## Requirements
### Test examples  
In order to test the Golang Sync examples is necessary that you have [Docker](https://www.docker.com/get-started) installed.

### Environment variables required
SYNC_BASEURL=https://sync.paybook.com/v1/  
SYNC_APIKEY=xxxxxxxxx    
SYNC_CREDENTIALS={"username":"test","password":"test"}  
SYNC_IDACCOUNT=xxxxxxx  
SYNC_IDCREDENTIAL=xxxxxxxxx  
SYNC_IDSITE=xxxxxxxxx  
SYNC_IDUSER=xxxxxxxxx  
SYNC_TOKEN=xxxxxxxxx  
SYNC_USERNAME="Sync user"
SYNC_TWOFA=`{"token":"test"}`

To set a env variable in your Linux or MacOS system type

```
export BASEURL=https://sync.paybook.com/v1/
```

To display a env variable type

```
echo $BASEURL
https://sync.paybook.com/v1/
```


## Docker
### Run docker

Change PATHS con volumes

```bash
docker run \  
--rm -it \  
-v $PAYBOOK/git/sync-code-samples/go:/go/src/paybook.com/sync-code-samples \
-w /go \
-e SYNC_BASEURL=https://sync.paybook.com/v1/ \
-e SYNC_APIKEY=xxxxxxxxx \
-e SYNC_CREDENTIALS=`{"username":"test","password":"test"}` \
-e SYNC_IDACCOUNT=xxxxxxxxx \
-e SYNC_IDCREDENTIAL=xxxxxxxxx \
-e SYNC_IDSITE=xxxxxxxxx \
-e SYNC_IDUSER=xxxxxxxxx \
-e SYNC_TOKEN=xxxxxxxxx \
-e SYNC_USERNAME="Usuario de Sync" \
-e SYNC_TWOFA=`{"token":"test"}` \
golang:1.12.1 \
/bin/bash
```

### Install dependencies

```
go get paybook.com/sync-code-samples
```

### Test 

```
go run /go/src/paybook.com/sync-code-samples/examples.go -test
Test Go samples:
SYNC_BASEURL= https://sync.paybook.com/v1/
```

### Build

```
go install paybook.com/sync-code-samples
```

### Run

```
sync-code-samples -test
```

## Use cases
### Users
#### Create
```
export SYNC_USERNAME="Test user"
sync-code-samples -users=create
```

### Get
```
sync-code-samples -users=get
```

### Modify
```
export SYNC_IDUSER=xxxxxx
export SYNC_USERNAME="Test user1"
sync-code-samples -users=modify
```

### Delete
```
export SYNC_IDUSER=xxxxxx
sync-code-samples -usesr=delete
```

## Sessions
### Get token
```
export SYNC_IDUSER=xxxxxx
sync-code-samples -sessions=get_token
```
### Verify token
```
export SYNC_TOKEN =xxxxxx
sync-code-samples -sessions=verify_token
```

## Catalogs
### Get Sites
```
export SYNC_TOKEN=xxxxxxxxxxx  
sync-code-samples -catalogs=sites
```

## Credentials
### Create
```
export SYNC_TOKEN=xxxxxxxxxxx  
export SYNC_IDSITE=xxxxxxxxxxx  
export SYNC_CREDENTIALS=`{"username":"test","password":"test"}`  
sync-code-samples -credentials=create
```

### Get
```
export SYNC_TOKEN=xxxxxxxxxxx  
sync-code-samples -credentials=get
```

### Delete
```
export SYNC_TOKEN=xxxxxxxxxxx  
export SYNC_IDCREDENTIAL=xxxxxxxxxxx  
sync-code-samples -credentials=delete
```

## Accounts
### Get
```
export SYNC_TOKEN=xxxxxxxxxxx  
export SYNC_IDCREDENTIAL=xxxxxxxxxxx 
sync-code-samples -accounts=get
```

## Transactions
### Get
```
export SYNC_TOKEN=xxxxxxxxxxx  
export SYNC_IDCREDENTIAL=xxxxxxxxxxx 
export SYNC_IDACCOUNT=xxxxxxxxxxx 
sync-code-samples -transactions=get
```

## Attachments
### Get
```
export SYNC_TOKEN=xxxxxxxxxxx  
export SYNC_IDCREDENTIAL=xxxxxxxxxxx 
export SYNC_IDACCOUNT=xxxxxxxxxxx 
sync-code-samples -attachments=get
```

## Exit from Docker container
Just type

```
exit
```
