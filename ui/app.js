import React from 'react'
import { Container, CssBaseline } from '@material-ui/core'
import { createTheme, MuiThemeProvider } from '@material-ui/core/styles'
import MyAppBar from './appbar'
import DatabaseUI from './db'
import LoginView from './login'
import { openAlertDialog} from './dialogs/alert'
import { fetchJSON } from './api'

/** Application theme */
const theme = createTheme({
	spacing: 6,

	palette: {
		type: 'dark',
		primary: {
			main: '#272c34'
		},
		secondary: {
			main: '#c1b685'
		},
		text: {
			primary: '#e0e0e0',
			secondary: '#a0a0a0'
		},
		background: {
			default: '#202020',
			paper: '#525252'
		},
		tonalOffset: 0.0,
		contrastThreshold: 3
	},
	typography: {
		fontSize: 14
	}
})

/**
 * App is the main component of the application.
 * 
 * @param {state} ctx - context for the application
 */
export default function App(props) {
	const database = React.useRef(null)

	const onLogin = (username, password) => {
		fetchJSON('/api/login', 'post', 'json', { Username: username, Password: password }, null, (rs) => {
			if (rs !== false) {
				props.ctx.setUsername(rs.Username)
				props.ctx.setIsAdmin(rs.IsAdmin)
				props.ctx.setAuthToken(rs.AuthToken)
				console.log("Logged in as", rs.Username, "admin:", rs.IsAdmin)
			}
		}, (err) => {
			console.log("Login error:", err)
			openAlertDialog(props.ctx, { text: "Login error." })
		})
	}

	return (
		<MuiThemeProvider theme={theme}>
			<Container maxWidth={false}>
				<CssBaseline />

				{props.ctx.dialogs.map((dialog, i) => <React.Fragment key={i}>{dialog}</React.Fragment>)}

				{(props.ctx.authToken == null) ?
					<LoginView onSubmit={onLogin} /> :
					<React.Fragment>
						<MyAppBar
							onAddClicked={() => database.current.addBookmark()}
							onSearchTextChanged={(text) => database.current.filterText = text}
							ctx={props.ctx} />
						<DatabaseUI ref={(r) => database.current = r} ctx={props.ctx} />
					</React.Fragment>}
			</Container>
		</MuiThemeProvider>
	)
}

