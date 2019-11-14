const _t = require('tailwindcss/defaultTheme')

module.exports = {
  theme: {
    extend: {
      colors: {
        primary: _t.colors.gray[300],
      },
      opacity: {
        '80': '0.8',
        '90': '0.9'
      },
      maxWidth: {
        ..._t.spacing,
      },
      maxHeight: {
        ..._t.spacing,
      }
    }
  },
  variants: {
    backgroundColor: ['responsive', 'odd', 'hover', 'focus']
  },
  plugins: [
    require('tailwindcss-transitions')(),
    require('@tailwindcss/custom-forms'),
  ],
}
