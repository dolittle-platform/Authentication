// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { createTheme } from '@material-ui/core/styles';

import { Rubik } from './fonts/Rubik';
import BackgroundImage from './images/background.jpg';

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
      fontFamily: 'Rubik-Regular',
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
  overrides: {
    MuiCssBaseline: {
      '@global': {
        '@font-face': [ Rubik.Regular, Rubik.Bold, Rubik.Medium ],
        html: {
          height: '100%',
        },
        body: {
          backgroundColor: '#242331',
          backgroundImage: `url(${BackgroundImage})`,
          backgroundPosition: 'top left',
          backgroundRepeat: 'no-repeat',
          backgroundSize: 'auto 100%',
        },
      },
    },
  },
});
