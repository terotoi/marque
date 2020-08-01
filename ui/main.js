import React from 'react'
import ReactDOM from 'react-dom'
import App from './app'
import 'typeface-roboto'

(function () {
	var ui = document.getElementById("ui")

	console.log("main.js")

	ReactDOM.render(
		<App />, ui)
})()
