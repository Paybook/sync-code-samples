"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_API_KEY = args['api_key'];
const ID_USER = args['id_user'];

const SYNC_API = 'https://opendata-api.syncfy.com/v1';

async function main(){
  try{

    var options_request = {
      uri : SYNC_API + '/users/' + ID_USER + '?pretty=1',
      headers : {
        'Authorization' : 'API_KEY api_key=' + SYNC_API_KEY,
        'Content-type' : 'application/json'
      },
      method : 'PUT',
      body : {
        name : 'Homer J Simpson'
      },
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
