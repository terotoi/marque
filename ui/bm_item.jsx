import React from 'react'
import {
	Divider, IconButton, ListItem, ListItemSecondaryAction, ListItemText
} from '@material-ui/core'
import { Delete, Edit } from '@material-ui/icons'

/**
 * BookmarkItem
 * 
 * @param {Object} item - the bookmark to show
 * @param {function} onDelete
 * @param {function} onDelete
 * @param {function} onTagClicked
 */
export default function BookmarkItem(props) {
	const item = props.item

	let h_tags
	if (item.Tags) {
		h_tags = item.Tags.map((tag, i) => {
			return <span key={i} onClick={(ev) => { ev.preventDefault(); props.onTagClicked(tag) }}>{tag} </span>
		})
	} else {
		h_tags = <span>No tags</span>
	}

	return (
		<React.Fragment>
			<ListItem alignItems="flex-start" button component="a" href={item.URL}>
				<ListItemText
					primary={item.Title}
					secondary={
						<React.Fragment>
							{h_tags}<br />
							{item.Notes ?
								<span onClick={(ev) => this.onDescClicked(ev, item)}>{item.Notes}<br /></span>
								: null}
							{item.Updated ? item.Updated.toLocaleDateString() + ' ' + item.Updated.toLocaleTimeString() : null}
						</React.Fragment>
					} />

				<ListItemSecondaryAction>
					<IconButton color='primary' onClick={() => props.onEdit(item)}>
						<Edit />
					</IconButton>
					<IconButton color='primary' onClick={() => props.onDelete(item)}>
						<Delete />
					</IconButton>
				</ListItemSecondaryAction>
			</ListItem>
			<Divider variant="inset" component="li" />
		</React.Fragment>)
}
