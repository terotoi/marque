import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogContentText from '@material-ui/core/DialogContentText'
import TextField from '@material-ui/core/TextField'

const styles = makeStyles((theme) => ({
	username: {
		display: 'none'
	},
	newPasswd: {
		marginTop: '2em'
	}
}))

/**
 * PasswordDialog is a dialog for changing a user's password.
 * 
 * @param {string} props.text - dialog text
 * @param {function} props.onClose - called when the dialog is closed for any reason
 * @param {function(oldPasswd, newPasswd)} props.onConfirm - function called
 *  when the dialog is confirmed
 */
export default function PasswordDialog(props) {
	const [oldPasswd, setOldPasswd] = React.useState("")
	const [newPasswd, setNewPasswd] = React.useState("")
	const classes = styles()

	const handleConfirm = () => {
		props.onClose()
		if (props.onConfirm)
			props.onConfirm(oldPasswd, newPasswd)
	}

	return (
		<Dialog
			open={true}
			onClose={props.onClose}
			aria-labelledby="passwd-dialog-title"
			aria-describedby="passwd-dialog-description">
			<DialogContent>
				<DialogContentText id="passwd-dialog-description">
					{props.text}
				</DialogContentText>

				<input type="text" name="username" value=""
					className={classes.username}
					onChange={() => {}}/>

				<TextField
					color="primary"
					autoFocus required fullWidth
					margin="dense"
					name="current-password"
					label="Enter old password"
					type="password"
					autoComplete="current-password"
					value={oldPasswd}
					onChange={(ev) => { setOldPasswd(ev.target.value) }} />

				<TextField
					className={classes.newPasswd}
					color="primary"
					required fullWidth
					margin="dense"
					name="new-password"
					label="Enter new password"
					type="password"
					autoComplete="new-password"
					required={true}
					value={newPasswd}
					onChange={(ev) => { setNewPasswd(ev.target.value) }} />
			</DialogContent>
			<DialogActions>
				<Button onClick={props.onClose}>
					Cancel
				</Button>
				<Button onClick={handleConfirm} disabled={newPasswd == ""}>
					Confirm
				</Button>
			</DialogActions>
		</Dialog>)
}

/**
 * Creates a PasswordDialog and adds it to the context using context.addDialog(dialog)
 * On close, context.removeDialog(dialog) will be called.
 * 
 * @param {string} props.text - main text of the dialog
 * @param {function(oldPasswd, newPasswd)} props.onConfirm - function called
*  when the dialog is confirmed
 * @param {state} ctx - application context
*/
export function openPasswordDialog(ctx, props) {
	const dialog =
		<PasswordDialog
			text={props.text}
			onConfirm={props.onConfirm}
			onClose={() => { ctx.removeDialog(dialog) }} />
	ctx.addDialog(dialog)
}

