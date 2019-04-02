"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_API_KEY = args['api_key'];

const SYNC_API = 'https://sync.paybook.com/v1';

async function main(){
  try{

    var options_request = {
      uri : SYNC_API + '/users?pretty=1',
      headers : {
        'Authorization' : 'API_KEY api_key=' + SYNC_API_KEY,
        'Content-type' : 'application/json'
      },
      method : 'GET',
      json : true
    };

    var result = await request(options_request);
    console.log(result);

    process.exit();
  }catch(error){
    console.log(error.error);
    process.exit();
  }
}

main();