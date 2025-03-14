"use strict";

const request = require('request-promise');
const args = require('minimist')(process.argv.slice(2))

const SYNC_TOKEN = args['token'];
const SYNC_ID_SITE = args['id_site'];
const SYNC_CREDENTIALS = JSON.parse(args['credentials']);

const SYNC_API = 'https://opendata-api.syncfy.com/v1';

function delay(ms){
  return new Promise(function(resolve){
    setTimeout(resolve,ms);
  })
}

async function main(){
  try{

    var options_request = {
      uri : SYNC_API + '/credentials?pretty=1',
      headers : {
        'Authorization' : 'Bearer ' + SYNC_TOKEN,
        'Content-type' : 'application/json'
      },
      method : 'POST',
      body : {
        id_site : SYNC_ID_SITE,
        credentials : SYNC_CREDENTIALS
      },
      json : true
    };

    var result = await request(options_request);
    var id_credential = result.response.id_credential;
    var url_job_status = result.response.status;

    console.log("");
    console.log("Credential ID : " + result.response.id_credential);
    console.log("URL Job Status : " + url_job_status);

    var exit = false;
    var job_code;
    while(!exit){
      await delay(1000);

      var options_request = {
        uri : url_job_status,
        headers : {
          'Authorization' : 'Bearer ' + SYNC_TOKEN,
          'Content-type' : 'application/json'
        },
        method : 'GET',
        json : true
      };

      var result = await request(options_request);
      var status_array = result.response;
      var last_status = status_array[status_array.length-1];
      job_code = last_status.code;

      console.log("");
      console.log("Code : " + job_code);
      switch (job_code) {
        case 100:
          console.log("Status : Register");
          console.log("Message : API register a new job");
          break;
        case 101:
          console.log("Status : Starting");
          console.log("Message : Sync got job information to start working");
          break;
        case 102:
          console.log("Status : Running");
          console.log("Message : Sync is running (login successful)");
          break;
        case 103:
          console.log("Status : TokenReceived");
          console.log("Message : Sync received the token");
          break;
        case 200:
          console.log("Status : Finish");
          console.log("Message : Data was processed correctly");
          exit = true;
          break;
        case 201:
          console.log("Status : Pending");
          console.log("Message : Data was processed correctly, pending data will continue to download in the background");
          exit = true;
          break;
        case 202:
          console.log("Status : NoTransactions");
          console.log("Message : Job finish correctly, but there where no transactions found");
          exit = true;
          break;
        case 203:
          console.log("Status : PartialTransactions");
          console.log("Message : Job finish correctly, but one or more accounts have no transactions");
          exit = true;
          break;
        case 204:
          console.log("Status : Incomplete");
          console.log("Message : Job finish correctly, but data is incomplete");
          exit = true;
          break;
        case 206:
          console.log("Status : NoAccounts");
          console.log("Message : There are no accounts");
          exit = true;
          break;
        case 401:
          console.log("Status : Unauthorized");
          console.log("Message : Invalid credentials (user and password are not valid)");
          exit = true;
          break;
        case 403:
          console.log("Status : Invalidtoken");
          console.log("Message : Invalid token");
          exit = true;
          break;
        case 405:
          console.log("Status : Locked");
          console.log("Message : Credentials are locked");
          exit = true;
          break;
        case 406:
          console.log("Status : Conflict");
          console.log("Message : User is already logged");
          exit = true;
          break;
        case 408:
          console.log("Status : UserAction");
          console.log("Message : User action required on the site");
          exit = true;
          break;
        case 409:
          console.log("Status : WrongSite");
          console.log("Message : Selected site is incorrect type for credentials");
          exit = true;
          break;
        case 410:
          console.log("Status : Waiting");
          console.log("Message : Waiting for two-fa");
          await send_twofa(last_status.address,last_status.twofa);
          break;
        case 411:
          console.log("Status : TwofaTimeout");
          console.log("Message : Timeout for user input on two-fa");
          exit = true;
          break;
        case 500:
          console.log("Status : Error");
          console.log("Message : Bank requires attention (Sync error)");
          exit = true;
          break;
        case 501:
          console.log("Status : Unavailable");
          console.log("Message : Bank temporarily unavailable (timeout)");
          exit = true;
          break;
        case 504:
          console.log("Status : ScriptTimeout");
          console.log("Message : Sync timeout");
          exit = true;
          break;
        case 509:
          console.log("Status : UndergoingMaintenance");
          console.log("Message : Site is undergoing maintenance");
          exit = true;
          break;
        default:
          throw {error : "Unknown status"};
          exit = true;
          break;
      }
    }

    console.log("");
    console.log("Connection finished with code: " + job_code);

    process.exit();
  }catch(error){
    console.log(error);
    process.exit();
  }
}


function get_twofa_value_from_user(twofa_elements){
  return new Promise(function(resolve, reject){
    try{
      if(twofa_elements.length != 1){
        reject("The length of the elements is not as expected");
      }
      var twofa_element_details = twofa_elements[0];

      var twofa_element_type = twofa_element_details.type;
      console.log("TWO-FA Element Type: " + twofa_element_type);

      var twofa_element_label = twofa_element_details.label;
      console.log("TWO-FA Element Label: " + twofa_element_label);

      var twofa_element_value;
      switch (twofa_element_type) {
        case "text": //it can be TOKEN or CAPTCHA
          if('imgBase64File' in twofa_element_details){
            console.log("It's a Captcha TWO-FA Element");
          }else{
            console.log("It's a Token TWO-FA Element");
          }
          twofa_element_value = "test";
          break;
        case "options":
          console.log("It's a Multi Image TWO-FA Element");
          console.log("There are " + twofa_element_details.options.length + " images");
          twofa_element_value = 1;
          break;
        case "textOptions":
          console.log("It's a Multi Text TWO-FA Element"); 
          console.log("There are " + twofa_element_details.options.length + " options");
          twofa_element_value = 1;
          break;
        default:
          reject("Unknown TWO-FA Element Type");
          break;
      }

      


      var twofa_element_key = twofa_element_details.name;
      console.log("TWO-FA Element Key: " + twofa_element_key);
      console.log("TWO-FA Element Value: " + twofa_element_value);

      var twofa_body = {
        twofa : {}
      };
      twofa_body.twofa[twofa_element_key] = twofa_element_value;

      resolve(twofa_body);
    }catch(error){
      reject(error);
    }
  });

}

function send_twofa(twofa_address,twofa_elements){
  return new Promise(async function(resolve,reject){
    try{

      var twofa_body = await get_twofa_value_from_user(twofa_elements);

      var options_request = {
        uri : twofa_address,
        headers : {
          'Authorization' : 'Bearer ' + SYNC_TOKEN,
          'Content-type' : 'application/json'
        },
        method : 'POST',
        body : twofa_body,
        json : true
      };

      var result = await request(options_request);

      if(result.response){
          resolve(true);
      }else{
        reject({error: "Error"});
      }


    }catch(error){
      reject(error);
    }
  })
}

main();
