const CopyWebpackPlugin = require('copy-webpack-plugin');
const TerserJSPlugin = require('terser-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const MomentLocalesPlugin = require('moment-locales-webpack-plugin');
const ManifestPlugin = require('webpack-manifest-plugin');
const CleanObsoleteChunks = require('webpack-clean-obsolete-chunks');
const LiveReloadPlugin = require('webpack-livereload-plugin');
const IgnoreEmitPlugin = require('ignore-emit-webpack-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const path = require('path');
 
const APP_ENV = process.env.NODE_ENV || 'development';

module.exports = {
  mode: APP_ENV,
  // devtool: APP_ENV == 'production' ? false: 'source-map',
  entry: [
    './resources/scripts/index.js',
    './resources/styles/index.css',
  ],
  output: { filename: 'bundle.js', path: `${__dirname}/public/assets` },
  resolve: {
    alias: {
      'svelte': path.resolve('node_modules', 'svelte'),
      '~': path.resolve('resources', 'scripts'),
    },
    extensions: ['.scss', '.mjs', '.js', '.svelte'],
    mainFields: ['svelte', 'browser', 'module', 'main']
  },
  optimization: {
    minimizer: [new TerserJSPlugin({}), new OptimizeCSSAssetsPlugin({})],
  },
  module: {
    rules: [

      /*
       ██████╗███████╗███████╗
      ██╔════╝██╔════╝██╔════╝
      ██║     ███████╗███████╗
      ██║     ╚════██║╚════██║
      ╚██████╗███████║███████║
       ╚═════╝╚══════╝╚══════╝

      */

      {
        test: /\.css$/,
        // exclude: /node_modules/,
        use: [
          { loader: MiniCssExtractPlugin.loader },
          { loader: 'css-loader' },
          { loader: 'postcss-loader' },
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
        test: /\.(svelte|html)$/,
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

      /*
      ███████╗ ██████╗ ███╗   ██╗████████╗███████╗
      ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔════╝
      █████╗  ██║   ██║██╔██╗ ██║   ██║   ███████╗
      ██╔══╝  ██║   ██║██║╚██╗██║   ██║   ╚════██║
      ██║     ╚██████╔╝██║ ╚████║   ██║   ███████║
      ╚═╝      ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚══════╝

      */

      {
        test: /\.(woff2?|ttf|otf|eot|svg)$/,
        exclude: /node_modules/,
        loader: 'file-loader',
        options: {
            name: '[path][name].[ext]'
        }
      },

    ]
  },
  plugins: [
    new CleanObsoleteChunks(),
    new CopyWebpackPlugin(
      [{ from: './resources', to: '' }], { copyUnmodified: true, ignore: ['styles/**', 'scripts/**', 'views/**', 'webfonts/**'] }
    ),
    // new IgnoreEmitPlugin(/\.s[ac]ss$/),
    new MiniCssExtractPlugin({ filename: 'bundle.css' }),
    new MomentLocalesPlugin(),
    // new ManifestPlugin({ fileName: 'manifest.json' }),
    APP_ENV != 'production' && new LiveReloadPlugin({}),
    APP_ENV != 'production' && new BundleAnalyzerPlugin()
  ].filter(p => p !== false)
}