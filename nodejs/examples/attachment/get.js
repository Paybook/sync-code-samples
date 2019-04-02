"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_TOKEN = args['token'];
const SYNC_ID_CREDENTIAL = args['id_credential'];
const SYNC_ID_ACCOUNT = args['id_account'];

const SYNC_API = 'https://sync.paybook.com/v1';

async function main(){
  try{

    var options_request = {
      uri : SYNC_API + '/attachments?pretty=1&id_credential=' + SYNC_ID_CREDENTIAL + "&id_account=" + SYNC_ID_ACCOUNT,
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
