// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';

import { configuration } from '../Configuration';
import { Theme } from '../styles/Theme';
import Logo from '../styles/images/logo.svg';

export const Footer = (): JSX.Element => {
    return (<>{
        configuration.showDolittleHeadline &&
            <Box css={{
                textAlign: 'center',
                paddingBottom: '40px',
            }}>
                <Logo css={{
                    display: 'block',
                    margin: 'auto',
                    width: '120px',
                    [Theme.breakpoints.up('sm')]: {
                        width: '165px',
                        marginLeft: '40px',
                    },
                }} />
            </Box>
    }</>);
};
