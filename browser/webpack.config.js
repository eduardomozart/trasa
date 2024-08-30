const path = require('path');
const webpack = require('webpack');
// const Crx3 = require('@snadeau/crx3-webpack-plugin');
const pkg = require('./package.json');

module.exports = {
  entry: {
    // Each entry in here would declare a file that needs to be transpiled
    // and included in the extension source.
    // For example, you could add a background script like:
    // background: './src/background.js',
    trasa: './src/trasa.js',
  },
  output: {
    // This copies each source entry into the extension dist folder named
    // after its entry config key.
    path: path.join(path.resolve(__dirname), 'extension', 'dist'),
    filename: '[name].js',
  },
  module: {
    // This transpiles all code (except for third party modules) using Babel.
    rules: [
      {
        exclude: /node_modules/,
        test: /\.js$/,
        // Babel options are in .babelrc
        use: ['babel-loader'],
      },
      // {
      //   exclude: /node_modules/,
      //   test: /\.(png|jpe?g|gif)$/i,
      //   use: ['file-loader'],
      // },
      {
        test: /\.(jpg|png)$/,
        use: {
          loader: 'url-loader',
        },
      },
    ],
  },
  resolve: {
    // This allows you to import modules just like you would in a NodeJS app.
    extensions: ['.js', '.jsx'],
    modules: ['src', 'node_modules'],
  },
  plugins: [
    // Since some NodeJS modules expect to be running in Node, it is helpful
    // to set this environment var to avoid reference errors.
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('production'),
    }),
    /* new Crx3({
      updateUrl: 'http://localhost:8000/',
      updateFilename: 'updates.xml',
      keyFile: './build.pem',
      contentPath: './extension',
      outputPath: './dist',
      name: `${pkg.name}`,
      zip: true,
      browserVersion: '64.0.3242'
    }) */
  ],
  // This will expose source map files so that errors will point to your
  // original source files instead of the transpiled files.
  devtool: 'sourcemap',
};
