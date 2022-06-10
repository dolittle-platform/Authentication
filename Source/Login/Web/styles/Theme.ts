// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { createTheme } from '@mui/material/styles';

import { Rubik } from './fonts/Rubik';
import Symbol from './images/symbol.svg?url';

export const Theme = createTheme({
  typography: {
    fontFamily: 'Rubik',
    h1: {
      fontFamily: 'Rubik-Bold',
      fontSize: '30px',
      letterSpacing: '0.3px',
    },
    h2: {
      fontFamily: 'Rubik-Medium',
      fontSize: '22px',
      letterSpacing: '-0.22px',
    },
    body1: {
      fontFamily: 'Rubik',
      fontSize: '18px',
      letterSpacing: '0.18px',
    },
    button: {
      fontFamily: 'Rubik-Medium',
      fontSize: '14px',
      letterSpacing: '0.84px',
    }
  },
  palette: {
    primary: {
      main: '#FAFAFA',
    },
    text: {
      primary: '#FAFAFA'
    }
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        '@font-face': Rubik.Regular,
        fallbacks: [
          { '@font-face': Rubik.Bold },
          { '@font-face': Rubik.Medium },
        ],
        html: {
          height: '100%',
        },
        body: {
          backgroundColor: '#0f1014',
          backgroundImage: `url(${Symbol})`,
          backgroundRepeat: 'no-repeat',
          backgroundSize: 'auto 142vh',
          backgroundPosition: '70% -25vh',
          '@media (min-width: 600px)': {
            backgroundSize: 'auto 124vh',
            backgroundPosition: '-42vh -18vh',
          },
        },
      }
    },
  },
});
