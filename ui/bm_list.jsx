import React from 'react'
import {
	List, Typography, withStyles
} from '@material-ui/core'
import BookmarkItem from './bm_item'

const styles = {
}

/**
 * BookmarkList shows the list of bookmarks.
 *
 * @param {string} filterText
 * @param {[]string} filterTags
 * @param {function} onEdit
 * @param {function} onDelete
 * @param {function} onTagClicked
 */
export default withStyles(styles)(class BookmarkList extends React.Component {
	constructor(props) {
		super(props)

		this.onBookmarkClicked = this.onBookmarkClicked.bind(this)
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
		if (!text && tags.length == 0) {
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

	render() {
		const filtered = this.filterBookmarks(this.props.bookmarks,
			this.props.filterText, this.props.filterTags)

		return (
			<React.Fragment>
				<Typography variant="body1">{filtered.length} / {this.props.bookmarks.length} bookmarks displayed.</Typography>
				<List component="nav">
					{filtered.map((bm, i) =>
						<BookmarkItem item={bm} key={i}
							onEdit={this.props.onEdit}
							onDelete={this.props.onDelete}
							onTagClicked={this.props.onTagClicked} />
					)}
				</List>
			</React.Fragment>
		)
	}
})