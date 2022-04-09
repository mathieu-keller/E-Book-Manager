const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = env => {
  var production = env && env.NODE_ENV === 'production';
  return {
    entry: ['./src/index.tsx'],
    mode: production ? 'production' : 'development',
    performance: {
      hints: production ? 'warning' : false,
    },
    devtool: production ? false : 'inline-source-map',
    module: {
      rules: [
        {
          test: /\.(ts|tsx)$/,
          exclude: [/\.test.(ts|tsx)$/, /node_modules/],
          use: {
            loader: 'ts-loader',
            options: {
              transpileOnly: true,
            },
          },
        },

        {
          test: /\.css$/,
          exclude: [/node_modules/],
          use: ['style-loader', 'css-loader', 'postcss-loader'],
        },
        {
          test: /\.(eot|svg|ttf|woff|woff2|png|jpg|gif)$/,
          type: 'asset',
        },
      ],
    },
    resolve: {
      extensions: ['.tsx', '.ts', '.js', '.css', '.scss'],
    },
    output: {
      filename: 'bundle.js',
      path: path.resolve(__dirname, 'dist/bundles'),
    },
    plugins: [
      new HtmlWebpackPlugin(
        Object.assign(
          {},
          {
            inject: true,
            template: './public/index.html',
            favicon: './public/favicon.png',
          }, production ? {
            minify: {
              removeComments: true,
              collapseWhitespace: true,
              removeRedundantAttributes: true,
              useShortDoctype: true,
              removeEmptyAttributes: true,
              removeStyleLinkTypeAttributes: true,
              keepClosingSlash: true,
              minifyJS: true,
              minifyCSS: true,
              minifyURLs: true,
            },
          } : undefined,
        ),
      )],
    devServer: !production ? {
      port: 3001,
      watchFiles: {paths: ['./src', './public']},
      hot: true,
      proxy: {
        '**': {
          changeOrigin: true,
          target: 'http://localhost:8080',
        },
      },
    } : {},
  };
};
