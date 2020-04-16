"use strict"
import {
	AppBar, IconButton, InputAdornment, Menu, MenuItem, TextField, Toolbar, Typography,
	Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle,
	makeStyles
} from '@material-ui/core'
import { Add, Search } from '@material-ui/icons'
import MenuIcon from '@material-ui/icons/Menu'
import React, { useState } from 'react'

const VERSION = '0.4'

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

// props:
//   onAddClicked: function()
//   onSearchTextChanged: function(text)
const MyAppBar = function (props) {
	const classes = useStyles()

	const [menuOpen, setMenuOpen] = useState(false)
	const [menuAnchor, setMenuAnchor] = useState(null)
	const [aboutOpen, setAboutOpen] = useState(false)

	return (
		<div>
			<AppBar position="static" className={classes.navbar}>
				<Toolbar>
					<IconButton edge="start" className={classes.menuButton} color="inherit"
						aria-label="menu"
						ref={(r) => setMenuAnchor(r)}
						onClick={() => setMenuOpen(true)}>
						<MenuIcon />
					</IconButton>
					<Typography variant="h4" className={classes.title}>Marque</Typography>
					
					<TextField id="standard-basic" label="Search" className={classes.search}
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
					<Menu
						anchorEl={menuAnchor}
						open={menuOpen}
						onClose={() => setMenuOpen(false)}>
						<MenuItem onClick={() => { setAboutOpen(true); setMenuOpen(false) }}>
							About</MenuItem>
					</Menu>
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
					Marque {VERSION} Copyright © 2020 Tero Oinas.<br/>
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
