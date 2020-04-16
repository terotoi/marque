import React from 'react'
import {
	Button, Dialog, DialogActions, DialogContent, DialogContentText,
	DialogTitle
} from '@material-ui/core'

// props:
//    title: string   // Alert title text
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
