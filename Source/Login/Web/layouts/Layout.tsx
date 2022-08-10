// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import { SxProps } from '@mui/system';

import { configuration } from '../Configuration';
import Symbol from '../styles/images/symbol.svg?url';
import Logo from '../styles/images/logo.svg';
import { Routes } from './Routes';

const dolittleBackgroundStyle: SxProps = configuration.showDolittleHeadline
    ? {
        backgroundImage: `url(${Symbol})`,
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'auto 142vh',
        backgroundPosition: '-60vh -25vh',
    }
    : {};

export const Layout = (): JSX.Element => (
    <Box sx={{
        textAlign: 'right',
        minHeight: '100vh',
        ...dolittleBackgroundStyle
    }}>
        <Box sx={{
            textAlign: 'center',
            width: '100vw',
            maxWidth: '33.8125rem',
            marginLeft: 'auto',
            '@media (min-width: 33.8125rem)': {
                marginRight: 'calc((100vw - 33.8125rem)*0.233)',
            },
            padding: '1.25rem',
        }}>
            <Routes />
            {
                // TODO: Why is there so much margin on the bottom?
                configuration.showDolittleHeadline &&
                    <Logo sx={{ width: 166, height: 39, mt: 18.5, mb: 18.5 }}/>
            }
        </Box>
    </Box>
);

export default Layout;
