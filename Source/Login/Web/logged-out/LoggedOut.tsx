// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import ChevronLeft from '@mui/icons-material/ChevronLeft';

import { configuration } from '../Configuration';

export const LoggedOut = (): JSX.Element => {
    return (
        <>
            <Box css={{ padding: '158px 64px 0 64px' }}>
                <Typography variant="h2" css={{ marginBottom: '30px' }}>You have been logged out.</Typography>
            </Box>
            <Button startIcon={<ChevronLeft/>} href={configuration.startPath} css={{ position: 'absolute', left: '20px', bottom: '20px' }}>Go back and log in again</Button>
        </>
    );
}
