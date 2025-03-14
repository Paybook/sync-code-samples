"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_TOKEN = args['token'];
const SYNC_ID_ATTACHMENT = args['id_attachment'];

const SYNC_API = 'https://opendata-api.syncfy.com/v1';

async function main(){
  try{
    var options_request = {
      uri : SYNC_API + '/attachments/' + SYNC_ID_ATTACHMENT,
      headers : {
        'Authorization' : 'Bearer ' + SYNC_TOKEN,
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
