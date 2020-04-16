"use strict"
import { makeStyles, Button } from '@material-ui/core'
import React, { useState } from 'react'

const styles = (theme) => {
	const styles = {
		taglist: {
			marginBottom: theme.spacing(2)
		},
		tag: {
			marginRight: theme.spacing(2),
			marginBottom: theme.spacing(1)
		}
	}
	return styles
}

const useStyles = makeStyles(styles)

// props:
// 		all: []string       All tags
//    selected: []string  Selected tags
//    onChange(selected: []string) Called when list of selected tags changes.
const TagList = function (props) {
	const classes = useStyles()

	const onTagClicked = function (tag) {
		const i = props.selected.indexOf(tag)
		let s
		if (i == -1) {
			s = [...props.selected, tag]
		} else {
			s = props.selected.filter((t) => t !== tag)
		}
		props.onChange(s)
	}

	const t_html = props.all.map((tag) => {
		return (<Button
			className={classes.tag}
			variant="outlined"
			size="small"
			color={props.selected.indexOf(tag) != -1 ? "primary" : "default"}
			key={tag}
			onClick={(ev) => onTagClicked(tag)}>
			{tag}</Button>)
	})

	return (
		<div className={classes.taglist}>
			{t_html}
		</div>)
}

export default TagList
