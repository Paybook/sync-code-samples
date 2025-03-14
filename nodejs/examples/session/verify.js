"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2));

const SYNC_TOKEN = args['token'];

const SYNC_API = 'https://opendata-api.syncfy.com/v1';

async function main(){
  try{

    var options_request = {
      uri : SYNC_API + '/sessions/'+ SYNC_TOKEN +'/verify?pretty=1',
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
