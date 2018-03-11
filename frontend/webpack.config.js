var webpack = require('webpack');
var path = require('path');

var BUILD_DIR = path.resolve(__dirname, 'public');
var APP_DIR = path.resolve(__dirname, 'src');

var config = {
    entry: APP_DIR + '/index.js',
    output: {
        path: BUILD_DIR,
        filename: 'bundle.js'
    },
    module : {
        loaders : [
            {
                test : /\.jsx?/,
                include : APP_DIR,
                loader : 'babel',
                query: {
                  presets: ['react', 'es2015']
                }
            }
        ]
    },
    plugins: [
      new webpack.EnvironmentPlugin({
        NODE_ENV: 'development', // use 'development' unless process.env.NODE_ENV is defined
        DEBUG: false
      })
    ]
};

module.exports = config;
