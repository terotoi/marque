import React from 'react'
import App from './app'

export default class Context extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			authToken: localStorage.getItem('authToken') || null,

			setAuthToken: (token) => {
				if (token === null)
					localStorage.removeItem('authToken')
				else
					localStorage.setItem('authToken', token)
				this.setState({ authToken: token })
			},

			username: localStorage.getItem('username') || '',
			setUsername: (name) => {
				localStorage.setItem('username', name)
				this.setState({ username: name })
			},

			isAdmin: localStorage.getItem('isAdmin') == 'true',
			setIsAdmin: (admin) => {
				localStorage.setItem('isAdmin', admin)
				this.setState({ isAdmin: admin })
			},

			logout: () => {
				this.state.setAuthToken(null)
				this.state.setUsername('')
				this.state.setIsAdmin(false)
			},

			dialogs: [],
			addDialog: (dialog) => {
				this.setState({ dialogs: [...this.state.dialogs, dialog] })
			},

			removeDialog: (dialog) => {
				this.setState({ dialogs: this.state.dialogs.filter((x) => x !== dialog) })
			}
		}
	}

	render() {
		return <App ctx={this.state} />
	}
}
