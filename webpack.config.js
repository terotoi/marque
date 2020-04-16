module.exports = {
	output: {
		path: __dirname + '/public/dist',
		filename: 'ui.js'
	},
	entry: './ui/main',
	performance: {
		hints: false
	},
	module: {
		rules: [
			{
				test: /\.(js|jsx)$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader'
				}
			}
		]
	}
}

