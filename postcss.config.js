module.exports = {
  ident: 'postcss',
  plugins: [
    require('postcss-import'),
    require('tailwindcss'),
    require('postcss-nested'),
    require('postcss-preset-env')({ stage: 1 }),
    require('autoprefixer'),
  ],
};