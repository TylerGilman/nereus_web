/** @type {import('tailwindcss').Config} */
const { colors: defaultColors } = require('tailwindcss/defaultTheme')

module.exports = {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
  safelist: [],
  theme: defaultColors,
}

