import React, { Component } from 'react';

class Widget extends Component {

	constructor(props) {
	    super(props);

	    this.state = {
	      token: ""
	    };

		if(!window.syncWidget){
			(function(w,d,s,id,r){
				var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?"http":"https";
				if(!d.getElementById(id)){
					w[r]={};
					w[r]=w[r]||function(){w[r].q=w[r].q||[].push(arguments)};
					js=d.createElement(s);
					js.id=id;
					js.type = 'text/javascript';
					js.src=p+"://www.paybook.com/sync/widget.js";
					fjs.parentNode.insertBefore(js,fjs);
				}
			})(window,document,"script","sync-widget", "syncWidget");
		}

	}

  componentDidMount() {
		window.syncWidget.options = {
			baseDiv: "sync_container",
			callback: this.userDefinedCallback,
			theme: "light",
			avoidAdmin: true,
			quickAnswer : true,
		}
  }

  componentWillUnmount() {
		if (!window.syncWidget) return false;
		delete window.syncWidget;
	}

	userDefinedCallback = (args) => {
		console.log(args);
	}

	handleChange(event) {
		this.setState({ [event.target.name]: event.target.value });
	}

	handleSubmit = async event => {
		event.preventDefault();
		var token = this.state.token;
		if (typeof window.syncWidget.setToken === "function") {
			window.syncWidget.setToken(token);
		}
	}

	render() {
		return (
			<div>
				<div>
					<span>Token</span>
					<input type="text" name="token" value={this.state.token} onChange={event => this.handleChange(event)}></input>
					<button className="btn btn-primary" onClick={this.handleSubmit}>Acceder</button>
				</div>
				<div id="sync_container" />
			</div>
		);
	}
}


export default Widget;
