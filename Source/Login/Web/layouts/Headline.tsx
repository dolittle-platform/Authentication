// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';

import Logo from '../styles/images/logo.svg';

export const Headline = (): JSX.Element => {
    return (
        <Box css={{ position: 'absolute', top: '246px', left: '144px' }}>
            <Logo width={347} />
            <Typography
                variant="h1"
                css={{ marginTop: '9px', marginLeft: '102px', maxWidth: '480px' }}
            >Transforming your business with real time events.</Typography>
        </Box>
    );
};
