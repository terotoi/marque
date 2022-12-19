import React, { useState } from 'react'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogContentText from '@material-ui/core/DialogContentText'
import DialogTitle from '@material-ui/core/DialogTitle'

// EditBookmarkDialog is used for adding and editing bookmarks.
// props:
//   bookmark: Bookmark
//   onClose: function()
//   onSubmit: function(Bookmark)
export default function EditBookmarkDialog(props) {
  const [url, setURL] = useState(props.bookmark.URL)
  const [tags, setTags] = useState(props.bookmark.Tags ? props.bookmark.Tags.join(' ') : '')
  const [title, setTitle] = useState(props.bookmark.Title)
  const [notes, setNotes] = useState(props.bookmark.Notes)

  const onSubmit = () => {
    // Allow both comma and space as a separator.
    const t = tags.split(' ').map((t) => t.split(',')).flat().filter((x) => x != '')

    if (url != "") {
      const bm = {
        ID: props.bookmark.ID,
        URL: url,
        Tags: t,
        Title: title,
        Notes: notes
      }
      props.onClose()
      props.onSubmit(bm)
    }
  }

  return (
    <Dialog open={true} onClose={props.onClose} aria-labelledby="form-dialog-title">
      <DialogTitle id="form-dialog-title">{props.bookmark.ID == -1 ? "Add" : "Edit"} a bookmark</DialogTitle>
      <DialogContent>
        <DialogContentText>Enter data for an URL:</DialogContentText>

        <TextField autoFocus margin="dense" id="url"
          label="URL" type="url" fullWidth required
          value={url} onChange={(ev) => setURL(ev.target.value)} />

        <TextField margin="dense" id="tags"
          label="Tags" type="text" fullWidth
          value={tags} onChange={(ev) => setTags(ev.target.value)} />
        <TextField margin="dense" id="title"
          label="Title (optional)" type="text" fullWidth
          value={title} onChange={(ev) => setTitle(ev.target.value)} />
        <TextField margin="dense" id="notes"
          label="Notes (optional)" type="text" fullWidth
          value={notes} onChange={(ev) => setNotes(ev.target.value)} />

      </DialogContent>
      <DialogActions>
        <Button onClick={props.onClose} color="primary">
          Cancel
        </Button>
        <Button onClick={onSubmit} color="primary">
          Submit
        </Button>
      </DialogActions>
    </Dialog>)
}

