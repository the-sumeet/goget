/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,svelte}"],
    theme: {
      extend: {},
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        light: '#EEEEEE',
        dark: '#222831',
        lightDark: '#393E46',
        accent: '#00ADB5',
        accentDark: '#006f75',
        green: '#13CE66',
        red: '#FF4949',
        yellow: '#FFD166',
      }
    },
    plugins: [],
  }