const CopyWebpackPlugin = require('copy-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const ManifestPlugin = require('webpack-manifest-plugin');
const CleanObsoleteChunks = require('webpack-clean-obsolete-chunks');
const LiveReloadPlugin = require('webpack-livereload-plugin');
const IgnoreEmitPlugin = require('ignore-emit-webpack-plugin');
const path = require('path');
 
module.exports = {
  mode: process.env.NODE_ENV || 'development',
  entry: [
    './assets/scripts/index.js',
    './assets/styles/index.styl',
  ],
  output: { filename: 'bundle.[hash].js', path: `${__dirname}/public/assets` },
  resolve: {
    alias: {
      'svelte': path.resolve('node_modules', 'svelte'),
      '~': path.resolve('assets', 'scripts'),
    },
    extensions: ['.styl', '.js', '.svelte'],
    mainFields: ['svelte', 'browser', 'module', 'main']
  },
  module: {
    rules: [

      /*
      ███████╗████████╗██╗   ██╗██╗     ██╗   ██╗███████╗
      ██╔════╝╚══██╔══╝╚██╗ ██╔╝██║     ██║   ██║██╔════╝
      ███████╗   ██║    ╚████╔╝ ██║     ██║   ██║███████╗
      ╚════██║   ██║     ╚██╔╝  ██║     ██║   ██║╚════██║
      ███████║   ██║      ██║   ███████╗╚██████╔╝███████║
      ╚══════╝   ╚═╝      ╚═╝   ╚══════╝ ╚═════╝ ╚══════╝

      */

      {
        test: /\.styl$/,
        exclude: /node_modules/,
        use: [
          { loader: MiniCssExtractPlugin.loader },
          { loader: 'css-loader' },
          {
            loader: 'stylus-loader',
            options: {
              paths: 'node_modules/bootstrap-stylus/stylus/'
            }
          }
        ]
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
        test: /\.(html|svelte)$/,
        exclude: /node_modules/,
        use: ['svelte-loader']
      },

    ]
  },
  plugins: [
    new CleanObsoleteChunks(),
    new CopyWebpackPlugin(
      [{ from: './assets', to: '' }], { copyUnmodified: true, ignore: ['styles/**', 'scripts/**'] }
    ),
    new IgnoreEmitPlugin(/\.styl$/),
    new MiniCssExtractPlugin({ filename: 'bundle.[contenthash].css' }),
    new ManifestPlugin({ fileName: 'manifest.json' }),
    new LiveReloadPlugin({}),
  ]
}