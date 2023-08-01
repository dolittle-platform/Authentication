// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Routes } from './Routes';

import { Box, Link } from '@mui/material';

import Symbol from '../styles/images/symbol.svg?url';
import Logo from '../styles/images/logo.svg';

import { configuration } from '../Configuration';

const styles = {
    rightContainer: {
        textAlign: 'center',
        maxWidth: 540,
        p: 2.5,
        pt: 25,
        ml: 'auto',
        '@media (min-width: 33.8125rem)': {
            mr: 'calc((100vw - 33.8125rem)*0.233)',
        },
    },
    showDolittlelogo:
        configuration.showDolittleHeadline ? {
            backgroundImage: `url(${Symbol})`,
            backgroundRepeat: 'no-repeat',
            backgroundSize: 'auto 100vh',
        } : {},
    logo: {
        width: 166,
        height: 28,
    },
};

export const Layout = () =>
    <Box sx={{ minHeight: '100vh', ...styles.showDolittlelogo }}>
        <Box sx={styles.rightContainer}>
            <Routes />
            {configuration.showDolittleHeadline &&
                <Link href={configuration.startPath} sx={{ width: 1, mt: 18, display: 'block' }}>
                    <Logo sx={styles.logo} />
                </Link>
            }
        </Box>
    </Box>;

export default Layout;
