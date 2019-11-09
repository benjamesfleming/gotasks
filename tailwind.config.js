const _t = require('tailwindcss/defaultTheme')

module.exports = {
  theme: {
    extend: {
      colors: {
        primary: _t.colors.gray[500],
      }
    }
  },
  variants: {},
  plugins: [
    require('tailwindcss-transitions')(),
  ],
}
