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
            <td>-</td>
        </tr>
        <tr>
            <td>101</td>
            <td>Starting</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>102</td>
            <td>Running</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>103</td>
            <td>TokenReceived</td>
            <td>-</td>
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
            <td>-</td>
        </tr>
        <tr>
            <td>201</td>
            <td>Pending</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>202</td>
            <td>NoTransactions</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>203</td>
            <td>PartialTransactions</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>204</td>
            <td>Incomplete</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>205</td>
            <td>Pending</td>
            <td>-</td>
        </tr>
      	<tr>
            <td>206</td>
            <td>NoAccounts</td>
            <td>-</td>
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

