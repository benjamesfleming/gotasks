const CopyWebpackPlugin = require('copy-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const ManifestPlugin = require('webpack-manifest-plugin');
const CleanObsoleteChunks = require('webpack-clean-obsolete-chunks');
const LiveReloadPlugin = require('webpack-livereload-plugin');
const IgnoreEmitPlugin = require('ignore-emit-webpack-plugin');
const path = require('path');
 
const APP_ENV = process.env.NODE_ENV || 'development';

module.exports = {
  mode: APP_ENV,
  // devtool: APP_ENV == 'producation' ? false: 'source-map',
  entry: [
    './resources/scripts/index.js',
    './resources/styles/index.scss',
  ],
  output: { filename: '[name].[hash].js', path: `${__dirname}/public/assets` },
  resolve: {
    alias: {
      'svelte': path.resolve('node_modules', 'svelte'),
      '~': path.resolve('resources', 'scripts'),
    },
    extensions: ['.scss', '.mjs', '.js', '.svelte'],
    mainFields: ['svelte', 'browser', 'module', 'main']
  },
  module: {
    rules: [

      /*
      ███████╗ █████╗ ███████╗███████╗
      ██╔════╝██╔══██╗██╔════╝██╔════╝
      ███████╗███████║███████╗███████╗
      ╚════██║██╔══██║╚════██║╚════██║
      ███████║██║  ██║███████║███████║
      ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝

      */

      {
        test: /\.s[ac]ss$/,
        exclude: /node_modules/,
        use: [
          { loader: MiniCssExtractPlugin.loader },
          { loader: 'css-loader' },
          { loader: 'sass-loader' },
        ]
      },

      {
        test: /\.css$/,
        exclude: /node_modules/,
        use: [
          { loader: MiniCssExtractPlugin.loader },
          { loader: 'css-loader' },
        ]
      },

      /*
      ███████╗ ██████╗██████╗ ██╗██████╗ ████████╗███████╗
      ██╔════╝██╔════╝██╔══██╗██║██╔══██╗╚══██╔══╝██╔════╝
      ███████╗██║     ██████╔╝██║██████╔╝   ██║   ███████╗
      ╚════██║██║     ██╔══██╗██║██╔═══╝    ██║   ╚════██║
      ███████║╚██████╗██║  ██║██║██║        ██║   ███████║
      ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝╚═╝        ╚═╝   ╚══════╝

      */

      {
        test: /\.m?js$/,
        exclude: /node_modules/,
        use: [
          { loader: 'babel-loader' },
        ],
      },

      /*
      ███████╗██╗   ██╗███████╗██╗  ████████╗███████╗
      ██╔════╝██║   ██║██╔════╝██║  ╚══██╔══╝██╔════╝
      ███████╗██║   ██║█████╗  ██║     ██║   █████╗
      ╚════██║╚██╗ ██╔╝██╔══╝  ██║     ██║   ██╔══╝
      ███████║ ╚████╔╝ ███████╗███████╗██║   ███████╗
      ╚══════╝  ╚═══╝  ╚══════╝╚══════╝╚═╝   ╚══════╝

      */

      {
        test: /\.svelte$/,
        // exclude: /node_modules/,
        use: [
          { loader: 'babel-loader' },
          { 
            loader: 'svelte-loader',
            options: {
              ...require('./svelte.config.js'),
              emitCss: true,
            }
          },
        ]
      },

    ]
  },
  plugins: [
    new CleanObsoleteChunks(),
    // new CopyWebpackPlugin(
    //   [{ from: './resources/static', to: '' }], { copyUnmodified: true, ignore: ['styles/**', 'scripts/**'] }
    // ),
    // new IgnoreEmitPlugin(/\.s[ac]ss$/),
    new MiniCssExtractPlugin({ filename: '[name].[contenthash].css' }),
    new ManifestPlugin({ fileName: 'manifest.json' }),
    new LiveReloadPlugin({}),
  ]
}