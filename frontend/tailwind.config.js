/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,svelte}", "./index.html"],
    theme: {
      extend: {},
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        light: '#EEEEEE',
        dark: '#000000',
        lightDark: '#282A3A',
        accent: '#C69749',
        accentDark: '#735F32',
        green: '#13CE66',
        red: '#FF4949',
        yellow: '#FFD166',
      }
    },
    plugins: [],
  }