import React from 'react'
import ReactDOM from 'react-dom'
import Context from './context'
import 'typeface-roboto'

(function () {
	var ui = document.getElementById("ui")

	ReactDOM.render(
		<Context />, ui)
})()
