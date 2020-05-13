# Status Codes

------

Credential Status Codes

------

<br/>

### Progress Information Codes

<table>
    <thead>
        <tr>
            <th colspan="3">1xx</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><b>Code</b></td>
            <td><b>Name</b></td>
            <td><b>Description</b></td>
        </tr>
        <tr>
            <td>100</td>
            <td>Register</td>
            <td>The API registers a new process (through a REST request)</td>
        </tr>
        <tr>
            <td>101</td>
            <td>Starting</td>
            <td>The process information was obtained to start operating</td>
        </tr>
      	<tr>
            <td>102</td>
            <td>Running</td>
            <td>The process is running (login successful)</td>
        </tr>
      	<tr>
            <td>103</td>
            <td>TokenReceived</td>
            <td>The process received the token</td>
        </tr>
    </tbody>
</table> 

------

<br/>

### Success Codes

<table>
    <thead>
        <tr>
            <th colspan="3">2xx</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><b>Code</b></td>
            <td><b>Name</b></td>
            <td><b>Description</b></td>
        </tr>
        <tr>
            <td>200</td>
            <td>Finish</td>
            <td>The data was obtained correctly</td>
        </tr>
        <tr>
            <td>201</td>
            <td>Pending</td>
            <td>The data was processed correctly, pending data will continue to download in the background</td>
        </tr>
      	<tr>
            <td>202</td>
            <td>NoTransactions</td>
            <td>The process ends successfully, but no transactions were found</td>
        </tr>
      	<tr>
            <td>203</td>
            <td>PartialTransactions</td>
            <td>The process ends successfully, but one or more accounts have no transactions</td>
        </tr>
      	<tr>
            <td>204</td>
            <td>Incomplete</td>
            <td>The job completed successfully, but the information is incomplete</td>
        </tr>
      	<tr>
            <td>205</td>
            <td>Pending</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>206</td>
            <td>NoAccounts</td>
            <td>The process ends successfully, but no Accounts were found</td>
        </tr>
    </tbody>
</table> 

------

<br/>

### User Interaction Codes

<table>
    <thead>
        <tr>
            <th colspan="3">4xx</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><b>Code</b></td>
            <td><b>Name</b></td>
            <td><b>Description</b></td>
        </tr>
        <tr>
            <td>401</td>
            <td>Unauthorized</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>403</td>
            <td>Invalidtoken</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>405</td>
            <td>Locked</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>406</td>
            <td>Conflict</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>408</td>
            <td>UserAction</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>409</td>
            <td>WrongSite</td>
            <td>-</td>
        </tr>
    </tbody>
</table> 

------

<br/>

### Multi-Factor Authentication Codes

<table>
    <thead>
        <tr>
            <th colspan="3">4xx</th>
        </tr>
    </thead>
    <tbody>
      	<tr>
            <td><b>Code</b></td>
            <td><b>Name</b></td>
            <td><b>Description</b></td>
        </tr>
        <tr>
            <td>410</td>
            <td>Waiting</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>411</td>
            <td>TwofaTimeout</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>413</td>
            <td>LoginError</td>
            <td>-</td>
        </tr>
    </tbody>
</table> 

------

<br/>

### Connection Error Codes

<table>
    <thead>
        <tr>
            <th colspan="3">5xx</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><b>Code</b></td>
            <td><b>Name</b></td>
            <td><b>Description</b></td>
        </tr>
        <tr>
            <td>500</td>
            <td>Error</td>
            <td>-</td>
        </tr>
        <tr>
            <td>501</td>
            <td>Unavailable</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>504</td>
            <td>ConnectionTimeout</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>509</td>
            <td>UndergoingMaintenance</td>
            <td>-</td>
        </tr>
    </tbody>
</table> 

