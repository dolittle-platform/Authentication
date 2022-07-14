// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { Box } from '@mui/material';
import { BgLogo, MainLogo } from '../styles/logos';

interface Props {
    children: React.ReactNode;
}

const styles = {
    backgroundLogoContainer: {
        maxInlineSize: '793px',
        minBlockSize: '100vh'
    },
    mainContainer: {
        inlineSize: '100%',
        maxInlineSize: '541px',
        position: 'absolute',
        top: '20%',
        right: '20%',
        transform: 'translate(20%, 0%)',
        textAlign: 'center',
        padding: '20px'
    }
};

export const LoginWrapper: React.FunctionComponent<Props> = (props: Props) => {
    const children = props!.children;

    return (
        <>
            <Box sx={styles.backgroundLogoContainer}>
                <BgLogo />
            </Box>
            <Box sx={styles.mainContainer}>
                {children}
                <MainLogo mt={18.5} mb={18.5} />
            </Box>
        </>
    );
};