// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { createTheme } from '@mui/material/styles';
import { Rubik } from './fonts/Rubik';

export const Theme = createTheme({
  typography: {
    fontFamily: '"Rubik", "Open sans", "Arial", sans- serif',
    h1: {
      fontFamily: 'Rubik-Bold',
      fontSize: '1.625rem', //26px
      fontWeight: 500,
    },
    h2: {
      fontFamily: 'Rubik-Medium',
      fontSize: '1.5rem', //24px
      fontWeight: 500,
    },
    h3: {
      fontSize: '1.375rem', //22px
      fontWeight: 500,
    },
    h4: {
      fontSize: '1.25rem', //20px
      fontWeight: 500,
    },
    h5: {
      fontFamily: 'Rubik-Light',
      fontSize: '1.25rem', //20px
      fontWeight: 300,
    },
    h6: {
      fontSize: '1rem', //16px
      fontWeight: 300,
      textTransform: 'uppercase',
    },
    subtitle1: {
      fontSize: '1.25rem', //20px
      fontWeight: 500,
    },
    subtitle2: {
      fontSize: '1.125rem', //18px
      fontWeight: 500,
    },
    body1: {
      fontFamily: 'Rubik',
      fontSize: '1rem', //16px
    },
    body2: {
      fontSize: '0.875rem', //14px
    },
    button: {
      fontFamily: 'Rubik-Medium',
      fontSize: '14px',
    }
  },
  palette: {
    primary: {
      main: '#8C9AF8',
    },
    text: {
      primary: '#fafafa',
    },
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        '@font-face': Rubik.Regular,
        fallbacks: [
          { '@font-face': Rubik.Bold },
          { '@font-face': Rubik.Medium },
          { '@font-face': Rubik.Light },
        ],
        body: {
          backgroundColor: '#0f1014',
        },
      }
    },
  },
});
