/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors');

module.exports = {
  content: ['./src/**/*.{js,ts,jsx,tsx,mdx}'],
  mode: 'jit',
  purge: ['./src/**/*.{js,jsx,ts,tsx,mdx}'],
  theme: {
    screens: {
      // sm: '480px',
      // md: '768px',
      lg: '976px',
      xl: '1440px',
    },
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      primary: colors.zinc,
      black: colors.zinc,
      white: colors.white,
      gray: colors.gray,
      red: colors.red,
      // orange: colors.orange,
      // amber: colors.amber,
      // yellow: colors.yellow,
      // lime: colors.lime,
      // green: colors.green,
      // emerald: colors.emerald,
      // teal: colors.teal,
      // cyan: colors.cyan,
      sky: colors.sky,
      // blue: colors.blue,
      // indigo: colors.indigo,
      // violet: colors.violet,
      // purple: colors.purple,
      // fuchsia: colors.fuchsia,
      // pink: colors.pink,
      // rose: colors.rose,
    },
    extend: {},
  },
  plugins: [],
};
