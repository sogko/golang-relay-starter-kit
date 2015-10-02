import express from 'express';
import path from 'path';
import webpack from 'webpack';
import WebpackDevServer from 'webpack-dev-server';

const APP_PORT = 3000;
const GRAPHQL_PORT = 8080;

// Serve the Relay app
var compiler = webpack({
  entry: path.resolve(__dirname, 'js', 'app.js'),
  module: {
    loaders: [
      {
        exclude: /node_modules/,
        loader: 'babel',
        query: {stage: 0, plugins: ['./build/babelRelayPlugin']},
        test: /\.js$/
      }
    ]
  },
  output: {filename: 'app.js', path: '/'}
});
var app = new WebpackDevServer(compiler, {
  contentBase: '/public/',
  proxy: {'/graphql': `http://localhost:${GRAPHQL_PORT}`},
  publicPath: '/js/',
  stats: {colors: true}
});
// Serve static resources
app.use('/', express.static(path.resolve(__dirname, 'public')));
app.listen(APP_PORT, () => {
  console.log(`App is now running on http://localhost:${APP_PORT}`);
});
