"use strict"
import { makeStyles, Box, Button } from '@material-ui/core'
import React, { useState } from 'react'

const styles = (theme) => {
	const styles = {
		taglist: {
			marginBottom: theme.spacing(2)
		},
		tag: {
			marginRight: theme.spacing(2),
			marginBottom: theme.spacing(1),
			flexGrow: 1,
			flexShrink: 1,
			flexBasis: '10em',
			width: 0,
			textAlign: 'center',
			border: '1px solid #777'
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

	console.log("TagList__")

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
		<Box className={classes.taglist} display="flex" flexDirection="row" flexWrap="wrap">
			{t_html}
		</Box>)
}

export default TagList
