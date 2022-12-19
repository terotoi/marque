import React from 'react'
import BookmarkList from './bm_list'
import EditBookmarkDialog from './edit'
import TagList from './taglist'
import AlertDialog from './dialogs/alert'
import { fetchJSON } from './api'

const FILTER_TEXT_DELAY = 500

function convertDate(str) {
  if (str) {
    const d = new Date(str)
    return d
  }
  return null
}

/**
 *  Database is an interface to the bookmark database on the server.
 * 
 * @param {state} ctx - application context
 */
export default class DatabaseUI extends React.Component {
  constructor(props) {
    super(props)

    this.ctx = props.ctx

    this.state = {
      bookmarks: [],
      tags: [],
      filterText: '',
      filterTags: [],
      editBookmark: null,
      alert: null
    }

    this.onTagClicked = this.onTagClicked.bind(this)

    fetchJSON('/api/get_bookmarks', 'get', 'json', null, this.ctx.authToken, (bms) => {
      if (bms === null)
        bms = []
      bms.forEach(bm => bm.Updated = convertDate(bm.Updated))
      this.setState({ bookmarks: bms, tags: this.collectTags(bms) })
    }, (err) => {
      this.setState({
        alert: {
          title: "Server error",
          text: err.toString()
        }
      })
    })
  }

  get bookmarks() {
    return this.state.bookmarks
  }

  set filterText(text) {
    if (this.delay)
      clearTimeout(this.delay)

    this.delay = setTimeout(() => {
      this.delay = null
      this.setState({ filterText: text })
    }, FILTER_TEXT_DELAY)
  }

  addBookmark() {
    this.setState({ editBookmark: this.newBookmark() })
  }

  newBookmark() {
    return {
      ID: -1,
      UserID: null,
      URL: '',
      Title: '',
      Tags: [],
      Notes: '',
    }
  }

  // Collect and return all unique tags in the given bookmarks array.
  collectTags(bms) {
    let tags = new Map()
    for (var bm of bms) {
      if (bm.Tags) {
        for (var tag of bm.Tags) {
          tags.set(tag, true)
        }
      }
    }

    let keys = [...tags.keys()]
    keys.sort()
    return keys
  }

  // Called after changes to bookmarks or tags
  update() {
    const tags = this.collectTags(this.state.bookmarks)
    const filterTags = this.state.filterTags.filter((ft) =>
      tags.indexOf(ft) != -1)

    this.setState({
      bookmarks: this.state.bookmarks,
      tags: tags,
      filterTags: filterTags
    })
  }

  createBookmark(bm) {
    fetchJSON('/api/bookmark/create', 'post', 'json', bm, this.ctx.authToken, (result) => {
      if (result === false)
        this.setState({ alert: { title: "Server error", text: "Failed to add a bookmark." } })
      else if (result === true)
        this.setState({ alert: { title: "Duplicate bookmark found." } })
      else {
        result.Updated = convertDate(result.Updated)
        this.state.bookmarks.unshift(result)
        this.update()
      }
    }, (err) => {
      this.setState({
        alert: {
          title: "Server error",
          text: err.toString()
        }
      })
    })
  }

  updateBookmark(bm) {
    fetchJSON('/api/bookmark/update', 'post', 'json', bm, this.ctx.authToken, (result) => {
      if (result === false)
        this.setState({ alert: { title: "Server error" } })
      else if (result === true)
        this.setState({ alert: { title: "Duplicate bookmark found." } })
      else {
        // Update locally
        for (var b of this.state.bookmarks) {
          if (b.ID == result.ID) {
            b.URL = result.URL
            b.Title = result.Title
            b.Tags = result.Tags
            b.Notes = result.Notes
            b.Updated = convertDate(result.Updated)
            this.update()
            break
          }
        }
      }
    }, (err) => {
      this.setState({
        alert: {
          title: "Server error",
          text: err.toString()
        }
      })
    })
  }

  deleteBookmark(bm) {
    fetchJSON('/api/bookmark/delete/' + bm.ID, 'post', 'json', null, this.ctx.authToken, (result) => {
      if (result === false)
        this.setState({ alert: ["Server error", "Failed to delete a bookmark."] })
      else {
        this.state.bookmarks.splice(this.state.bookmarks.indexOf(bm), 1)
        this.update()
      }
    }, (err) => {
      this.setState({
        alert: {
          title: "Server error",
          text: err.toString()
        }
      })
    })
  }

  onTagClicked(tag) {
    const i = this.state.filterTags.indexOf(tag)
    let selected
    if (i == -1) {
      selected = [...this.state.filterTags, tag]
    } else {
      selected = this.state.filterTags.filter((t) => t !== tag)
    }
    this.setState({ filterTags: selected })
  }

  render() {
    return (
      <React.Fragment>
        {this.state.alert ?
          <AlertDialog
            title={this.state.alert.title}
            text={this.state.alert.text}
            onClose={() => this.setState({ alert: null })} /> : null}

        {(this.state.editBookmark !== null) ?
          <EditBookmarkDialog
            bookmark={this.state.editBookmark}
            onSubmit={(bm) => {
              bm.ID != -1 ?
                this.updateBookmark(bm) :
                this.createBookmark(bm)
            }}
            onClose={() => this.setState({ editBookmark: null })} /> : null}

        <TagList
          all={this.state.tags}
          selected={this.state.filterTags}
          onTagClicked={this.onTagClicked} />
        <BookmarkList
          bookmarks={this.state.bookmarks}
          filterText={this.state.filterText}
          filterTags={this.state.filterTags}
          onEdit={(bm) => this.setState({ editBookmark: bm })}
          onDelete={(bm) => this.deleteBookmark(bm)}
          onTagClicked={this.onTagClicked} />
      </React.Fragment>
    )
  }
}
