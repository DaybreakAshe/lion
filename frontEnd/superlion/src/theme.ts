import { createTheme } from '@mui/material/styles'

// import 'slick-carousel/slick/slick.css'
// import 'slick-carousel/slick/slick-theme.css'

// define theme
export const theme = createTheme({
  spacing: 4,
  palette: {
    primary: { 
      main: '#307DCF',
      light: '#4B48FD'
    },
    secondary: {
      main: '#FB6B5A'
    },
    background: {
      default: '#F9F9F9',
      paper: '#ffff'
    },
  },
  breakpoints: {
    values: {
      xs: 0,
      sm: 600,
      md: 960,
      lg: 1280,
      xl: 1920,
    },
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          scrollbarColor: '#F9F9F9',
          '&::-webkit-scrollbar, & *::-webkit-scrollbar': {
            backgroundColor: '#F9F9F9'
          },
          '&::-webkit-scrollbar-thumb, & *::-webkit-scrollbar-thumb': {
            borderRadius: 8,
            backgroundColor: '#6b6b6b',
            backgroundClip: 'content-box',
            minHeight: 24,
            border: '4px solid transparent'
          },
          '&::-webkit-scrollbar-thumb:focus, & *::-webkit-scrollbar-thumb:focus': {
            backgroundColor: '#959595'
          },
          '&::-webkit-scrollbar-thumb:active, & *::-webkit-scrollbar-thumb:active': {
            backgroundColor: '#959595'
          },
          '&::-webkit-scrollbar-thumb:hover, & *::-webkit-scrollbar-thumb:hover': {
            backgroundColor: '#959595'
          },
          '&::-webkit-scrollbar-corner, & *::-webkit-scrollbar-corner': {
            backgroundColor: '#2b2b2b'
          },
          a: {
            textDecorator: 'none'
          }
        }
      }
    }
  }
})
