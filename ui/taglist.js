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
			flexGrow: 0,
			flexShrink: 0,
			width: '8%',
			minWidth: '8em',
			border: '1px solid #777',
			overflow: 'hidden'
		}
	}
	return styles
}

const useStyles = makeStyles(styles)

/**
 * TagList
 * 
 * @param {[]string} all - all tags
 * @param {[]string} selected - currently selected tags
 * @param {function} onTagClicked -
 */
const TagList = function (props) {
	const classes = useStyles()

	const t_html = props.all.map((tag) => {
		return (
			<Button
				className={classes.tag}
				variant="outlined"
				size="small"
				color={props.selected.indexOf(tag) != -1 ? "secondary" : "inherit"}
				key={tag}
				onClick={(ev) => props.onTagClicked(tag)}>
				{tag}</Button>)
	})

	return (
		<Box className={classes.taglist} display="flex" flexDirection="row" flexWrap="wrap">
			{t_html}
		</Box>)
}

export default TagList
