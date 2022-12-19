import React from 'react'
import {
	Button, Dialog, DialogActions, DialogContent, DialogContentText,
	DialogTitle
} from '@material-ui/core'

// props:
//    title: string  // Alert title text
//    text: string   // Alert text
//    onClose: function()
export default function AlertDialog(props) {
	const onClose = () => {
		props.onClose()
	}

	return (
		<Dialog
			open={true}
			onClose={onClose}
			aria-labelledby="alert-dialog-title"
			aria-describedby="alert-dialog-description">
			<DialogTitle id="alert-dialog-title">{props.title}</DialogTitle>
			<DialogContent>
				<DialogContentText id="alert-dialog-description">
					{props.text || ''}
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


/**
 * Creates an AlertDialog and adds it to the context using ctx.addDialog(dialog)
 * On close, ctx.removeDialog(dialog) will be called.
 * 
 * @param {string} props.title - title of the dialog (optional)
 * @param {string} props.text - main text of the dialog
 * @param {state} ctx - application context
 */
export function openAlertDialog(ctx, props) {
	const dialog =
		<AlertDialog
		  title={props.title || props.text}
			text={props.text}
			onConfirm={props.onConfirm}
			onClose={() => { ctx.removeDialog(dialog) }} />
	ctx.addDialog(dialog)
}

