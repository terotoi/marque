import React, { useState } from 'react'
import {
	AppBar, IconButton, InputAdornment, Menu, MenuItem, TextField, Toolbar, Typography,
	Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle,
	makeStyles
} from '@material-ui/core'
import { Add, Search } from '@material-ui/icons'
import AccountCircle from '@material-ui/icons/AccountCircle'
import MenuIcon from '@material-ui/icons/Menu'
import { openAlertDialog } from './dialogs/alert'
import { openPasswordDialog } from './dialogs/password'
import { fetchJSON } from './api'

const VERSION = '0.4.1'

const styles = (theme) => {
	const styles = {
		navbar: {
			marginBottom: theme.spacing(2)
		},
		menuButton: {
			marginRight: theme.spacing(2)
		},
		title: {
			flexGrow: 3
		},
		search: {
			flexGrow: 1
		}
	}
	return styles
}

const useStyles = makeStyles(styles)

/**
 * AppBar
 *
 * @param {function} props.onAddClicked - called when add bookmark is clicked
 * @param {function} props.onSearchTextChanged - called when user changed the search text
 * @param {state} ctx - context for the application
 */
const MyAppBar = function (props) {
	const classes = useStyles()
	const [mainMenuAnchor, setMainMenuAnchor] = useState(null)
	const [accountMenuAnchor, setAccountMenuAnchor] = useState(null)
	const [aboutOpen, setAboutOpen] = useState(false)

	function onChangePassword() {
		function onPasswordConfirm(oldPasswd, newPasswd) {
			fetchJSON('/api/user/set_password', 'post', 'json', {
				Username: props.ctx.username,
				OldPassword: oldPasswd,
				NewPassword: newPasswd
			},
				props.ctx.authToken,
				() => {
					openAlertDialog(props.ctx, {
						text: "Password changed."
					})
				},
				(error) => { openAlertDialog(props.ctx, { title: "Error", text: error || "Server error" }) })
		}

		openPasswordDialog(props.ctx, {
			text: 'Change your password.',
			onConfirm: onPasswordConfirm
		})
	}

	return (
		<div>
			<AppBar position="static" className={classes.navbar}>
				<Toolbar>
					<IconButton edge="start" className={classes.menuButton} color="inherit"
						aria-label="menu"
						onClick={(ev) => { setMainMenuAnchor(ev.currentTarget) }}>
						<MenuIcon />
					</IconButton>
					<Typography variant="h4" className={classes.title}>Marque</Typography>

					<TextField
						id="search_field"
						label="Search"
						autoComplete="off"
						name="search"
						className={classes.search}
						onChange={(ev) => props.onSearchTextChanged(ev.target.value)}
						InputProps={{
							startAdornment: (
								<InputAdornment position="start">
									<Search />
								</InputAdornment>
							),
						}} />

					<IconButton color="inherit" aria-label="add a bookmark" onClick={props.onAddClicked}>
						<Add />
					</IconButton>

					<IconButton
						onClick={(ev) => setAccountMenuAnchor(ev.currentTarget)}>
						<AccountCircle />
					</IconButton>

					{/** Main menu **/}
					{mainMenuAnchor === null ? null :
						<Menu
							anchorEl={mainMenuAnchor}
							open={true}
							onClose={() => setMainMenuAnchor(null)}>
							<MenuItem onClick={() => { setAboutOpen(true); setMainMenuAnchor(null) }}>
								About</MenuItem>
						</Menu>}

					{/** Account menu **/}
					{accountMenuAnchor === null ? null :
						<Menu
							anchorEl={accountMenuAnchor}
							open={true}
							onClose={() => setAccountMenuAnchor(null)}>

							<MenuItem>{props.ctx.username ? "Username: " + props.ctx.username : "Not logged in"}</MenuItem>
							<MenuItem onClick={onChangePassword}>Change password</MenuItem>
							{props.ctx.isAdmin ?
								<MenuItem disabled>Manage users</MenuItem> : null}
							<MenuItem onClick={() => {
								setAccountMenuAnchor(null)
								props.ctx.logout()
							}}>Logout</MenuItem>
						</Menu>}
				</Toolbar>
			</AppBar>

			{aboutOpen ? <AboutDialog onClose={() => setAboutOpen(false)} /> : null}
		</div >
	)
}

// props:
//   onClose()
function AboutDialog(props) {
	const onClose = () => {
		props.onClose()
	}

	return (
		<Dialog
			open={true}
			onClose={onClose}
			aria-labelledby="about-dialog-title"
			aria-describedby="about-dialog-description">
			<DialogTitle id="about-dialog-title">Marque</DialogTitle>
			<DialogContent>
				<DialogContentText id="about-dialog-description">
					Marque {VERSION} Copyright © 2020 Tero Oinas.<br />
					This program is published under the Generic Public License v2.0.
				</DialogContentText>
			</DialogContent>
			<DialogActions>
				<Button onClick={onClose} color="primary" autoFocus>
					OK
				</Button>
			</DialogActions>
		</Dialog>
	)
}


export default MyAppBar
