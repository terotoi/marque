import {
	Divider, IconButton, List, ListItem, ListItemSecondaryAction, ListItemText,
	Typography, withStyles
} from '@material-ui/core'
import { Delete, Edit, ContactSupportOutlined, ContactsOutlined } from '@material-ui/icons'
import React from 'react'

const styles = {
}

// BookmarkList shows the list of bookmarks.
// props:
//   filterText: string
//   filterTags: []string
//   onEdit: function(Bookmark)
export default withStyles(styles)(class BookmarkList extends React.Component {
	constructor(props) {
		super(props)

		this.onBookmarkClicked = this.onBookmarkClicked.bind(this)
		this.onSelectTag = this.onSelectTag.bind(this)

		this.delay = null
	}

	newestBookmarks(bms, num) {
		bms.sort((a, b) => {
			return b.Updated - a.Updated
		})
		return bms.slice(0, num)
	}

	filterBookmarks(bms, text, tags) {
		text = text.trim()
		if(!text && tags.length == 0) {
			return this.newestBookmarks(bms, 10)
		}

		let r = []
		for (var bm of bms) {
			let match = true
			for (var tag of tags) {
				if (bm.Tags == null || bm.Tags.indexOf(tag) == -1) {
					match = false
					break
				}
			}

			if (match)
				r.push(bm)
		}

		if (text) {
			// Special case for all bookmarks
			text = text.toLocaleLowerCase()

			// Split filter into parts at commas.
			var kws = text.split(' ')
			kws = kws.map((f) => f.trim())

			r = r.filter((bm) => {
				const t = (bm.Title + ' ' + (bm.Tags ? bm.Tags.join(' ') : '') +
					bm.Notes).toLocaleLowerCase()
				for (var kw of kws) {
					if (t.indexOf(kw) == -1)
						return false
				}
				return true
			})
		}

		r.sort((a, b) => {
			return b.Updated - a.Updated
		})
		return r
	}

	onBookmarkClicked(bm) {
		window.open(bm.URL, '_blank')
	}

	onDescClicked(ev, bm) {
		ev.stopPropagation()
		this.props.onEdit(bm)
	}

	onSelectTag(tag, ev) {
		ev.stopPropagation()
	}

	// TODO: move to an own component
	renderItem(item, i) {
		const classes = this.props.classes

		let h_tags
		if (item.Tags) {
			h_tags = item.Tags.map((tag, i) => {
				return <span key={i} onClick={(ev) => this.onSelectTag(tag, ev)}>{tag} </span>
			})
		} else {
			h_tags = <span>No tags</span>
		}

		return (
			<React.Fragment key={i}>
				<ListItem alignItems="flex-start">
					<ListItemText
						onClick={() => this.onBookmarkClicked(item)}
						primary={item.Title}
						secondary={
							<React.Fragment>
								{h_tags}<br/>
								{item.Notes ?
									<span onClick={(ev) => this.onDescClicked(ev, item)}>{item.Notes}<br/></span>
									: null}
								{item.Updated ? item.Updated.toLocaleDateString() + ' ' + item.Updated.toLocaleTimeString() : null}
							</React.Fragment>
						} />

					<ListItemSecondaryAction>
						<IconButton color='primary' onClick={() => this.props.onEdit(item)}>
							<Edit />
						</IconButton>
						<IconButton color='primary' onClick={() => this.props.onDeleteClicked(item)}>
							<Delete />
						</IconButton>
					</ListItemSecondaryAction>
				</ListItem>
				<Divider variant="inset" component="li" />
			</React.Fragment>)
	}

	render() {
		const filtered = this.filterBookmarks(this.props.bookmarks,
			this.props.filterText, this.props.filterTags)

		return (
			<React.Fragment>
				<Typography variant="body1">{filtered.length} / {this.props.bookmarks.length} bookmarks displayed.</Typography>
				<List component="nav">
					{filtered.map((bm, i) => {
						return this.renderItem(bm, i)
					})}
				</List>
			</React.Fragment>
		)
	}
})