"use strict"
import { Container, CssBaseline } from '@material-ui/core'
import { teal } from '@material-ui/core/colors'
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles'
import React from 'react'
import MyAppBar from './appbar'
import DatabaseUI from './dbui'

const theme = createMuiTheme({
	palette: { primary: teal, type: 'dark' }
})

/** App is the main component of the application. **/
class App extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
		}

		this.databaseUI = null
	}

	render() {
		return (
			<MuiThemeProvider theme={theme}>
				<Container maxWidth={false}>
					<CssBaseline />
					<MyAppBar
						onAddClicked={() => this.databaseUI.addBookmark()}
						onSearchTextChanged={(text) => this.databaseUI.filterText = text} />
					<DatabaseUI ref={(r) => this.databaseUI = r} />

				</Container>
			</MuiThemeProvider>
		);
	}
}

export default App

