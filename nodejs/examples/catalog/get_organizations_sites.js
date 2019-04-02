"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_TOKEN = args['token'];
const IS_TEST = args['is_test'];

const SYNC_API = 'https://sync.paybook.com/v1';

async function main(){
  try{
    
    var options_request = {
      uri : SYNC_API + '/catalogues/organizations/sites?pretty=1' + ( IS_TEST ? '&is_test=1' : ''),
      headers : {
        'Authorization' : 'TOKEN token=' + SYNC_TOKEN,
        'Content-type' : 'application/json'
      },
      method : 'GET',
      json : true
    };

    var result = await request(options_request);
    var response = result.response;
    for(var i=0;i<response.length;i++){
      var organization = response[i].name;
      console.log("Organization " + organization);
      for(var j=0;j<response[i].sites.length;j++){
        var site = response[i].sites[j];
        var site_id = site.id_site;
        var site_name = site.name;
        console.log("\t Site ID: " + site_id);
        console.log("\t Site Name: " + site_name);
        console.log("");
      }
      console.log("");
      console.log("");
    }


    process.exit();
  }catch(error){
    console.log(error.error);
    process.exit();
  }
}

main();
