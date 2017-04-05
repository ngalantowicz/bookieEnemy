const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const path = require('path');

module.exports = {

  entry: './client/src/index.jsx',
  output: {
    path: path.resolve(__dirname, 'client/src'),
    filename: 'app.min.js'
  },
  devServer: {
    contentBase: path.join(__dirname, 'client/src'),
    compress: true,
    stats: 'errors-only',
    open: true,
    historyApiFallback: true
  },

  module: {
    rules: [
      {
        test: /\.css$/,
        use: ExtractTextPlugin.extract({fallback: 'style-loader', use: ['css-loader'], publicPath: '/dist'})
      }, {
        test: /\.jsx$/,
        exclude: /(node_modules|bower_components)/,
        use: 'babel-loader'
      }
    ]
  },

  plugins: [
    new HtmlWebpackPlugin({
      title: 'Bookie Enemy',
      hash: true,
      template: './client/src/index.html'
    }),
    new ExtractTextPlugin({filename: 'style.css', disable: false, allChunks: true})
  ]
};
